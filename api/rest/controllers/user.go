package controllers

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/samims/merchant-api/app"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/logger"
)

type User interface {
	SignUp(http.ResponseWriter, *http.Request)
}

type user struct {
	cfg config.Configuration
	svc app.Services
}

func (ctlr *user) SignUp(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var user models.User
	err = json.Unmarshal(body, &user)

	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	svc_resp, err := ctlr.svc.UserService().SignUp(ctx, user)

	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	logger.Log.Info(svc_resp)
	userBytes, _ := json.Marshal(svc_resp)

	w.Write(userBytes)

}

func NewUser(cfg config.Configuration, svc app.Services) User {
	return &user{
		cfg: cfg,
		svc: svc,
	}
}
