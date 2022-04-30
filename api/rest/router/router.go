package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samims/merchant-api/api/rest/controllers"
	"github.com/samims/merchant-api/app"
	"github.com/samims/merchant-api/config"
)

// Init initializes router
func Init(cfg config.Configuration, svc app.Services) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)

	pingController := controllers.NewPing(cfg, svc)

	router.Get("/", pingController.Get)
	return router

}
