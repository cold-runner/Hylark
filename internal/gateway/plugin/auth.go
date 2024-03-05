package plugin

import (
	"github.com/cloudwego/hertz/pkg/app"
)

type Auth struct {
}

func NewAuth() Mw {
	return Auth{}
}

func (a Auth) Handle(ctx *app.RequestContext) error {

	return nil
}
