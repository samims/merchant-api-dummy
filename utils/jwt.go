package utils

import (
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
)

func GenerateJWT(id int64, email, secretKey string) (string, error) {
	var mySigningKey = []byte(secretKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["email"] = email
	claims["user_id"] = id
	claims["exp"] = time.Now().Add(time.Minute * constants.JWT_EXPIRATION_DELTA_IN_MINUTES).Unix()

	// signed token
	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		logger.Log.Error("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}
