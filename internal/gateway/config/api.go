package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var apis ApiList

type ApiList struct {
	APIList []struct {
		APIInfo struct {
			Name        string                    `mapstructure:"name"`
			Path        string                    `mapstructure:"path"`
			Method      string                    `mapstructure:"method"`
			Description string                    `mapstructure:"description"`
			MWS         map[string]map[string]any `mapstructure:"mws"`
		} `mapstructure:"apiInfo"`
	} `mapstructure:"apiList"`
}

func init() {
	v := viper.New()
	v.SetConfigFile("api.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(errors.Errorf("fatal read config file! err: %v", err))
	}
	if err := v.Unmarshal(&apis); err != nil {
		panic(errors.Errorf("fatal unmarshal config file! err: %v", err))
	}
	// TODO 配置校验
}

func GetApiList() ApiList {
	if c == nil {
		panic("apiList is not init")
	}
	return apis
}
