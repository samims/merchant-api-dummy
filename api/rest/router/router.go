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
	userController := controllers.NewUser(cfg, svc)
	merchantController := controllers.NewMerchant(cfg, svc)

	router.Get("/ping", pingController.Get)
	router.Post("/signup", userController.SignUp)
	router.Get("/users", userController.GetAll)
	router.Patch("/users/{id}", userController.Update)
	router.Post("/merchants", merchantController.Create)
	router.Get("/merchants/{id}", merchantController.Get)
	router.Patch("/merchants/{id}", merchantController.Update)
	router.Delete("/merchants/{id}", merchantController.Delete)
	router.Get("/merchants/{id}/members", merchantController.GetTeamMembers)
	router.Post("/merchants/{id}/members", merchantController.AddTeamMember)
	router.Delete("/merchants/{id}/members", merchantController.RemoveTeamMember)

	return router

}
