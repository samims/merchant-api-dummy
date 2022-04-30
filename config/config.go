package config

import (
	"github.com/spf13/viper"
)

// Blueprint of configuration
type Configuration interface {
	AppConfig() AppConfig
	ApiConfig() ApiConfig
	PostgresConfig() PostgresConfig
}

// holds the required config instances
type configuration struct {
	apiConfig      ApiConfig
	appConfig      AppConfig
	postgresConfig PostgresConfig
}

// return the ApiConfig instance
func (config *configuration) ApiConfig() ApiConfig {
	return config.apiConfig

}

func (config *configuration) AppConfig() AppConfig {
	return config.appConfig
}

// return the PostgresConfig instance
func (config *configuration) PostgresConfig() PostgresConfig {
	return config.postgresConfig
}

// return configuration with with the required param
func Init(
	v *viper.Viper,
) Configuration {
	return &configuration{
		apiConfig:      NewApiConfig(v),
		appConfig:      NewAppConfig(v),
		postgresConfig: NewPostgresConfig(v),
	}
}
