package gw

import (
	"context"

	"Hylark/internal/gateway/api"
	"Hylark/internal/gateway/response"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func (g *Gateway) createApi(c context.Context, ctx *app.RequestContext) {
	path := ctx.Query("path")
	if _, ok := g.apis[path]; ok {
		ctx.JSON(consts.StatusBadRequest, response.RespApiAlreadyExist)
		return
	}

	g.apis[path] = api.NewApi(
		ctx.Query("name"),
		ctx.Query("method"),
		path,
		ctx.Query("description"),
	)

	ctx.JSON(consts.StatusOK, nil)
}

func (g *Gateway) deleteApi(c context.Context, ctx *app.RequestContext) {
	delete(g.apis, ctx.Query("api"))
	ctx.JSON(consts.StatusOK, nil)
}

func (g *Gateway) updateApi(c context.Context, ctx *app.RequestContext) {
	delete(g.apis, ctx.Query("api"))
	ctx.JSON(consts.StatusOK, nil)
}

func (g *Gateway) getApiList(c context.Context, ctx *app.RequestContext) {
	ctx.JSON(consts.StatusOK, g.apis)
}

func (g *Gateway) getApiPlugins(c context.Context, ctx *app.RequestContext) {
	param := ctx.Param("specieApi")
	specieApi, ok := g.apis[param]
	if !ok {
		ctx.JSON(consts.StatusBadRequest, response.RespApiNotExist)
		return
	}
	ctx.JSON(consts.StatusOK, specieApi.GetPlugins())
}
