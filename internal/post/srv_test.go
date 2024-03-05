package user

import (
	"context"
	"github.com/cold-runner/Hylark/internal/pkg"
	"github.com/cold-runner/Hylark/kitex_gen/post"
	"testing"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	"github.com/cold-runner/Hylark/kitex_gen/post/srv"
)

func TestSrv_Register(t *testing.T) {
	c := srv.MustNewClient("srv",
		client.WithHostPorts(":8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	ctx := context.Background()
	req := &post.CreatePostRequest{
		Token:      nil,
		UserId:     pkg.Convert(""),
		CategoryId: nil,
		Title:      nil,
		Summary:    nil,
		Content:    nil,
		Picture:    nil,
		LinkUrl:    nil,
		TagId:      nil,
	}

	resp, err := c.CreatePost(ctx, req)
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if isBizErr {
		t.Logf("%v, extra: %v", bizErr, bizErr.BizExtra()["msg"])
		return
	}

	t.Logf("register called success! resp: %v", *resp)

}

func Test_SendSmsCode(t *testing.T) {
	c := srv.MustNewClient("srv",
		client.WithHostPorts(":8888"),
		client.WithTransportProtocol(transport.TTHeader),
		client.WithMetaHandler(transmeta.ClientTTHeaderHandler),
	)
	ctx := context.Background()
	resp, err := c.CreatePost(ctx, &post.CreatePostRequest{})
	bizErr, isBizErr := kerrors.FromBizStatusError(err)
	if isBizErr {
		t.Logf("%v, extra: %v", bizErr, bizErr.BizExtra()["msg"])
		return
	}
	t.Logf("register called success! resp: %v", *resp)

}
