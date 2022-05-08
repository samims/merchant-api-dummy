package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/samims/merchant-api/api/rest/controllers"
	"github.com/samims/merchant-api/api/rest/middlewares"
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
	// router.Use(middlewares.IsAuthorized(cfg))

	router.Get("/ping", pingController.Get)
	// users crud apis
	router.Post("/signup", userController.SignUp)
	router.Post("/signin", userController.SignIn)

	router.Get("/users", middlewares.IsAuthorized(userController.GetAll, cfg))
	router.Patch("/users/{id}", userController.Update)

	// merchants crud apis
	router.Post("/merchants", middlewares.IsAuthorized(merchantController.Create, cfg))
	router.Get("/merchants/{id}", middlewares.IsAuthorized(merchantController.Get, cfg))
	router.Patch("/merchants/{id}", middlewares.IsAuthorized(merchantController.Update, cfg))
	router.Delete("/merchants/{id}", middlewares.IsAuthorized(merchantController.Delete, cfg))

	// merchant's teams apis
	router.Get("/merchants/{id}/members", middlewares.IsAuthorized(merchantController.GetTeamMembers, cfg))
	router.Post("/merchants/{id}/members", middlewares.IsAuthorized(merchantController.AddTeamMember, cfg))
	router.Delete("/merchants/{id}/members", middlewares.IsAuthorized(merchantController.RemoveTeamMember, cfg))

	return router

}
