package utils

import (
	"errors"

	"github.com/google/uuid"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
	"golang.org/x/crypto/bcrypt"
)

// this file contains the cipher & decipher functions

// GenerateHash generates a hash from a string
func GenerateBCryptHash(s string) (string, error) {
	if len(s) == 0 {
		return "", errors.New(constants.ErrorEmptyString)
	}

	hashBytes, err := bcrypt.GenerateFromPassword([]byte(s), bcrypt.DefaultCost)
	if err != nil {
		logger.Log.WithError(err).Error("GenerateBCryptHash")
		return "", err
	}

	hashString := string(hashBytes)
	return hashString, nil
}

// ValidateBCryptHash validates a hash against a string
func ValidateBCryptHash(s, hashedString string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedString), []byte(s))
	if err != nil {
		logger.Log.WithError(err).Error("ValidateBCryptHash")
		return err
	}
	return nil
}

// GenerateUUIDString generates a UUID string
func GenerateUUIDString() (string, error) {
	generatedUUID, err := uuid.NewRandom()
	if err != nil {
		logger.Log.WithError(err).Error("GenerateUUIDString")
		return "", err
	}
	return generatedUUID.String(), nil

}
