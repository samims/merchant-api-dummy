package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/samims/merchant-api/app"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
	"github.com/samims/merchant-api/utils"
)

type User interface {
	SignUp(http.ResponseWriter, *http.Request)
	SignIn(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Update(http.ResponseWriter, *http.Request)
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

// SignIn controller is a controller that signs in a user by calling the service
func (ctlr *user) SignIn(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var loginModel models.LoginModel
	err = json.Unmarshal(body, &loginModel)

	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	errorList := utils.Validate(loginModel)
	if len(errorList) > 0 {
		logger.Log.Error(errorList)
		utils.Renderer(w, loginModel, errorList...)
		return
	}

	resp, err := ctlr.svc.UserService().SignIn(ctx, loginModel)
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

func (ctlr *user) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	userIdFromURL, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	userID, err := strconv.ParseInt(r.Header.Get("UserID"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}
	ctx = context.WithValue(ctx, constants.UserIDContextKey, userID)

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

	userPublic, err := ctlr.svc.UserService().Update(ctx, userIdFromURL, user)
	utils.Renderer(w, userPublic, err)

}

func NewUser(cfg config.Configuration, svc app.Services) User {
	return &user{
		cfg: cfg,
		svc: svc,
	}
}
