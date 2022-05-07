package utils

import (
	"encoding/json"
	"net/http"

	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
)

// ErrorResponse is a struct that contains the error message model for the response
type ErrorResponse struct {
	Message string `json:"message"`
	Details string `json:"details"`
}

// Renderer is a function that handles the success response
func Renderer(w http.ResponseWriter, data interface{}, errList ...error) {
	logger.Log.Info(errList)
	groupError := "Renderer"
	// error handling for the response

	w.Header().Add("Content-Type", "application/json")
	erroResp := []ErrorResponse{}
	for _, err := range errList {

		if err == nil {
			continue
		}
		errMsg, ok := constants.ErrorString[err.Error()]
		if !ok {
			errMsg = err.Error()
		}

		resp := ErrorResponse{
			Message: errMsg,
			Details: err.Error(),
		}
		erroResp = append(erroResp, resp)

	}
	if len(erroResp) == 0 {
		w.Header().Add("Content-Type", "application/json")

		jsonResp, err := json.Marshal(data)
		if err != nil {
			logger.Log.WithError(err).Error(groupError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(jsonResp)
		return
	}
	w.WriteHeader(http.StatusBadRequest)

	jsonResp, err := json.Marshal(erroResp)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)

}
