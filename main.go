package main

import (
	"fmt"
	"net/http"

	"github.com/samims/merchant-api/api/rest/router"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/logger"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	config := config.Init(v)

	r := router.Init(config)

	port := config.ApiConfig().Port()
	logger.Log.Infof("Server listening on port %s...", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
