package utils

import (
	"errors"
	"net/http"
	"testing"

	"github.com/samims/merchant-api/constants"
	"github.com/samims/merchant-api/utils/mocks"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type ResponseUtilsTestSuite struct {
	suite.Suite
	mockedResponseWriter *mocks.ResponseWriterMock
	mockedHeader         http.Header
}

func (suite *ResponseUtilsTestSuite) SetupTest() {
	suite.mockedResponseWriter = mocks.NewResponseWriterMock()
	suite.mockedHeader = GetMockedHeader()

}

// GetMockedHeader returns mocked header
func GetMockedHeader() http.Header {
	return map[string][]string{}
}

// TestRendererForSuccess tests that renderer for success
func (suite *ResponseUtilsTestSuite) TestRendererForSuccessShouldWriteData() {
	// Arrange
	data := map[string]interface{}{"test": "test"}

	// mock
	suite.mockedResponseWriter.On("Header").Return(suite.mockedHeader)
	suite.mockedResponseWriter.On("WriteHeader", mock.Anything).Return()
	suite.mockedResponseWriter.On("Write", mock.Anything).Return(0, nil)

	// Act
	Renderer(suite.mockedResponseWriter, data, nil)

	// Assert
	suite.mockedResponseWriter.AssertCalled(suite.T(), "Header")
	suite.mockedResponseWriter.AssertCalled(suite.T(), "WriteHeader", mock.Anything)
	suite.mockedResponseWriter.AssertCalled(suite.T(), "Write", mock.Anything)

}

// TestRendererForFailureShouldWriteError tests that renderer for failure should write error
func (suite *ResponseUtilsTestSuite) TestRendererForFailureShouldWriteError() {
	// Arrange
	data := map[string]interface{}{"test": "test"}
	err := errors.New(constants.ErorNameIsNotValid)

	// mock
	suite.mockedResponseWriter.On("Header").Return(suite.mockedHeader)
	suite.mockedResponseWriter.On("WriteHeader", mock.Anything).Return()
	suite.mockedResponseWriter.On("Write", mock.Anything).Return(0, nil)

	// Act
	Renderer(suite.mockedResponseWriter, data, err)

	// Assert
	suite.mockedResponseWriter.AssertCalled(suite.T(), "Header")
	suite.mockedResponseWriter.AssertCalled(suite.T(), "WriteHeader", mock.Anything)
	suite.mockedResponseWriter.AssertCalled(suite.T(), "Write", mock.Anything)
}

func TestResponseUtil(t *testing.T) {
	suite.Run(t, new(ResponseUtilsTestSuite))
}
