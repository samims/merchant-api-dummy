package config

import (
	"github.com/spf13/viper"
)

// Blueprint of configuration
type Configuration interface {
	AppConfig() AppConfig
	ApiConfig() ApiConfig
}

// holds the required config instances
type configuration struct {
	apiConfig ApiConfig
	appConfig AppConfig
}

// return the ApiConfig instance
func (config *configuration) ApiConfig() ApiConfig {
	return config.apiConfig

}

func (config *configuration) AppConfig() AppConfig {
	return config.appConfig
}

// return configuration with with the required param
func Init(
	v *viper.Viper,
) Configuration {
	return &configuration{
		apiConfig: NewApiConfig(v),
		appConfig: NewAppConfig(v),
	}
}
