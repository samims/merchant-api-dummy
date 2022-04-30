package controllers

import (
	"net/http"

	"github.com/samims/merchant-api/config"
)

// Ping is the blueprint of ping controller
type Ping interface {
	Get(w http.ResponseWriter, r *http.Request)
}

// ping is the dependency struct of Ping interface
type ping struct {
	cfg config.Configuration
}

func (p *ping) Get(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Pong!!"))
}

func NewPing(cfg config.Configuration) Ping {
	return &ping{
		cfg: cfg,
	}
}
