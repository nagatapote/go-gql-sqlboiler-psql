package config

type Config struct {
	DatabaseConfig
}

var Conf Config

type DatabaseConfig struct {
	DatabaseURL string `envconfig:"DATABASE_URL" required:"true"`
	SSLMode     string `envconfig:"POSTGRES_SSLMODE" required:"true"`
}
