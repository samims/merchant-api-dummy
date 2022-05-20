package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/astaxie/beego/orm"
	"github.com/icrowley/fake"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type MerchantRepoTestSuite struct {
	suite.Suite
	ctx      context.Context
	mockedDB *mocks.OrmerMock
	mockedQS *mocks.QuerySeterMock
	SUT      MerchantRepo
}

// SetupTest is called before each test
func (suite *MerchantRepoTestSuite) SetupTest() {
	suite.ctx = context.Background()
	suite.mockedDB = mocks.NewOrmerMock(suite.T())
	suite.mockedQS = mocks.NewQuerySeterMock(suite.T())
	suite.SUT = NewMerchantRepo(suite.mockedDB)
}

// TestMerchantRepoSaveShouldReturnErrorIfWriteToDBFails tests that save should return error if write to DB fails
func (suite *MerchantRepoTestSuite) TestMerchantRepoSaveShouldReturnErrorIfWriteToDBFails() {
	// Arrange
	doc := &models.Merchant{
		Name: "",
	}

	// mock
	suite.mockedDB.On("Insert", mock.Anything).Return(int64(0), errors.New("Write to DB failed"))

	// Act
	err := suite.SUT.Save(suite.ctx, doc)

	// Assert
	suite.Error(err)
	// check dock has no attribute id
	suite.Equal(doc.Id, int64(0), "doc.Id should be 0 as it is not saved to DB")
}

// TestMerchantRepoSaveShouldReturnErrorIfDocIsNil tests that save should return error if doc is nil
func (suite *MerchantRepoTestSuite) TestMerchantRepoSaveShouldReturnErrorIfDocIsNil() {

	// Mock
	suite.mockedDB.On("Insert", mock.Anything).Return(int64(0), errors.New("Write to DB failed"))
	// Act
	err := suite.SUT.Save(suite.ctx, nil)

	// Assert
	suite.Error(err)
}

// TestMerchantRepoSaveSuccess tests that save should not return error if write to DB succeeds
func (suite *MerchantRepoTestSuite) TestMerchantRepoSaveSuccess() {
	// Arrange
	doc := &models.Merchant{
		Name: fake.Brand(),
	}

	// mock
	suite.mockedDB.On("Insert", mock.Anything).Return(int64(1), nil)

	// Act
	err := suite.SUT.Save(suite.ctx, doc)

	// Assert
	suite.NoError(err)
	suite.NotEqual(doc.Id, int64(0), "doc.Id should not be 0 as it is saved to DB")
}

// TestMerchantRepoUpdateShouldReturnErrorIfWriteToDBFails tests that update should return error if write to DB fails
func (suite *MerchantRepoTestSuite) TestMerchantRepoUpdateShouldReturnErrorIfWriteToDBFails() {
	// Arrange
	doc := &models.Merchant{
		BaseModel: models.BaseModel{Id: 1},
		Name:      fake.Brand(),
		URL:       fake.DomainName(),
	}

	// mock
	suite.mockedDB.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(int64(0), errors.New("Write to DB failed"))

	// Act
	err := suite.SUT.Update(suite.ctx, doc, []string{"name", "url"})

	// Assert
	suite.Error(err)
}

// TestMerchantRepoUpdateSuccess tests that update should not return error if write to DB succeeds
func (suite *MerchantRepoTestSuite) TestMerchantRepoUpdateSuccess() {
	// Arrange
	doc := models.Merchant{
		BaseModel: models.BaseModel{Id: int64(1)},
		Name:      fake.Brand(),
		URL:       fake.DomainName(),
	}

	// Mock
	suite.mockedDB.On("Update", mock.Anything, mock.Anything, mock.Anything).Return(int64(doc.Id), nil)

	// Act
	err := suite.SUT.Update(suite.ctx, &doc, []string{"name", "url"})

	// Assert
	suite.NoError(err)
}

// TestMerchantRepoDeleteReturnsErrorIfOrmRaises errors when deleting from DB
func (suite *MerchantRepoTestSuite) TestMerchantRepoDeleteReturnsErrorIfOrmRaisesErrorsWhenDeletingFromDB() {
	// Arrange
	doc := models.Merchant{
		BaseModel: models.BaseModel{Id: 1},
	}

	// Mock
	suite.mockedDB.On("Delete", mock.Anything).Return(int64(0), errors.New("Write to DB failed"))

	// Act
	err := suite.SUT.Delete(suite.ctx, doc)

	// Assert
	suite.Error(err)
}

// TestMerchantRepoDeleteSuccess tests that delete should not return error if write to DB succeeds
func (suite *MerchantRepoTestSuite) TestMerchantRepoDeleteSuccess() {
	// Arrange
	doc := models.Merchant{
		BaseModel: models.BaseModel{Id: 1},
	}

	// Mock
	suite.mockedDB.On("Delete", mock.Anything).Return(int64(1), nil)

	// Act
	err := suite.SUT.Delete(suite.ctx, doc)

	// Assert
	suite.NoError(err)
}

// // TestMerchantRepoGetALLShouldReturnErrorIfReadFromDBFails tests that getAll should return error if read from DB fails
func (suite *MerchantRepoTestSuite) TestMerchantRepoGetALLShouldReturnErrorIfReadFromDBFails() {

	// Mock
	suite.mockedDB.On("QueryTable", mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("OrderBy", mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("All", mock.Anything).Return(int64(0), orm.ErrNoRows)

	// Act
	res, err := suite.SUT.GetAll(suite.ctx)

	// Assert
	suite.Error(err)
	suite.Equal(err, orm.ErrNoRows, "err should be orm.ErrNoRows as mock returns orm.ErrNoRows")
	suite.IsType(res, []models.Merchant{}, "res should be []*models.Merchant")

	suite.mockedDB.AssertExpectations(suite.T())
	suite.mockedQS.AssertExpectations(suite.T())

}

// TestMerchantRepoGetALLSuccess tests that getAll success scenario
func (suite *MerchantRepoTestSuite) TestMerchantRepoGetALLSuccess() {

	// Mock
	suite.mockedDB.On("QueryTable", mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("OrderBy", mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("All", mock.Anything).Return(int64(1), nil).Run(
		func(args mock.Arguments) {
			// declare the var with type
			arg := args.Get(0).(*[]models.Merchant)
			// assign the value to the pointer
			*arg = []models.Merchant{
				{
					BaseModel: models.BaseModel{Id: int64(1)},
				},
			}
		},
	)

	// Act
	res, err := suite.SUT.GetAll(suite.ctx)

	// Assert
	suite.NoError(err)
	suite.IsType([]models.Merchant{}, res, "res should be []models.Merchant")
	suite.Equal(1, len(res), "res should have 1 element")

	suite.mockedDB.AssertExpectations(suite.T())
	suite.mockedQS.AssertExpectations(suite.T())
}

// TestMerchantRepoFindOneReadFromDBFails should return error if read from DB fails
func (suite *MerchantRepoTestSuite) TestMerchantRepoFindOneReadFromDBFails() {
	// arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	// Mock
	suite.mockedDB.On("QueryTable", mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("Filter", mock.Anything, mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("One", mock.Anything).Return(orm.ErrNoRows)

	// Act
	res, err := suite.SUT.FindOne(suite.ctx, merchant)

	// Assert
	suite.Error(err)
	suite.Empty(res)

}

func (suite *MerchantRepoTestSuite) TestMerhcantRepoFindOneTestSuccess() {
	// Arrange
	merchant := models.Merchant{
		BaseModel: models.BaseModel{
			Id: int64(1),
		},
	}

	// Mock
	suite.mockedDB.On("QueryTable", mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("Filter", mock.Anything, mock.Anything).Return(suite.mockedQS)
	suite.mockedQS.On("One", mock.Anything).Return(nil).Run(
		func(args mock.Arguments) {
			arg := args.Get(0).(*models.Merchant)
			*arg = merchant
		},
	)

	// Act
	res, err := suite.SUT.FindOne(suite.ctx, merchant)

	// Assert
	suite.NoError(err)
	suite.Equal(&merchant, res)

}

func TestUserService(t *testing.T) {
	suite.Run(t, new(MerchantRepoTestSuite))
}
