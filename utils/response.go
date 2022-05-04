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
func Renderer(w http.ResponseWriter, data interface{}, err error) {
	groupError := "Renderer"
	// success response handling
	if err == nil {
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

	// error handling for the response
	errMsg := constants.ErrorString[err.Error()]
	w.Header().Add("Content-Type", "application/json")
	w.Header().Add("Message", errMsg)
	w.WriteHeader(constants.ErrorCode[err.Error()])

	resp := ErrorResponse{
		Message: errMsg,
		Details: err.Error(),
	}

	jsonResp, err := json.Marshal(resp)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Write(jsonResp)

}
