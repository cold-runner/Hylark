package plugin

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/pkg/errors"
)

type MwName string

const (
	AUTH    MwName = "auth"
	LIMITER        = "limiter"
	FUSE           = "fuse"
)

type Mw interface {
	Handle(ctx *app.RequestContext) error
}

func NewPlugin(mwName MwName) Mw {
	switch mwName {
	case AUTH:
		return NewAuth()
	case LIMITER:
		return NewLimiter()
	case FUSE:
		return NewFuse()
	}
	return nil
}

func ParseMw(mwName string) (MwName, error) {
	switch mwName {
	case "auth":
		return AUTH, nil
	case "limiter":
		return LIMITER, nil
	case "fuse":
		return FUSE, nil
	default:
		return "", errors.New("unsupported middleware!")
	}
}
