// Code generated by Kitex v0.8.0. DO NOT EDIT.

package srv

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	kitex "github.com/cloudwego/kitex/pkg/serviceinfo"
	user "github.com/cold-runner/Hylark/kitex_gen/user"
)

func serviceInfo() *kitex.ServiceInfo {
	return srvServiceInfo
}

var srvServiceInfo = NewServiceInfo()

func NewServiceInfo() *kitex.ServiceInfo {
	serviceName := "srv"
	handlerType := (*user.Srv)(nil)
	methods := map[string]kitex.MethodInfo{
		"Register":    kitex.NewMethodInfo(registerHandler, newSrvRegisterArgs, newSrvRegisterResult, false),
		"SendSmsCode": kitex.NewMethodInfo(sendSmsCodeHandler, newSrvSendSmsCodeArgs, newSrvSendSmsCodeResult, false),
		"Certificate": kitex.NewMethodInfo(certificateHandler, newSrvCertificateArgs, newSrvCertificateResult, false),
	}
	extra := map[string]interface{}{
		"PackageName":     "user",
		"ServiceFilePath": `idl/userSrv.thrift`,
	}
	svcInfo := &kitex.ServiceInfo{
		ServiceName:     serviceName,
		HandlerType:     handlerType,
		Methods:         methods,
		PayloadCodec:    kitex.Thrift,
		KiteXGenVersion: "v0.8.0",
		Extra:           extra,
	}
	return svcInfo
}

func registerHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.SrvRegisterArgs)
	realResult := result.(*user.SrvRegisterResult)
	success, err := handler.(user.Srv).Register(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSrvRegisterArgs() interface{} {
	return user.NewSrvRegisterArgs()
}

func newSrvRegisterResult() interface{} {
	return user.NewSrvRegisterResult()
}

func sendSmsCodeHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.SrvSendSmsCodeArgs)
	realResult := result.(*user.SrvSendSmsCodeResult)
	success, err := handler.(user.Srv).SendSmsCode(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSrvSendSmsCodeArgs() interface{} {
	return user.NewSrvSendSmsCodeArgs()
}

func newSrvSendSmsCodeResult() interface{} {
	return user.NewSrvSendSmsCodeResult()
}

func certificateHandler(ctx context.Context, handler interface{}, arg, result interface{}) error {
	realArg := arg.(*user.SrvCertificateArgs)
	realResult := result.(*user.SrvCertificateResult)
	success, err := handler.(user.Srv).Certificate(ctx, realArg.Req)
	if err != nil {
		return err
	}
	realResult.Success = success
	return nil
}
func newSrvCertificateArgs() interface{} {
	return user.NewSrvCertificateArgs()
}

func newSrvCertificateResult() interface{} {
	return user.NewSrvCertificateResult()
}

type kClient struct {
	c client.Client
}

func newServiceClient(c client.Client) *kClient {
	return &kClient{
		c: c,
	}
}

func (p *kClient) Register(ctx context.Context, req *user.RegisterRequest) (r *user.RegisterResponse, err error) {
	var _args user.SrvRegisterArgs
	_args.Req = req
	var _result user.SrvRegisterResult
	if err = p.c.Call(ctx, "Register", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) SendSmsCode(ctx context.Context, req *user.SendSmsCodeRequest) (r *user.SendSmsCodeResponse, err error) {
	var _args user.SrvSendSmsCodeArgs
	_args.Req = req
	var _result user.SrvSendSmsCodeResult
	if err = p.c.Call(ctx, "SendSmsCode", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}

func (p *kClient) Certificate(ctx context.Context, req *user.CertificateRequest) (r *user.CertificateResponse, err error) {
	var _args user.SrvCertificateArgs
	_args.Req = req
	var _result user.SrvCertificateResult
	if err = p.c.Call(ctx, "Certificate", &_args, &_result); err != nil {
		return
	}
	return _result.GetSuccess(), nil
}
