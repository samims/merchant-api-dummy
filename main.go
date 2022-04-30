package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/logger"
	"github.com/spf13/viper"
)

func main() {
	v := viper.New()
	config := config.Init(v)

	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		usedPort := config.ApiConfig().Port()
		logger.Log.Info("Used port " + usedPort)
		w.Write([]byte("Welcome"))
	})
	logger.Log.Info("Server listening on port 3000")

	http.ListenAndServe(":3000", r)
}
