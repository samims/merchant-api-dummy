package utils

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type CipherUtilTestSuite struct {
	suite.Suite
}

func (suite *CipherUtilTestSuite) SetupTest() {
}

// TestGenerateBcryptHashSuccess tests that generate bcrypt hash
func (suite *CipherUtilTestSuite) TestGenerateBcryptHashSuccess() {
	// Arrange
	s := "test"

	// Act
	hash, err := GenerateBCryptHash(s)

	// Assert
	suite.NoError(err)
	suite.NotEmpty(hash)
	suite.NoError(ValidateBCryptHash(s, hash))
}

// TestGenerateBcryptHasFailsForEmptyString tests that generate bcrypt hash fails for empty string
func (suite *CipherUtilTestSuite) TestGenerateBcryptHashFailsForEmptyString() {
	// Arrange
	s := ""

	// Act
	hash, err := GenerateBCryptHash(s)

	// Assert
	suite.Error(err)
	suite.Empty(hash)
}

// TestValidateBcryptHashSuccess tests that validate bcrypt hash
func (suite *CipherUtilTestSuite) TestValidateBcryptHashSuccess() {
	// Arrange
	s := "test"
	hash, err := GenerateBCryptHash(s)
	suite.NoError(err)
	suite.NotEmpty(hash)

	// Act
	err = ValidateBCryptHash(s, hash)

	// Assert
	suite.NoError(err)
}

// TestValidateBcryptHashFailsForUnmathedHash tests that validate bcrypt hash fails when hash value of string is not correct
func (suite *CipherUtilTestSuite) TestValidateBcryptHashFailsForUnmatchedHash() {
	// Arrange
	s := "test"
	hash, err := GenerateBCryptHash(s)
	suite.NoError(err)
	suite.NotEmpty(hash)

	// Act
	err = ValidateBCryptHash(s, "wrong hash")

	// Assert
	suite.Error(err)
}

// TestGenerateUUIDStringSuccess tests that it generates uuid string
func (suite *CipherUtilTestSuite) TestGenerateUUIDStringSuccess() {
	// Act
	uuid, err := GenerateUUIDString()

	// Assert
	suite.NotEmpty(uuid)
	suite.NoError(err)
}

func TestCipherUtil(t *testing.T) {
	suite.Run(t, new(CipherUtilTestSuite))
}
