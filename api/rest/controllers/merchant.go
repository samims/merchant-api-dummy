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

type Merchant interface {
	Create(w http.ResponseWriter, r *http.Request)
	Get(w http.ResponseWriter, r *http.Request)
	Update(w http.ResponseWriter, r *http.Request)
	Delete(w http.ResponseWriter, r *http.Request)
	GetTeamMembers(w http.ResponseWriter, r *http.Request)
	AddTeamMember(w http.ResponseWriter, r *http.Request)
	RemoveTeamMember(w http.ResponseWriter, r *http.Request)
}

type merchant struct {
	cfg config.Configuration
	svc app.Services
}

func (ctlr *merchant) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// get UserID from header
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

	var merchant models.Merchant
	err = json.Unmarshal(body, &merchant)

	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	validationErrors := utils.Validate(merchant)
	if len(validationErrors) > 0 {
		logger.Log.Error(validationErrors)
		utils.Renderer(w, merchant, validationErrors...)
		return
	}
	resp, err := ctlr.svc.MerchantService().Create(ctx, merchant)
	utils.Renderer(w, resp, err)

}

func (ctrl *merchant) Get(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	merchantId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)

	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}

	merchantPublic, err := ctrl.svc.MerchantService().Get(ctx, merchantId)
	utils.Renderer(w, merchantPublic, err)

}

func (ctrl *merchant) Update(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()

	// get UserID from header
	userID, err := strconv.ParseInt(r.Header.Get("UserID"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}
	ctx = context.WithValue(ctx, constants.UserIDContextKey, userID)

	merhcantID, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, err)
	}

	var merchant models.Merchant
	err = json.Unmarshal(body, &merchant)

	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.InternalServerError))
		return
	}
	merchantPublic, err := ctrl.svc.MerchantService().Update(ctx, merhcantID, merchant)
	utils.Renderer(w, merchantPublic, err)
}

func (ctrl *merchant) Delete(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// get UserID from header
	userID, err := strconv.ParseInt(r.Header.Get("UserID"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}
	ctx = context.WithValue(ctx, constants.UserIDContextKey, userID)

	merchantId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, err)
		return
	}
	resp, err := ctrl.svc.MerchantService().Delete(ctx, merchantId)
	utils.Renderer(w, resp, err)
}

func (ctrl *merchant) GetTeamMembers(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	merchantId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}

	// pointer variable created to make page and limit as optional param
	var page, pageSize *int64
	pageNum, err := strconv.ParseInt(r.URL.Query().Get("page"), 10, 64)
	if err == nil {
		page = &pageNum
	}

	pageSizeNum, err := strconv.ParseInt(r.URL.Query().Get("size"), 10, 64)
	if err == nil {
		pageSize = &pageSizeNum
	}

	merchantPublic, err := ctrl.svc.MerchantService().GetTeamMembers(ctx, merchantId, page, pageSize)
	utils.Renderer(w, merchantPublic, err)
}

func (ctrl *merchant) AddTeamMember(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// get UserID from header
	userID, err := strconv.ParseInt(r.Header.Get("UserID"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}
	ctx = context.WithValue(ctx, constants.UserIDContextKey, userID)

	merchantId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var bodyMap map[string]int64
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	merchant, err := ctrl.svc.MerchantService().AddTeamMember(ctx, merchantId, bodyMap["user_id"])
	utils.Renderer(w, merchant, err)
}

func (ctrl *merchant) RemoveTeamMember(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	// get UserID from header
	userID, err := strconv.ParseInt(r.Header.Get("UserID"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}
	ctx = context.WithValue(ctx, constants.UserIDContextKey, userID)

	merchantId, err := strconv.ParseInt(chi.URLParam(r, "id"), 10, 64)
	if err != nil {
		logger.Log.Error(err)
		utils.Renderer(w, nil, errors.New(constants.BadRequest))
		return
	}

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var bodyMap map[string]int64
	err = json.Unmarshal(body, &bodyMap)
	if err != nil {
		logger.Log.Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	merchant, err := ctrl.svc.MerchantService().RemoveTeamMember(ctx, merchantId, bodyMap["user_id"])
	utils.Renderer(w, merchant, err)
}

func NewMerchant(cfg config.Configuration, svc app.Services) Merchant {
	return &merchant{
		svc: svc,
		cfg: cfg,
	}
}
