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

func (suite *MerchantServiceTestSuite) TestMerchantServiceVerifyObjectPermissionTestRaisesErrorIfUserDoesNotExist() {
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

// TestVerifyObjectPermissionTestRaisesErrorIfMerchantDoesNotExist is a unit test for MerchantService.VerifyObjectPermission
// It verifies that the VerifyObjectPermission method raises an error if merchant's user is not same as current user
func (suite *MerchantServiceTestSuite) TestMerchantServiceVerifyObjectPermissionTestRaisesErrorIfMerchantDoesNotExist() {
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

// TestMerchantCreateRaisesIfUserDoesnotExist is a unit test for MerchantService.Create
// It verifies that the Create method raises an error if the user does not exist
func (suite *MerchantServiceTestSuite) TestMerchantServiceCreateRaisesIfUserDoesnotExist() {
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

// TestMerchantCreateRaisesIfUserAlreadyAttachedToMerchant is a unit test for MerchantService.Create
// It verifies that the Create method raises an error if the user is already attached to a merchant
func (suite *MerchantServiceTestSuite) TestMerchantServiceCreateRaisesIfUserAlreadyAttachedToMerchant() {
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

// TestMerchantCreateRaisesIfMerchantDataWriteToDBFailed is a unit test for MerchantService.Create
// It verifies that the Create method raises an error if the merchant data is not written to the DB
// due to an error in the DB layer
func (suite *MerchantServiceTestSuite) TestMerchantServiceCreateRaisesIfMerchantDataWriteToDBFailed() {
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

// TestMerchantCreateRaisesIfUpdateUserAfterMerchantSaveReturnsError tests if user attachment with merchant
// fails it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceCreateRaisesIfUpdateUserAfterMerchantSaveReturnsError() {
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

// TestMerchantCreateRaisesIfMerchantRollbackFails test if merchant creation is faild due to
// the cause that the user already has a merchant and then rollback is faild then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceCreateRaisesIfMerchantRollbackFails() {
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
	suite.mockMerchantRepo.On("Delete", ctx, mock.Anything).Return(orm.ErrArgs)

	// Act
	result, err := suite.SUT.Create(ctx, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrArgs.Error(), err.Error())
	suite.Empty(result)

}

// TestMerchantCreateSuccess tests if merchant creation is successful
// it should return PublicMerchant object and no error
func (suite *MerchantServiceTestSuite) TestMerchantServiceCreateSuccess() {
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
	suite.userRepoMock.On("Update", ctx, mock.Anything, mock.Anything).Return(nil)

	// Act
	result, err := suite.SUT.Create(ctx, merchant)

	// Assert
	suite.NoError(err)
	// assert that result is models.PublicMerchant
	suite.Equal(reflect.TypeOf(models.PublicMerchant{}), reflect.TypeOf(result))
	suite.Equal(merchant.Name, result.Name)

	suite.userRepoMock.AssertExpectations(suite.T())
}

// TestMerchantGetRaisesIfUserNotFound tests if user is not found then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceGetRaisesIfUserNotFound() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	// Mock
	suite.mockMerchantRepo.On("FindOne", suite.ctx, mock.Anything).Return(&merchant, orm.ErrNoRows)

	// Act
	result, err := suite.SUT.Get(suite.ctx, merchant.Id)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrNoRows.Error(), err.Error())
	suite.Equal(models.PublicMerchant{}, result)
	suite.Empty(result)

}

// TestMerchantGetSuccess tests if merchant is found then it should return PublicMerchant object and no error
func (suite *MerchantServiceTestSuite) TestMerchantServiceGetSuccess() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	// Mock
	suite.mockMerchantRepo.On("FindOne", suite.ctx, mock.Anything).Return(&merchant, nil)

	// Act
	result, err := suite.SUT.Get(suite.ctx, merchant.Id)

	// Assert
	suite.NoError(err)
	suite.Equal(reflect.TypeOf(models.PublicMerchant{}), reflect.TypeOf(result))

}

// TestMerchantServiceUpdateRaisesIfUserDoNotHavePermission tests if user do not have permission to update merchant
// then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceUpdateRaisesIfUserDoNotHavePermission() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)

	// Act
	result, err := suite.SUT.Update(ctx, merchant.Id+1, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(constants.PermissionDenied, err.Error())
	suite.Equal(models.PublicMerchant{}, result)

}

// TestMerchantServiceUpdateRaisesIfMerchantNotFound tests if merchant is not found then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceUpdateRaisesIfMerchantNotFound() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", ctx, mock.Anything).Return(nil, orm.ErrNoRows)

	// Act
	result, err := suite.SUT.Update(ctx, merchant.Id, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrNoRows.Error(), err.Error())
	suite.Equal(models.PublicMerchant{}, result)

	suite.userRepoMock.AssertExpectations(suite.T())

}

// TestMerchantServiceUpdateRaisesIfUpdateQueryFails tests if update query fails then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceUpdateRaisesIfUpdateQueryFails() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", ctx, mock.Anything).Return(&merchant, nil)
	suite.mockMerchantRepo.On("Update", ctx, mock.Anything, mock.Anything).Return(orm.ErrStmtClosed)

	// Act
	result, err := suite.SUT.Update(ctx, merchant.Id, merchant)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrStmtClosed.Error(), err.Error())
	suite.Equal(models.PublicMerchant{}, result)

	suite.userRepoMock.AssertExpectations(suite.T())
	suite.mockMerchantRepo.AssertExpectations(suite.T())
}

// TestMerchantServiceUpdateSuccess tests if merchant is found then it should return PublicMerchant object and no error
func (suite *MerchantServiceTestSuite) TestMerchantServiceUpdateSuccess() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", ctx, mock.Anything).Return(&merchant, nil)
	suite.mockMerchantRepo.On("Update", ctx, mock.Anything, mock.Anything).Return(nil)

	// Act
	result, err := suite.SUT.Update(ctx, merchant.Id, merchant)

	// Assert
	suite.NoError(err)
	suite.Equal(reflect.TypeOf(models.PublicMerchant{}), reflect.TypeOf(result))

	suite.userRepoMock.AssertExpectations(suite.T())
	suite.mockMerchantRepo.AssertExpectations(suite.T())
}

// TestMerchantServiceDeleteRaisesIfUserDoNotHavePermission tests if user do not have permission to delete merchant
// then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceDeleteRaisesIfUserDoNotHavePermission() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)

	// Act
	resp, err := suite.SUT.Delete(ctx, merchant.Id+1)

	// Assert
	suite.Error(err)
	suite.Equal(constants.PermissionDenied, err.Error())
	suite.Equal(map[string]interface{}{"success": false}, resp)

	suite.userRepoMock.AssertExpectations(suite.T())

}

// TestMerchantServiceDeleteRaisesIfDelete fails then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceDeleteRaisesIfDeleteFails() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", ctx, mock.Anything).Return(&merchant, nil)
	suite.mockMerchantRepo.On("Delete", ctx, mock.Anything).Return(orm.ErrStmtClosed)

	// Act
	resp, err := suite.SUT.Delete(ctx, merchant.Id)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrStmtClosed.Error(), err.Error())
	suite.Equal(map[string]interface{}{"success": false}, resp)

	suite.userRepoMock.AssertExpectations(suite.T())
	suite.mockMerchantRepo.AssertExpectations(suite.T())
}

// TestMerchantServiceDeleteSuccess tests if merchant is found then it should return success no error
func (suite *MerchantServiceTestSuite) TestMerchantServiceDeleteSuccess() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", ctx, mock.Anything).Return(&merchant, nil)
	suite.mockMerchantRepo.On("Delete", ctx, mock.Anything).Return(nil)

	// Act
	resp, err := suite.SUT.Delete(ctx, merchant.Id)

	// Assert
	suite.NoError(err)
	suite.Equal(map[string]interface{}{"success": true}, resp)

	suite.userRepoMock.AssertExpectations(suite.T())
	suite.mockMerchantRepo.AssertExpectations(suite.T())
}

// TestMerchantServiceRaisesPaginationError tests if non positive page number provided it sould return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceRaisesPaginationError() {
	// Arrange
	page := int64(-1)
	pageSize := int64(10)
	// Act
	_, err := suite.SUT.GetTeamMembers(suite.ctx, 1, &page, &pageSize)
	// Assert
	suite.Error(err)
	suite.Equal(constants.PaginationError, err.Error())

}

// TestMerchantServiceGetTeamMembersRaisesIfMerchantNotFound tests if merchant is not found then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceGetTeamMembersRaisesIfMerchantNotFound() {
	// Arrange
	page := int64(1)
	pageSize := int64(10)

	// Mock
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(nil, orm.ErrNoRows)
	// Act
	_, err := suite.SUT.GetTeamMembers(suite.ctx, 1, &page, &pageSize)
	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrNoRows.Error(), err.Error())

	suite.mockMerchantRepo.AssertExpectations(suite.T())
}

// TestMerchantServiceGetTeamMembersRaisesIfRelatedMerchantFetchFail tests if related merchant fetch fails then it should return error
func (suite *MerchantServiceTestSuite) TestMerchantServiceGetTeamMembersRaisesIfReltedUserFetchFail() {
	// Arrange
	page := int64(1)
	pageSize := int64(10)

	// Mock
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(&models.Merchant{}, nil)
	suite.userRepoMock.On("GetAll", mock.Anything, mock.Anything).Return([]models.User{}, int64(0), orm.ErrNoRows)

	// Act
	res, err := suite.SUT.GetTeamMembers(suite.ctx, 1, &page, &pageSize)
	// Assert
	suite.Error(err)

	suite.Equal(reflect.TypeOf(models.TeamMemberResponse{}), reflect.TypeOf(res))
	suite.Equal(res.Members, []models.PublicUser{})
	suite.Equal(orm.ErrNoRows.Error(), err.Error())

	suite.mockMerchantRepo.AssertExpectations(suite.T())
	suite.userRepoMock.AssertExpectations(suite.T())

}

// TestMerchantServiceGetTeamMembersSuccess tests if merchant is found then it should return success no error
func (suite *MerchantServiceTestSuite) TestMerchantServiceGetTeamMembersSuccess() {
	// Arrange
	page := int64(1)
	pageSize := int64(10)
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	user1 := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}
	user2 := models.User{
		BaseModel: models.BaseModel{
			Id: int64(2),
		},
		Merchant: &merchant,
	}

	users := []models.User{user1, user2}

	// Mock
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(&models.Merchant{}, nil)
	suite.userRepoMock.On("GetAll", mock.Anything, mock.Anything).Return(users, int64(0), nil)

	// Act
	res, err := suite.SUT.GetTeamMembers(suite.ctx, 1, &page, &pageSize)

	// Assert
	suite.NoError(err)
	suite.Equal(reflect.TypeOf(models.TeamMemberResponse{}), reflect.TypeOf(res))
	suite.Greaterf(len(res.Members), 0, "should have at least one member")

	suite.NotEmpty(res.Members)

	// check for pagination
	suite.NotEqual(res.CurrentPage, 0)
	suite.NotEqual(res.TotalRecord, 0)

	suite.mockMerchantRepo.AssertExpectations(suite.T())
	suite.userRepoMock.AssertExpectations(suite.T())
}

// TestAddTeamMemberRaisesForPermissionDenied tests if user fetch for authenticated user id fails then it should return error
func (suite *MerchantServiceTestSuite) TestAddTeamMemberRaisesIfAuthenticatedUserFetchOnVerifyPermissionFails() {
	// Arrange

	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&models.User{}, orm.ErrNoRows)

	// Act
	res, err := suite.SUT.AddTeamMember(ctx, merchant.Id, user.Id)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrNoRows.Error(), err.Error())
	suite.Equal(res["success"], false)

	suite.userRepoMock.AssertExpectations(suite.T())

}

// TestAddTeamMemberRaisesForPermissionDenied tests if user is not a merchant member then it should return error
func (suite *MerchantServiceTestSuite) TestAddTeamMemberRaisesForObjectLevelPermissionError() {
	// Arrange

	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)

	// Act
	res, err := suite.SUT.AddTeamMember(ctx, merchant.Id+1, user.Id)

	// Assert
	suite.Error(err)
	suite.Equal(err.Error(), constants.PermissionDenied)
	suite.Equal(res["success"], false)

	suite.userRepoMock.AssertExpectations(suite.T())

}

// TestAddTeamMemberRaisesIfMerchantFetchFails tests if merchant fetch fails then it should return error
func (suite *MerchantServiceTestSuite) TestAddTeamMemberRaisesIfMerchantFetchFails() {
	// Arrange

	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(&models.Merchant{}, orm.ErrNoRows)

	// Act
	res, err := suite.SUT.AddTeamMember(ctx, merchant.Id, user.Id)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrNoRows.Error(), err.Error())
	suite.Equal(res["success"], false)

	suite.mockMerchantRepo.AssertExpectations(suite.T())
	suite.userRepoMock.AssertExpectations(suite.T())
}

// TestAddTeamMemberRaisesIfUserAlreadyPartOfMerchant tests if user already part of a merchant then it should return error
func (suite *MerchantServiceTestSuite) TestAddTeamMemberRaisesIfUserAlreadyPartOfMerchant() {
	// Arrange

	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil)
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(&merchant, nil)

	// Act
	res, err := suite.SUT.AddTeamMember(ctx, merchant.Id, user.Id)

	// Assert
	suite.Error(err)
	suite.Equal(constants.UserAlreadyPartOfAMerchant, err.Error())
	suite.Equal(res["success"], false)

	suite.mockMerchantRepo.AssertExpectations(suite.T())
	suite.userRepoMock.AssertExpectations(suite.T())
}

// TestAddTeamMemberRaisesIfUserUpdateFails tests if user update fails when adding merchant to the new member(user)
func (suite *MerchantServiceTestSuite) TestAddTeamMemberRaisesIfUserUpdateFails() {
	// Arrange

	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	newMemberUser := models.User{
		BaseModel: models.BaseModel{
			Id: int64(2),
		},
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil).Once()
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(&merchant, nil)
	suite.userRepoMock.On("FindOne", mock.Anything, mock.Anything).Return(&newMemberUser, nil).Once()
	suite.userRepoMock.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(orm.ErrNoRows)

	// Act
	res, err := suite.SUT.AddTeamMember(ctx, merchant.Id, newMemberUser.Id)

	// Assert
	suite.Error(err)
	suite.Equal(orm.ErrNoRows.Error(), err.Error())
	suite.Equal(res["success"], false)

	suite.mockMerchantRepo.AssertExpectations(suite.T())
	suite.userRepoMock.AssertExpectations(suite.T())
}

// test success scenario
// TestAddTeamMemberSuccess tests if user is added to merchant team successfully then it should return success
func (suite *MerchantServiceTestSuite) TestAddTeamMemberSuccess() {
	// Arrange

	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}
	user := models.User{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
		Merchant: &merchant,
	}

	newMemberUser := models.User{
		BaseModel: models.BaseModel{
			Id: int64(2),
		},
	}

	ctx := context.WithValue(suite.ctx, constants.UserIDContextKey, user.Id)

	// Mock
	suite.userRepoMock.On("FindOne", ctx, mock.Anything).Return(&user, nil).Once()
	suite.mockMerchantRepo.On("FindOne", mock.Anything, mock.Anything).Return(&merchant, nil)
	suite.userRepoMock.On("FindOne", mock.Anything, mock.Anything).Return(&newMemberUser, nil).Once()
	suite.userRepoMock.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(nil)

	// Act
	res, err := suite.SUT.AddTeamMember(ctx, merchant.Id, newMemberUser.Id)

	// Assert
	suite.NoError(err)
	suite.Equal(res["success"], true)

	suite.mockMerchantRepo.AssertExpectations(suite.T())
	suite.userRepoMock.AssertExpectations(suite.T())
}

func TestMerchantService(t *testing.T) {
	suite.Run(t, new(MerchantServiceTestSuite))
}
