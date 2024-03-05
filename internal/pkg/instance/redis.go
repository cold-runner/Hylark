package instance

import (
	"context"
	"strconv"

	"github.com/redis/go-redis/v9"
)

type RedisConfig struct {
	Host                  string `json:"host"                     mapstructure:"host"                     description:"Cache application host address"`
	Port                  int    `json:"port"                     mapstructure:"port"`
	User                  string `json:"username"                 mapstructure:"user"`
	Password              string `json:"password"                 mapstructure:"password"`
	Database              int    `json:"database"                 mapstructure:"database"`
	MasterName            string `json:"master-name"              mapstructure:"master-name"`
	MaxIdle               int    `json:"optimisation-max-idle"    mapstructure:"optimisation-max-idle"`
	MaxActive             int    `json:"optimisation-max-active"  mapstructure:"optimisation-max-active"`
	Timeout               int    `json:"timeout"                  mapstructure:"timeout"`
	EnableCluster         bool   `json:"enable-cluster"           mapstructure:"enable-cluster"`
	UseSSL                bool   `json:"use-ssl"                  mapstructure:"use-SSL"`
	SSLInsecureSkipVerify bool   `json:"ssl-insecure-skip-verify" mapstructure:"SSL-insecure-skip-verify"`
}

func NewRedis(opt *RedisConfig) (*redis.Client, error) {
	var err error
	client := redis.NewClient(&redis.Options{
		Addr:     opt.Host + ":" + strconv.Itoa(opt.Port),
		Username: opt.User,
		Password: opt.Password,
		DB:       opt.Database,
	})
	_, err = client.Ping(context.Background()).Result()

	if err != nil {
		return nil, err
	}
	return client, nil
}
