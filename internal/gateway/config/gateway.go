package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var c *Conf

func init() {
	v := viper.New()
	v.SetConfigFile("config.yaml")
	if err := v.ReadInConfig(); err != nil {
		panic(errors.Errorf("fatal read config file! err: %v", err))
	}

	c = &Conf{}
	if err := v.Unmarshal(c); err != nil {
		panic(errors.Errorf("fatal unmarshal config file! err: %v", err))
	}

	if err := c.Log.validate(c); err != nil {
		panic(err)
	}

	if err := c.Server.validate(c); err != nil {
		panic(err)
	}

	if err := c.Auth.validate(c); err != nil {
		panic(err)
	}

}

type Conf struct {
	Server Server `mapstructure:"server"`
	Auth   Auth   `mapstructure:"auth"`
	Log    Log    `mapstructure:"log"`
}

func GetConfig() *Conf {
	if c == nil {
		panic("config is not init")
	}
	return c
}
