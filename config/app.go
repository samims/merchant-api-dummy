package config

import (
	"github.com/samims/merchant-api/logger"
	"github.com/spf13/viper"
)

type AppConfig interface {
	GetBuildEnv() string
}

// config for app
type appConfig struct {
	env *viper.Viper
}

// GetBuildEnv return environment type
func (config *appConfig) GetBuildEnv() string {
	config.env.AutomaticEnv()
	return config.env.GetString("app_build_env")
}

// NewAppConfig initializes and return AppConfig
func NewAppConfig(env *viper.Viper) AppConfig {
	logger.Log.Info("App config reading...")
	return &appConfig{
		env: env,
	}
}
