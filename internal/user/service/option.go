package service

import (
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/sms/tencent"
	prodFact "github.com/cold-runner/Hylark/internal/user/factory/production"
	prodRepo "github.com/cold-runner/Hylark/internal/user/repository/production"
	"github.com/cold-runner/Hylark/internal/user/store/mysql"
	"log/slog"

	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/cache/redis"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/oss/upyun"
)

type Option func(*Srv)

func empty(*Srv) {}

func WithCache(cacheType string) Option {
	switch cacheType {
	case "redis":
		return func(srv *Srv) {
			srv.Cache = redis.NewCache(&srv.RedisConfig)
		}
	default:
		slog.Warn("cache is not instantiated!")
	}
	return empty
}

func WithSms(smsType string) Option {
	switch smsType {
	case "tencent":
		return func(srv *Srv) {
			srv.Sms = tencent.NewSms(&srv.TencentSmsConfig)
		}
	default:
		slog.Warn("sms client is not instantiated!")
	}
	return empty
}

func WithRepository(repoType string) Option {
	switch repoType {
	case "mysql":
		return func(srv *Srv) {
			srv.Repository = prodRepo.NewRepository(mysql.NewStore(&srv.MysqlConfig))
		}
	default:
		slog.Warn("repository is not instantiated!")
	}
	return empty
}

func WithFactory(factoryType string) Option {
	switch factoryType {
	case "mysql":
		return func(srv *Srv) {
			srv.Factory = prodFact.NewFactory(mysql.NewStore(&srv.MysqlConfig))
		}
	default:
		slog.Warn("factory is not instantiated!")
	}
	return empty
}

func WithOss(ossType string) Option {
	switch ossType {
	case "upyun":
		return func(srv *Srv) {
			srv.Oss = upyun.NewOss(&srv.UpyunConfig)
		}
	default:
		slog.Warn("oss client is not instantiated!")
	}
	return empty
}
