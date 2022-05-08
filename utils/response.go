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

// Renderer is a function that renders the response for both success and error
func Renderer(w http.ResponseWriter, data interface{}, errList ...error) {
	groupError := "Renderer"
	w.Header().Add("Content-Type", "application/json")

	statusCode := http.StatusOK

	erroResp := []ErrorResponse{}
	for index, err := range errList {

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

		// need to check only for the first error in the list, repeted occurances will be ignored
		if index == 0 {
			statusCode, ok = constants.ErrorCode[err.Error()]
			// if not found, use the default status code
			if !ok {
				statusCode = http.StatusInternalServerError
			}
		}
		erroResp = append(erroResp, resp)

	}
	w.WriteHeader(statusCode)
	if len(erroResp) == 0 {

		jsonResp, err := json.Marshal(data)
		if err != nil {
			logger.Log.WithError(err).Error(groupError)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// for success resp
		w.Write(jsonResp)
		return
	}

	// error response
	jsonResp, err := json.Marshal(erroResp)
	if err != nil {
		logger.Log.WithError(err).Error(groupError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// write the error response body
	w.Write(jsonResp)

}
