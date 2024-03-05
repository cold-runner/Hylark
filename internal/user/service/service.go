package service

import (
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/cache"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/oss"
	"github.com/cold-runner/Hylark/internal/pkg/infrastructure/sms"
	"github.com/cold-runner/Hylark/internal/pkg/instance"
	"github.com/cold-runner/Hylark/internal/user/factory"
	"github.com/cold-runner/Hylark/internal/user/repository"
	"github.com/cold-runner/Hylark/kitex_gen/user"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"time"
)

type Srv struct {
	// 其他服务的客户端

	cache.Cache
	sms.Sms
	oss.Oss

	factory.Factory
	repository.Repository

	Config
}

type JwtConfig struct {
	Key        string        `mapstructure:"key"`
	Algorithm  string        `mapstructure:"algorithm"`
	Issuer     string        `mapstructure:"issuer"`
	ExpireTime time.Duration `mapstructure:"expire-time"`
	Subject    string        `mapstructure:"subject"`
}

type Config struct {
	JwtConfig                 `mapstructure:"jwt"`
	instance.MysqlConfig      `mapstructure:"mysql"`
	instance.RedisConfig      `mapstructure:"redis"`
	instance.TencentSmsConfig `mapstructure:"tencent-sms"`
	instance.UpyunConfig      `mapstructure:"upyun"`
}

func NewSrv(opts ...Option) user.Srv {
	s := &Srv{}

	s.initConfig()

	for _, o := range opts {
		o(s)
	}

	return s
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
