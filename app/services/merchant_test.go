package services

import (
	"context"
	"errors"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository/mocks"
	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/logger"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MerchantServiceTestSuite struct {
	suite.Suite
	ctx context.Context

	mockMerchantRepo *mocks.MerchantRepoMock
	userRepoMock     *mocks.UserRepoMock

	SUT MerchantService
}

// SetupTest is called before each test
func (suite *MerchantServiceTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.mockMerchantRepo = mocks.NewMerchantRepoMock(suite.T())
	suite.userRepoMock = mocks.NewUserRepoMock(suite.T())
	suite.SUT = NewMerchantService(suite.mockMerchantRepo, suite.userRepoMock)
}

func (suite *MerchantServiceTestSuite) TestVerifyObjectPermissionTestRaisesErrorIfUserDoesNotExist() {
	// Arrange
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1)},
	}
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1)},
	}
	suite.ctx = context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)
	// mock
	suite.userRepoMock.On("FindOne", suite.ctx, mock.Anything).Return(nil, errors.New(constants.UserNotFound))

	// Act
	err := suite.SUT.VerifyObjectPermission(suite.ctx, merchant.Id)

	// Assert
	suite.Error(err)
	suite.Equal(constants.UserNotFound, err.Error())

	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *MerchantServiceTestSuite) TestVerifyObjectPermissionTestRaisesErrorIfMerchantDoesNotExist() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1)},
	}

	user := models.User{
		BaseModel: models.BaseModel{Id: int64(1)},
		Merchant:  &merchant,
	}

	suite.ctx = context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// mock
	suite.userRepoMock.On("FindOne", suite.ctx, mock.Anything).Return(&user, nil)

	// Act
	err := suite.SUT.VerifyObjectPermission(suite.ctx, merchant.Id+1)

	// Assert
	suite.Error(err)
	logger.Log.Info(err.Error())
	suite.Equal(constants.PermissionDenied, err.Error())

	suite.mockMerchantRepo.AssertExpectations(suite.T())

}

func (suite *MerchantServiceTestSuite) TestMerchantCreateRaisesIfUserDoesnotExist() {
	// Arrange
	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, int64(1))
	merchant := models.Merchant{
		Name: fake.Company(),
		URL:  fake.DomainName(),
	}

	// mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(nil, errors.New(constants.UserNotFound))

	// Act
	result, err := suite.SUT.Create(ctx, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(constants.UserNotFound, err.Error())
	suite.Empty(result)

	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *MerchantServiceTestSuite) TestMerchantCreateRaisesIfUserAlreadyAttachedToMerchant() {
	// Arrange
	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, int64(1))
	merchant := models.Merchant{
		Name: fake.Company(),
		URL:  fake.DomainName(),
	}
	user := models.User{
		BaseModel: models.BaseModel{Id: int64(1)},
		Merchant:  &merchant,
	}

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)

	// Act
	result, err := suite.SUT.Create(ctx, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(constants.UserAlreadyPartOfAMerchant, err.Error())
	suite.Empty(result)

	suite.userRepoMock.AssertExpectations(suite.T())
}

func (suite *MerchantServiceTestSuite) TestMerchantCreateRaisesIfWriteToDBFailed() {
	// Arrange
	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, int64(1))
	merchant := models.Merchant{
		Name: fake.Company(),
		URL:  fake.DomainName(),
	}
	user := models.User{
		BaseModel: models.BaseModel{Id: int64(1)},
	}

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)

	// Statement Closed error of beego orm
	suite.mockMerchantRepo.On("Save", ctx, mock.Anything).Return(orm.ErrStmtClosed)

	// Act
	result, err := suite.SUT.Create(ctx, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrStmtClosed.Error(), err.Error())
	suite.Empty(result)

	suite.userRepoMock.AssertExpectations(suite.T())

}

func (suite *MerchantServiceTestSuite) TestMerchantCreateRaisesIfUpdateUserAfterMerchantSaveReturnsError() {
	// Arrange
	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, int64(1))
	merchant := models.Merchant{
		Name: fake.Company(),
		URL:  fake.DomainName(),
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("Save", ctx, mock.Anything).Return(nil)
	suite.userRepoMock.On("Update", ctx, mock.Anything, mock.Anything).Return(orm.ErrStmtClosed)
	suite.mockMerchantRepo.On("Delete", ctx, mock.Anything).Return(nil)

	// Act
	result, err := suite.SUT.Create(ctx, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrStmtClosed.Error(), err.Error())
	suite.Empty(result)

}

func TestMerchantService(t *testing.T) {
	suite.Run(t, new(MerchantServiceTestSuite))
}
