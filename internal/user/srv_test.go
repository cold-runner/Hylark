package user

import (
	"context"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/kitex_gen/user"
	"github.com/cold-runner/Hylark/kitex_gen/user/srv"
	"testing"
)

func TestRegister(t *testing.T) {
	cli := srv.MustNewClient("userSrv",
		client.WithHostPorts(":8888"),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	req := &user.RegisterRequest{
		Phone:    pkg.Convert("13942321313"),
		Password: pkg.Convert("Aa123654"),
		SmsCode:  pkg.Convert("123123"),
	}
	resp, err := cli.Register(context.Background(), req)
	if resp == nil {
		t.Errorf("err: %v", err)
		return
	}

	t.Logf("%#v", resp)
}

func TestPhonePasswordLogin(t *testing.T) {
	cli := srv.MustNewClient("srv",
		client.WithHostPorts(":8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)

	req := &user.PasswordLoginRequest{
		Phone:    pkg.Convert("18342728255"),
		Password: pkg.Convert("Aa12365"),
	}
	resp, err := cli.PasswordLogin(context.Background(), req)

	if err != nil {
		bizErr, isBizErr := kerrors.FromBizStatusError(err)
		if isBizErr {
			t.Errorf("biz err occured! err: %v", bizErr)
			return
		}
		t.Errorf("rpc call err! err: %v", err)
	}

	t.Logf("%v", resp.GetToken())
}
