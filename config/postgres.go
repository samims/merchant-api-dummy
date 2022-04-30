package config

import (
	"fmt"

	"github.com/samims/merchant-api/logger"
	"github.com/spf13/viper"
)

type PostgresConfig interface {
	Host() string
	Port() string
	Database() string
	User() string
	Password() string
	ConnectionURL() string
}

type postgresConfig struct {
	env *viper.Viper
}

func (cfg *postgresConfig) Host() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_host")
}

func (cfg *postgresConfig) Port() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_port")
}

func (cfg *postgresConfig) Database() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_database")
}

//
func (cfg *postgresConfig) User() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_user")
}

//
func (cfg *postgresConfig) Password() string {
	cfg.env.AutomaticEnv()
	return cfg.env.GetString("postgres_password")
}

// ConnectionURL returns connection url for postgresql database
func (config *postgresConfig) ConnectionURL() string {
	url := config.env.GetString("postgres_url")
	if len(url) > 0 {
		return url
	}
	return fmt.Sprintf(`postgres://%v:%v@%v:%v/%v?sslmode=disable`, config.User(), config.Password(), config.Host(), config.Port(), config.Database())
}

func NewPostgresConfig(env *viper.Viper) PostgresConfig {
	logger.Log.Info("DB config reading...")
	return &postgresConfig{
		env: env,
	}
}
