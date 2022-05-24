package repository

import (
	"context"
	"errors"
	"testing"

	"github.com/icrowley/fake"
	"github.com/samims/merchant-api/app/models"
	"github.com/samims/merchant-api/app/repository/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type UserRepoTestSuite struct {
	suite.Suite
	ctx        context.Context
	mockedDB   *mocks.OrmerMock
	mockedQS   *mocks.QuerySeterMock
	pagination *models.Pagination
	SUT        UserRepo
}

func (suite *UserRepoTestSuite) SetupTest() {
	page := int64(1)
	size := int64(10)
	suite.ctx = context.Background()
	suite.mockedDB = mocks.NewOrmerMock(suite.T())
	suite.mockedQS = mocks.NewQuerySeterMock(suite.T())
	suite.pagination = &models.Pagination{
		Page: &page,
		Size: &size,
	}
	suite.SUT = NewUserRepo(suite.mockedDB)
}

func (suite *UserRepoTestSuite) TestUserRepoSaveShouldReturnErrorIfWriteToDBFails() {
	// Arrange
	doc := models.User{
		FirstName: fake.FirstName(),
	}

	// mock
	suite.mockedDB.On("Insert", mock.Anything).Return(int64(0), errors.New("Write to DB failed"))

	// Act
	err := suite.SUT.Save(suite.ctx, &doc)

	// Assert
	suite.Error(err)
	// check dock has no attribute id
	suite.Equal(doc.Id, int64(0), "doc.Id should be 0 as it is not saved to DB")
}

// TestUserRepoSaveSuccess tests that save user to DB should succeed
func (suite *UserRepoTestSuite) TestUserRepoSaveSuccess() {
	// Arrange
	doc := models.User{
		FirstName: fake.FirstName(),
	}

	// mock
	suite.mockedDB.On("Insert", mock.Anything).Return(int64(1), nil)

	// Act
	err := suite.SUT.Save(suite.ctx, &doc)

	// Assert
	suite.NoError(err)
	suite.Equal(doc.Id, int64(1), "doc.Id should be 1 as it is saved to DB")
}

// TestUserRepoGetAllShouldReturnErrorIfReadFromDBFails tests if read from db fails then error should returned
// func (suite *UserRepoTestSuite) TestUserRepoGetAllShouldReturnErrorIfReadFromDBFails() {
// 	// Arrange
// 	userQuery := models.UserQuery{
// 		User:       models.User{},
// 		Pagination: suite.pagination,
// 	}
// 	// mock
// 	suite.mockedDB.On("QueryTable", mock.Anything).Return(suite.mockedQS)
// 	suite.mockedQS.On("OrderBy", mock.Anything).Return(suite.mockedQS)
// 	suite.mockedQS.On("Offset", mock.Anything).Return(suite.mockedQS)
// 	suite.mockedQS.On("Limit", mock.Anything).Return(suite.mockedQS)
// 	// suite.mockedQS.On("Filter", mock.Anything, mock.Anything).Return(suite.mockedQS)
// 	suite.mockedQS.On("RelatedSel").Return(suite.mockedQS)
// 	suite.mockedQS.On("All", mock.Anything).Return(nil, errors.New("Read from DB failed"))

// 	// Act
// 	_, _, err := suite.SUT.GetAll(suite.ctx, userQuery)

// 	// Assert
// 	suite.Error(err)

// }

// Run the test suite
func TestUserRepoTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepoTestSuite))
}
