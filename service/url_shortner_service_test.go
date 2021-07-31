package service

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"main/mocks"
	"testing"
)

type UrlShortnerServiceTestSuite struct {
	suite.Suite
	urlShortnerService        UrlShortnerService
	mockController            *gomock.Controller
	mockRandomStringGenerator *mocks.MockRandomStringGenerator
}

func TestNewUrlShortnerService(t *testing.T) {
	suite.Run(t, new(UrlShortnerServiceTestSuite))
}

func (suite *UrlShortnerServiceTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRandomStringGenerator = mocks.NewMockRandomStringGenerator(suite.mockController)
	suite.urlShortnerService = NewUrlShortnerService(suite.mockRandomStringGenerator)
}

func (suite *UrlShortnerServiceTestSuite) TestShorten_ShouldReturnRandomStringOfLength6() {
	suite.mockRandomStringGenerator.EXPECT().GetRandString(6).Return("adf3Lk")

	result := suite.urlShortnerService.Shorten()

	suite.Equal("/adf3Lk", result.ShortenedUrl)
}
