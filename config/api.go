package config

import "github.com/spf13/viper"

type ApiConfig interface {
	Port() string
}

// configuration for the API
type apiConfig struct {
	env *viper.Viper
}

// Returns the port the api is running
func (config *apiConfig) Port() string {
	config.env.AutomaticEnv()
	port := config.env.GetString("api_port")
	return port
}

func NewApiConfig(env *viper.Viper) ApiConfig {
	return &apiConfig{
		env: env,
	}
}
