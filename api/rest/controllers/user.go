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
	"github.com/samims/merchant-api/utils"
)

type User interface {
	SignUp(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
}

type user struct {
	cfg config.Configuration
	svc app.Services
}

// SignUp is a controller that creates a new user by calling the service
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

	errorList := utils.Validate(user)
	if len(errorList) > 0 {
		logger.Log.Error(errorList)
		utils.Renderer(w, user, errorList...)
		return
	}

	resp, err := ctlr.svc.UserService().SignUp(ctx, user)

	utils.Renderer(w, resp, err)

}

func (ctlr *user) GetAll(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	svc_resp, err := ctlr.svc.UserService().GetAll(ctx)

	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	userBytes, _ := json.Marshal(svc_resp)
	w.Write(userBytes)
}

func NewUser(cfg config.Configuration, svc app.Services) User {
	return &user{
		cfg: cfg,
		svc: svc,
	}
}
