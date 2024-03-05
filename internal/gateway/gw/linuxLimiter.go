//go:build linux

package gw

import (
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/hertz-contrib/limiter"
)

func addLimiter(h *server.Hertz) {
	h.Use(limiter.AdaptiveLimit())
}
