// Code generated by Kitex v0.8.0. DO NOT EDIT.

package srv

import (
	"context"
	client "github.com/cloudwego/kitex/client"
	callopt "github.com/cloudwego/kitex/client/callopt"
	user "github.com/cold-runner/Hylark/kitex_gen/user"
)

// Client is designed to provide IDL-compatible methods with call-option parameter for kitex framework.
type Client interface {
	Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error)
	SendSmsCode(ctx context.Context, req *user.SendSmsCodeRequest, callOptions ...callopt.Option) (r *user.SendSmsCodeResponse, err error)
	PasswordLogin(ctx context.Context, req *user.PasswordLoginRequest, callOptions ...callopt.Option) (r *user.PasswordLoginResponse, err error)
	Certificate(ctx context.Context, req *user.CertificateRequest, callOptions ...callopt.Option) (r *user.CertificateResponse, err error)
}

// NewClient creates a client for the service defined in IDL.
func NewClient(destService string, opts ...client.Option) (Client, error) {
	var options []client.Option
	options = append(options, client.WithDestService(destService))

	options = append(options, opts...)

	kc, err := client.NewClient(serviceInfo(), options...)
	if err != nil {
		return nil, err
	}
	return &kSrvClient{
		kClient: newServiceClient(kc),
	}, nil
}

// MustNewClient creates a client for the service defined in IDL. It panics if any error occurs.
func MustNewClient(destService string, opts ...client.Option) Client {
	kc, err := NewClient(destService, opts...)
	if err != nil {
		panic(err)
	}
	return kc
}

type kSrvClient struct {
	*kClient
}

func (p *kSrvClient) Register(ctx context.Context, req *user.RegisterRequest, callOptions ...callopt.Option) (r *user.RegisterResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Register(ctx, req)
}

func (p *kSrvClient) SendSmsCode(ctx context.Context, req *user.SendSmsCodeRequest, callOptions ...callopt.Option) (r *user.SendSmsCodeResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.SendSmsCode(ctx, req)
}

func (p *kSrvClient) PasswordLogin(ctx context.Context, req *user.PasswordLoginRequest, callOptions ...callopt.Option) (r *user.PasswordLoginResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.PasswordLogin(ctx, req)
}

func (p *kSrvClient) Certificate(ctx context.Context, req *user.CertificateRequest, callOptions ...callopt.Option) (r *user.CertificateResponse, err error) {
	ctx = client.NewCtxWithCallOptions(ctx, callOptions)
	return p.kClient.Certificate(ctx, req)
}
