package middlewares

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/logger"
)

//set error message in Error struct
// func SetError(err Error, message string) Error {
// 	err.IsError = true
// 	err.Message = message
// 	return err
// }

func IsAuthorized(handler http.HandlerFunc, cfg config.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Header["Authorization"] == nil {

			err := errors.New("no token found")
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
			logger.Log.Error(err)
			return
		}

		var mySigningKey = []byte(cfg.AppConfig().GetSecretKey())
		authorizationToken := strings.Split(r.Header["Authorization"][0], " ")[1]

		token, err := jwt.Parse(authorizationToken, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("parsing error")
			}
			return mySigningKey, nil
		})
		logger.Log.Error(err)

		if err != nil {
			err := errors.New("unauthorized")
			json.NewEncoder(w).Encode(map[string]interface{}{"error": err.Error()})
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// get user id from token
			userID := int64(claims["user_id"].(float64))
			email := claims["email"].(string)

			r.Header.Set("UserID", fmt.Sprintf("%d", userID))
			r.Header.Set("Email", email)
			handler.ServeHTTP(w, r)
			return

			// if claims["role"] == "admin" {

			// 	r.Header.Set("Role", "admin")
			// 	handler.ServeHTTP(w, r)
			// 	return

			// } else if claims["role"] == "user" {

			// 	r.Header.Set("Role", "user")
			// 	handler.ServeHTTP(w, r)
			// 	return
			// }
		}
		err = errors.New("not authorized")
		json.NewEncoder(w).Encode(err)
	}
}
