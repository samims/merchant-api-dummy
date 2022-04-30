package app

import (
	"github.com/samims/merchant-api/app/services"
	"github.com/samims/merchant-api/config"
)

type Services interface {
	PingServices() services.Ping
}

type svc struct {
}

func (svc *svc) PingServices() services.Ping {
	ping := services.NewPing()
	return ping
}

func InitServices(cfg config.Configuration) Services {
	return &svc{}
}
