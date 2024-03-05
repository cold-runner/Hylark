package plugin

import "github.com/cloudwego/hertz/pkg/app"

type Fuse struct {
}

func NewFuse() Mw {
	return &Fuse{}
}

func (f Fuse) Handle(ctx *app.RequestContext) error {
	return nil
}
