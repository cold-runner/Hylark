package service

import (
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/cache"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/oss"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/searchEngine"
	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"github.com/cold-runner/Hylark/internal/post/factory"
	"github.com/cold-runner/Hylark/internal/post/repository"
	"github.com/cold-runner/Hylark/kitex_gen/post"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Srv struct {
	// 其他服务的客户端

	cache.Cache
	searchEngine.SearchEngine
	oss.Oss

	factory.Factory
	repository.Repository

	Config
}

type Config struct {
	instance.MysqlConfig      `mapstructure:"mysql"`
	instance.RedisConfig      `mapstructure:"redis"`
	instance.UpyunConfig      `mapstructure:"upyun"`
	instance.ZincClientConfig `mapstructure:"zinc"`
}

func (s *Srv) initConfig() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(errors.Errorf("fatal read config file! err: %v", err))
	}

	if err := v.Unmarshal(&s.Config); err != nil {
		panic(errors.Errorf("user service unmarshal config file failed! err: %v", err))
	}
}

func NewSrv(opts ...Option) post.Srv {
	s := &Srv{}

	s.initConfig()

	for _, o := range opts {
		o(s)
	}

	return s
}
