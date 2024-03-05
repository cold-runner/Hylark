package instance

type ZincClientConfig struct {
	Host     string `mapstructure:"host"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	UseHttps bool   `mapstructure:"use-https"`
}
