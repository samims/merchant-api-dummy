package app

import (
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/app/services"
	"github.com/samims/merchant-api/config"
)

type Services interface {
	PingService() services.PingService
	UserService() services.UserService
}

type svc struct {
	pingService services.PingService
	userService services.UserService
}

func (svc *svc) PingService() services.PingService {
	return svc.pingService
}

func (svc *svc) UserService() services.UserService {
	return svc.userService
}

func InitServices(cfg config.Configuration) Services {
	db := cfg.PostgresConfig().GetDB()

	userRepo := repository.NewUserRepo(db)

	pingSvc := services.NewPingService()
	userSvc := services.NewUserService(userRepo)

	return &svc{
		pingService: pingSvc,
		userService: userSvc,
	}
}
