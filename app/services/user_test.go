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

// TestGetAllSuccess tests that get all success in ideal case
func (suite *UserServiceTestSuite) TestGetAllSuccess() {
	// Arrange
	users := []models.User{
		suite.user,
	}

	// mock
	suite.mockUserRepo.On("GetAll", suite.ctx, mock.Anything).Return(users, int64(len(users)), nil)

	// Act
	result, err := suite.SUT.GetAll(suite.ctx)

	// Assert
	suite.NoError(err)
	suite.Equal(len(result), len(users))
}

// TestGetAllFailsForDatabaseError tests that get all fails for database error
func (suite *UserServiceTestSuite) TestGetAllFailsForDatabaseError() {
	// mock
	suite.mockUserRepo.On("GetAll", suite.ctx, mock.Anything).Return(nil, int64(0), orm.ErrStmtClosed)

	// Act
	result, err := suite.SUT.GetAll(suite.ctx)

	// Assert
	suite.Error(err)
	suite.Equal(len(result), 0)
	suite.Equal(err.Error(), orm.ErrStmtClosed.Error())
}

// TestUpdateSuccess tests that update success in ideal case
func (suite *UserServiceTestSuite) TestUpdateSuccess() {
	// mock
	suite.mockUserRepo.On("FindOne", suite.ctx, mock.Anything).Return(&suite.user, nil)
	suite.mockUserRepo.On("Update", suite.ctx, mock.Anything, mock.Anything).Return(nil)

	// Act
	result, err := suite.SUT.Update(suite.ctx, suite.user.Id, suite.user)

	// Assert
	suite.NoError(err)
	suite.Equal(result.ID, suite.user.Id)

}

// TestUpdateFailsIfUserNotFound tests that update fails if user not found
func (suite *UserServiceTestSuite) TestUpdateFailsIfUserNotFound() {
	// mock
	suite.mockUserRepo.On("FindOne", suite.ctx, mock.Anything).Return(nil, orm.ErrNoRows)

	// Act
	_, err := suite.SUT.Update(suite.ctx, suite.user.Id, suite.user)

	// Assert
	suite.Error(err)
	suite.Equal(err.Error(), orm.ErrNoRows.Error())
}

// TestUpdateFailsDuringWriteToDatabase tests that update fails during write to database
func (suite *UserServiceTestSuite) TestUpdateFailsDuringWriteToDatabase() {
	// mock
	suite.mockUserRepo.On("FindOne", suite.ctx, mock.Anything).Return(&suite.user, nil)
	suite.mockUserRepo.On("Update", suite.ctx, mock.Anything, mock.Anything).Return(orm.ErrStmtClosed)

	// Act
	_, err := suite.SUT.Update(suite.ctx, suite.user.Id, suite.user)

	// Assert
	suite.Error(err)
	suite.Equal(err.Error(), orm.ErrStmtClosed.Error())
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTestSuite))
}
