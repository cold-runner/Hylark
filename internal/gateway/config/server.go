package config

import (
	"github.com/pkg/errors"
	"net"
	"os"
	"strconv"
)

type Server struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
	Tls  `mapstructure:"tls"`
}

type Tls struct {
	Use           bool   `mapstructure:"use"`
	CACertificate string `mapstructure:"CA"`
	Certificate   string `mapstructure:"certificate"`
	PrivateKey    string `mapstructure:"privateKey"`
}

func (s Server) validate(c *Conf) error {
	if net.ParseIP(s.Host) == nil {
		return errors.New("host is invalid!")
	}
	p, err := strconv.Atoi(s.Port)
	if err != nil {
		return errors.New("port is invalid!")
	}
	if p < 1024 || p > 49151 {
		return errors.New("unsupported port range, valid range is [1024, 49151]")
	}

	if s.Tls.Use {
		if _, err := os.ReadFile(s.Tls.CACertificate); err != nil {
			return errors.Errorf("read CA file failed! err: %v", err)
		}
		if _, err := os.ReadFile(s.Tls.Certificate); err != nil {
			return errors.Errorf("read tls cert file failed! err: %v", err)
		}
		if _, err := os.ReadFile(s.Tls.PrivateKey); err != nil {
			return errors.Errorf("read tls cert file failed! err: %v", err)
		}
	}

	return nil
}
