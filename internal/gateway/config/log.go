package config

import "C"
import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

type Log struct {
	Level  string `mapstructure:"level"`
	Output string `mapstructure:"output"`
	Format string `mapstructure:"format"`
}

func (l Log) validate(c *Conf) error {
	_, err := zap.ParseAtomicLevel(c.Log.Level)
	if err != nil {
		return errors.New("unsupported log level! [debug info warn error dpanic panic fatal]")
	}

	if l.Format != "json" && l.Format != "console" {
		return errors.New("unsupported format! [json console]")
	}

	return nil
}
