package services

import (
	"context"
	"errors"
	"reflect"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository/mocks"
	"github.com/samims/merchant-api/constants"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserServiceTestSuite struct {
	suite.Suite
	ctx context.Context

	mockUserRepo *mocks.UserRepoMock

	// stub
	user models.User
	SUT  UserService
}

// SetupTest is called before each test
func (suite *UserServiceTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.mockUserRepo = mocks.NewUserRepoMock(suite.T())

	suite.user = models.User{
		Email:        fake.EmailAddress(),
		FirstName:    fake.FirstName(),
		LastName:     fake.LastName(),
		PasswordHash: fake.CharactersN(10),
	}

	suite.SUT = NewUserService(suite.mockUserRepo)

}

// TestSignUpWithEmptyEmailShouldFail tests that sign up with empty email should fail
func (suite *UserServiceTestSuite) TestSignUpSuccess() {
	// Arrange

	// mock
	suite.mockUserRepo.On("Save", suite.ctx, mock.Anything).Return(nil).Run(func(args mock.Arguments) {
		user := args.Get(1).(*models.User)
		user.Id = int64(1)
	})

	// Act
	result, err := suite.SUT.SignUp(suite.ctx, suite.user)

	// Assert
	suite.NoError(err)
	typeResult := reflect.TypeOf(result)
	suite.Equal(typeResult, reflect.TypeOf(models.PublicUser{}))
	suite.Equal(result.ID, int64(1))

}

// TestEmailNeedToBeUnique tests that email need to be unique
func (suite *UserServiceTestSuite) TestEmailNeedToBeUnique() {
	// mock
	suite.mockUserRepo.On("Save", suite.ctx, mock.Anything).Return(errors.New(constants.UniqueEmailError))

	// Act
	result, err := suite.SUT.SignUp(suite.ctx, suite.user)

	// Assert
	suite.Error(err)
	suite.Equal(result.ID, int64(0))
	suite.Equal(err.Error(), constants.UniqueEmailError)

}

// TestSignUpFailsForPasswordHasGenerationError tests that sign up fails for password has generation error
func (suite *UserServiceTestSuite) TestSignUpFailsForPasswordHasGenerationError() {
	// Arrange
	suite.user.PasswordHash = ""

	// Act
	result, err := suite.SUT.SignUp(suite.ctx, suite.user)

	// Assert
	suite.Error(err)
	suite.Equal(result.ID, int64(0))
	suite.Equal(err.Error(), constants.ErrorEmptyString)

}

// TestSignUpFailsForDatabaseError tests that sign up fails for database error
func (suite *UserServiceTestSuite) TestSignUpFailsForDatabaseError() {
	// mock
	suite.mockUserRepo.On("Save", suite.ctx, mock.Anything).Return(orm.ErrStmtClosed)

	// Act
	result, err := suite.SUT.SignUp(suite.ctx, suite.user)

	// Assert
	suite.Error(err)
	suite.Equal(result.ID, int64(0))
	suite.Equal(err.Error(), orm.ErrStmtClosed.Error())

}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
