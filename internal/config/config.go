package config

import (
	"github.com/joho/godotenv"
	"os"
	"time"
)

const (
	EnvLocal = "local"
)

type (
	Config struct {
		Environment string
		Postgres    PostgresConfig
		Http        HttpConfig
	}
	PostgresConfig struct {
		UserName string `mapstructure:"username"`
		Password string `mapstructure:"password"`
		Host     string `mapstructure:"host"`
		Port     string `mapstructure:"port"`
		DbName   string `mapstructure:"dbname"`
		SslMode  string `mapstructure:"sslmode"`
	}

	HttpConfig struct {
		Host               string        `mapstructure:"host"`
		Port               string        `mapstructure:"port"`
		ReadTimeout        time.Duration `mapstructure:"readTimeout"`
		WriteTimeout       time.Duration `mapstructure:"writeTimeout"`
		MaxHeaderMegabytes int           `mapstructure:"maxHeaderBytes"`
	}
)

func Init() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, err
	}
	var cfg Config
	setFromEnv(&cfg)

	return &cfg, nil
}

func setFromEnv(cfg *Config) {
	cfg.Postgres.UserName = os.Getenv("POSTGRES_USERNAME")
	cfg.Postgres.Password = os.Getenv("POSTGRES_PASSWORD")
	cfg.Postgres.Host = os.Getenv("POSTGRES_HOST")
	cfg.Postgres.Port = os.Getenv("POSTGRES_PORT")
	cfg.Postgres.DbName = os.Getenv("POSTGRES_DB_NAME")
	cfg.Postgres.SslMode = os.Getenv("POSTGRES_SSL_MODE")

	cfg.Http.Port = os.Getenv("HTTP_PORT")
	cfg.Environment = os.Getenv("APP_ENV")
}
