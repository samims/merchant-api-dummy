package main

import (
	"fmt"
	"net/http"

	"github.com/samims/merchant-api/api/rest/router"
	"github.com/samims/merchant-api/app"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/logger"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	cfg := config.Init(v)

	svc := app.InitServices(cfg)

	r := router.Init(cfg, svc)

	port := cfg.ApiConfig().Port()
	logger.Log.Infof("Server listening on port %s...", port)

	http.ListenAndServe(fmt.Sprintf(":%s", port), r)
}
