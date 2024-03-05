package api

import (
	"github.com/cold-runner/Hylark/internal/gateway/plugin"

	"github.com/cloudwego/hertz/pkg/app"
)

type Api struct {
	name        string
	path        string // path is unique
	method      string
	description string

	plugins []plugin.Mw
}

func NewApi(name, path, method, description string) *Api {
	return &Api{
		name:        name,
		path:        path,
		method:      method,
		description: description,
	}
}

func (a *Api) AttachMw(mws ...plugin.Mw) {
	a.plugins = mws
}

func (a *Api) Handle(ctx *app.RequestContext) error {
	for _, mw := range a.plugins {
		if err := mw.Handle(ctx); err != nil {
			return err
		}
	}
	return nil
}

func (a *Api) GetPlugins() []plugin.Mw {
	return a.plugins
}
