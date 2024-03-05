package user

import (
	"context"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/kitex_gen/user"
	"github.com/cold-runner/Hylark/kitex_gen/user/srv"
)

func TestSrv_Register(t *testing.T) {
	c := srv.MustNewClient("srv",
		client.WithHostPorts(":8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	ctx := context.Background()
	req := &user.RegisterRequest{
		Phone:    pkg.Convert("18342728255"),
		Password: pkg.Convert("Aa123443"),
		SmsCode:  pkg.Convert("1218"),
	}

	resp, err := c.Register(ctx, req)
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if isBizErr {
		t.Logf("%v, extra: %v", bizErr, bizErr.BizExtra()["msg"])
		return
	}

	t.Logf("register called success! resp: %v", resp)

}

func Test_SendSmsCode(t *testing.T) {
	c := srv.MustNewClient("srv",
		client.WithHostPorts(":8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	ctx := context.Background()
	resp, err := c.SendSmsCode(ctx, &user.SendSmsCodeRequest{Phone: pkg.Convert("18342728255")})

	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if isBizErr {
		// 判断是否是成功响应
		t.Logf("%v, extra: %v", bizErr, bizErr.BizExtra()["msg"])
		return
	}
	t.Logf("register called success! resp: %v", resp)
}
