package app

import (
	"github.com/samims/merchant-api/app/repository"
	"github.com/samims/merchant-api/app/services"
	"github.com/samims/merchant-api/config"
)

type Services interface {
	PingService() services.PingService
	UserService() services.UserService
	MerchantService() services.MerchantService
}

type svc struct {
	pingService     services.PingService
	userService     services.UserService
	merchantService services.MerchantService
}

func (svc *svc) PingService() services.PingService {
	return svc.pingService
}

func (svc *svc) UserService() services.UserService {
	return svc.userService
}

func (svc *svc) MerchantService() services.MerchantService {
	return svc.merchantService

}

func InitServices(cfg config.Configuration) Services {
	db := cfg.PostgresConfig().GetDB()

	userRepo := repository.NewUserRepo(db)
	merchantRepo := repository.NewMerchantRepo(db)

	pingSvc := services.NewPingService()
	userSvc := services.NewUserService(userRepo, cfg)
	merchantSvc := services.NewMerchantService(merchantRepo, userRepo)

	return &svc{
		pingService:     pingSvc,
		userService:     userSvc,
		merchantService: merchantSvc,
	}
}
