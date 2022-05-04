package utils

import (
	"fmt"

	"github.com/go-playground/validator/v10"
)

func Validate(doc interface{}) []error {
	var resp []error
	err := validator.New().Struct(doc)
	if err != nil {
		errorList := err.(validator.ValidationErrors)
		for _, e := range errorList {
			resp = append(resp, fmt.Errorf(`'%v' is not a valid %v`, e.Value(), e.StructField()))
		}
	}
	return resp
}
