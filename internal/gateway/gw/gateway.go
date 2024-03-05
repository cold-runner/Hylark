package gw

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cold-runner/Hylark/internal/gateway/api"
	"github.com/cold-runner/Hylark/internal/gateway/config"
	"github.com/cold-runner/Hylark/internal/gateway/plugin"
	"github.com/cold-runner/Hylark/internal/gateway/response"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/app/server"
	hertzconfig "github.com/cloudwego/hertz/pkg/common/config"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/genericclient"
	"github.com/cloudwego/kitex/pkg/generic"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	hertzzap "github.com/hertz-contrib/logger/zap"
	"github.com/pkg/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"golang.org/x/net/http2"
	"gopkg.in/natefinch/lumberjack.v2"
)

type Gateway struct {
	*server.Hertz
	apis    map[string]*api.Api
	clients map[string]genericclient.Client
	config  *config.Conf
	syncer  time.Ticker
}

func New() *Gateway {
	g := &Gateway{config: config.GetConfig()}

	g.initLog()
	g.initHertz()
	g.initClients()
	g.initApis()
	g.initRouter()

	return g
}

func (g *Gateway) initLog() {
	c := g.config
	// 日志
	logLevel, err := zap.ParseAtomicLevel(c.Log.Level)
	if err != nil {
		panic(errors.New("unsupported log level!"))
	}
	enc := zapcore.EncoderConfig{
		MessageKey:     "msg",
		LevelKey:       "level",
		NameKey:        "name",
		TimeKey:        "ts",
		CallerKey:      "caller",
		FunctionKey:    "func",
		StacktraceKey:  "stacktrace",
		LineEnding:     "\n",
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
	var coreEnc zapcore.Encoder
	switch c.Log.Format {
	case "json":
		coreEnc = zapcore.NewJSONEncoder(enc)
	case "console":
		coreEnc = zapcore.NewConsoleEncoder(enc)
	default:
		panic(errors.New("unsupported log response!"))
	}
	logger := hertzzap.NewLogger(
		hertzzap.WithCoreLevel(logLevel),
		hertzzap.WithCoreEnc(coreEnc),
	)
	if c.Log.Output != "stdout" {
		lumberJackLogger := &lumberjack.Logger{
			Filename:   "log/" + c.Log.Output,
			MaxSize:    10,
			MaxBackups: 50000,
			MaxAge:     1000,
			Compress:   true,
			LocalTime:  true,
		}
		logger.SetOutput(lumberJackLogger)
	}

	hlog.SetLogger(logger)
}

func (g *Gateway) initHertz() {
	conf := g.config
	var opts []hertzconfig.Option
	opts = append(opts, server.WithHostPorts(conf.Server.Host+":"+conf.Server.Port))
	if !g.config.Server.Tls.Use {
		hlog.Warn("tls is not used")
	} else {
		opts = append(opts, server.WithTLS(tlsConfig(&conf.Server.Tls)))
	}
	g.Hertz = server.Default(opts...)
}

func tlsConfig(c *config.Tls) *tls.Config {
	// load server certificate
	cert, err := tls.LoadX509KeyPair(c.Certificate, c.PrivateKey)
	if err != nil {
		hlog.Fatalf("load server certificate file failed! err: %v", err)
	}
	// load root certificate
	certBytes, err := os.ReadFile(c.CACertificate)
	if err != nil {
		hlog.Fatalf("load CA file failed! err: %v", err)
	}
	caCertPool := x509.NewCertPool()
	ok := caCertPool.AppendCertsFromPEM(certBytes)
	if !ok {
		hlog.Fatalf("Failed to parse root certificate! err: %v", err)
	}
	return &tls.Config{
		// add certificate
		Certificates: []tls.Certificate{cert},
		MaxVersion:   tls.VersionTLS13,
		// enable client authentication
		ClientAuth: tls.RequireAndVerifyClientCert,
		ClientCAs:  caCertPool,
		// cipher suites supported
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
		},
		// set application protocol http2
		NextProtos: []string{http2.NextProtoTLS},
	}
}

func (g *Gateway) initClients() {
	g.clients = make(map[string]genericclient.Client)
	idlPath := "../../idl"
	dirs, err := os.ReadDir(idlPath)
	if err != nil {
		hlog.Fatalf("new thrift file provider failed: %v", err)
	}

	for _, entry := range dirs {
		// 只注册文件名包含Srv的thrift文件
		if entry.IsDir() || !strings.Contains(entry.Name(), "Srv") {
			continue
		}
		provider, err := generic.NewThriftFileProvider(idlPath + "/" + entry.Name())
		if err != nil {
			hlog.Fatalf("new thrift file provider failed: %v", err)
		}
		thriftGeneric, err := generic.HTTPThriftGeneric(provider)
		if err != nil {
			hlog.Fatal("new http thrift thriftGeneric failed: %v", err)
		}
		// thrift file name is service name
		svcName := strings.ReplaceAll(entry.Name(), ".thrift", "")
		cli, err := genericclient.NewClient(
			svcName,
			thriftGeneric,
			client.WithTransportProtocol(transport.TTHeader),
			client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
		)
		if err != nil {
			hlog.Fatalf("new thriftGeneric client failed: %v", err)
		}

		// 注册客户端
		g.clients[svcName] = cli
	}
}

func (g *Gateway) initApis() {
	apiList := config.GetApiList()
	g.apis = make(map[string]*api.Api)
	for _, apiInfo := range apiList.APIList {
		info := apiInfo.APIInfo
		g.apis[info.Path] = api.NewApi(
			info.Name,
			info.Path,
			info.Path,
			info.Description,
		)
		var mws []plugin.Mw
		for mw := range info.MWS {
			middleware, err := plugin.ParseMw(mw)
			if err != nil {
				hlog.Error("parse middleware failed! err: %v", err)
				continue
			}
			mws = append(mws, plugin.NewPlugin(middleware))
		}
		g.apis[info.Path].AttachMw(mws...)
	}
}

func (g *Gateway) initRouter() {
	addLimiter(g.Hertz)

	adminGroup := g.Group("/admin")
	{
		adminGroup.POST("/api/create", g.createApi)
		adminGroup.DELETE("/api/delete", g.deleteApi)
		adminGroup.POST("/api/update", g.updateApi)
		adminGroup.GET("/apiList", g.getApiList)
		adminGroup.GET("/:api/plugins", g.getApiPlugins)
	}

	g.Any("/:svc/:api", g.process)
}

func (g *Gateway) process(ctx context.Context, c *app.RequestContext) {
	reqSvc := c.Param("svc")
	cli, ok := g.clients[reqSvc]
	if !ok {
		c.JSON(consts.StatusNotFound, response.RespSvcNotExist)
		hlog.Debug("svc not found")
		return
	}

	reqApi := c.Param("specificApi")
	specificApi, ok := g.apis[reqSvc+reqApi]
	if !ok {
		c.JSON(consts.StatusNotFound, response.RespApiNotExist)
		hlog.Debug("specificApi not found")
		return
	}

	// 中间件处理
	if err := specificApi.Handle(c); err != nil {
		// 错误处理
		c.JSON(consts.StatusBadRequest, err)
	}

	// 构建请求
	// TODO path处理
	req, err := http.NewRequest(string(c.Method()), string(c.Path()), c.RequestBodyStream())
	if err != nil {
		c.JSON(consts.StatusInternalServerError, response.RespServerError)
		hlog.Errorf("new http request failed: %v", err)
		return
	}

	customReq, err := generic.FromHTTPRequest(req)
	if err != nil {
		c.JSON(consts.StatusInternalServerError, response.RespServerError)
		hlog.Errorf("convert request failed: %v", err)
		return
	}

	// 泛化调用
	resp, err := cli.GenericCall(ctx, "", customReq)

	// TODO 处理结果
	if err != nil {
		c.JSON(consts.StatusInternalServerError, response.RespServerError)
		hlog.Errorf("genericCall err:%v", err)
		return
	}
	realResp, ok := resp.(*generic.HTTPResponse)
	if !ok {
		c.JSON(consts.StatusInternalServerError, response.RespServerError)
		hlog.Errorf("format resp err:%v", err)
		return
	}

	c.JSON(consts.StatusOK, realResp)
}
