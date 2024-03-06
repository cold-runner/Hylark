package user

import (
	"context"
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/kitex_gen/user"
	"github.com/cold-runner/Hylark/kitex_gen/user/srv"
	"testing"
)

var (
	cli = srv.MustNewClient("userSrv",
		client.WithHostPorts(":8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	c = context.Background()
)

func TestRegister(t *testing.T) {
	req := &user.RegisterRequest{
		Phone:    pkg.Convert("13942321313"),
		Password: pkg.Convert("Aa123654"),
		SmsCode:  pkg.Convert("123124"),
	}
	resp, err := cli.Register(c, req)
	if resp == nil {
		t.Errorf("err: %v", err)
		return
	}

	t.Logf("%#v", resp)
}

func TestPhonePasswordLogin(t *testing.T) {
	req := &user.PasswordLoginRequest{
		Phone:    pkg.Convert("13942321313"),
		Password: pkg.Convert("Aa123654"),
	}
	resp, err := cli.PasswordLogin(c, req)

	if err != nil {
		bizErr, isBizErr := kerrors.FromBizStatusError(err)
		if !isBizErr {
			t.Errorf("rpc call failed. err: %v", err)
			return
		}
		t.Errorf("biz err occured. err: %v, extra: %v", bizErr, bizErr.BizExtra())
		return
	}

	t.Logf("%v", resp.GetToken())
}

func TestUpdateInfo(t *testing.T) {
	req := &user.UpdateUserInfoRequest{
		Token:        pkg.Convert("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1dWlkIjoiZjI1NjU0NDItM2E0OC00MTc4LTgyMmEtNjRhMTczOWNlZTcxIiwiaXNzIjoiU1lOVS1za3lsYWIiLCJzdWIiOiJIeWxhcmsiLCJleHAiOjE3MDk2ODk3OTZ9.A-k2C-FVv3OzCeiXY1DnkYUOuXkoYeOVi--rqHpcjDk"),
		Gender:       pkg.Convert("男"),
		College:      nil,
		Major:        nil,
		Grade:        pkg.Convert("大四"),
		Province:     pkg.Convert("辽宁"),
		Age:          nil,
		Introduction: pkg.Convert("一些介绍"),
		Avatar:       nil,
	}

	resp, err := cli.UpdateUserInfo(c, req)
	if err != nil {
		bizErr, isBizErr := kerrors.FromBizStatusError(err)
		if !isBizErr {
			t.Errorf("rpc call failed. err: %v", err)
			return
		}
		t.Errorf("biz err occured. err: %v, extra: %v", bizErr, bizErr.BizExtra())
		return
	}

	fmt.Println(resp.String())
}
