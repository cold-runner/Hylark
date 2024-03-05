package instance

import "github.com/upyun/go-sdk/v3/upyun"

type UpyunConfig struct {
	Bucket   string `mapstructure:"bucket"`
	Operator string `mapstructure:"operator"`
	Password string `mapstructure:"password"`
	UseHttp  bool   `mapstructure:"use-http"`
}

func NewUpyun(opt *UpyunConfig) *upyun.UpYun {
	return upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   opt.Bucket,
		Operator: opt.Operator,
		Password: opt.Password,
		UseHTTP:  opt.UseHttp,
	})
}
