package main

import (
	"log/slog"
	"os"
	"os/signal"
	"time"

	"github.com/cold-runner/Hylark/internal/user/service"
	"github.com/cold-runner/Hylark/kitex_gen/user/srv"

	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/server"
)

func main() {
	s := srv.NewServer(
		service.NewSrv(
			service.WithCache("redis"),
			service.WithSms("tencent"),
			service.WithOss("upyun"),
			service.WithRepository("mysql"),
			service.WithFactory("mysql"),
		),
		server.WithMetaHandler(transmeta.ServerTTHeaderHandler),
		server.WithExitWaitTime(5*time.Second),
	)

	go func() {
		exitChan := make(chan os.Signal)
		signal.Notify(exitChan, os.Interrupt, os.Kill)
		select {
		case <-exitChan:
			slog.Info("<CR+C> exit!")
			_ = s.Stop()
		}
	}()

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
