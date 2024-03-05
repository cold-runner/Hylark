package config

type Auth struct {
	Use         bool   `mapstructure:"use"`
	IdentityKey string `mapstructure:"identityKey"`
	Key         string `mapstructure:"key"`
}

func (a Auth) validate(c *Conf) error {
	return nil
}
