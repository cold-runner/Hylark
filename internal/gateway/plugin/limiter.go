package plugin

import "github.com/cloudwego/hertz/pkg/app"

type Limiter struct {
}

func NewLimiter() Mw {
	return &Limiter{}
}

func (l Limiter) Handle(ctx *app.RequestContext) error {
	return nil
}
