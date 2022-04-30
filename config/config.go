package config

import "github.com/spf13/viper"

// Blueprint of configuration
type Configuration interface {
	ApiConfig() ApiConfig
}

// holds the required config instances
type configuration struct {
	apiConfig ApiConfig
}

// return the ApiConfig instance
func (config *configuration) ApiConfig() ApiConfig {
	return config.apiConfig
}

// return configuration with with the required param
func Init(
	v *viper.Viper,
) Configuration {
	return &configuration{
		apiConfig: NewApiConfig(v),
	}
}
