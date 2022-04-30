package controllers

import (
	"net/http"

	"github.com/samims/merchant-api/app"
	"github.com/samims/merchant-api/config"
)

// Ping is the blueprint of ping controller
type Ping interface {
	Get(w http.ResponseWriter, r *http.Request)
}

// ping is the dependency struct of Ping interface
type ping struct {
	cfg config.Configuration
	svc app.Services
}

func (p *ping) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(p.svc.PingServices().Get()))
}

func NewPing(cfg config.Configuration, svc app.Services) Ping {
	return &ping{
		cfg: cfg,
		svc: svc,
	}
}
