package middlewares

import (
	"errors"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/samims/merchant-api/config"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/utils"
)

func IsAuthorized(handler http.HandlerFunc, cfg config.Configuration) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// when request does not have authorization header
		if r.Header["Authorization"] == nil {
			err := errors.New(constants.ErrorTokenNotFound)
			utils.Renderer(w, nil, err)
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

		// when token is invalid
		if err != nil {
			err := errors.New(constants.ErrorInvalidAuthToken)
			utils.Renderer(w, nil, err)
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
		}
		err = errors.New(constants.Unauthorized)
		utils.Renderer(w, nil, err)
	}
}
