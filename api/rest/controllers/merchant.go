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

type Merchant interface {
	Create(w http.ResponseWriter, r *http.Request)
}

type merchant struct {
	cfg config.Configuration
	svc app.Services
}

func (ctlr *merchant) Create(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
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

func NewMerchant(cfg config.Configuration, svc app.Services) Merchant {
	return &merchant{
		svc: svc,
		cfg: cfg,
	}
}
