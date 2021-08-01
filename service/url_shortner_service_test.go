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
	mockCacheUtil             *mocks.MockCacheUtil
}

func TestNewUrlShortnerService(t *testing.T) {
	suite.Run(t, new(UrlShortnerServiceTestSuite))
}

func (suite *UrlShortnerServiceTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRandomStringGenerator = mocks.NewMockRandomStringGenerator(suite.mockController)
	suite.mockCacheUtil = mocks.NewMockCacheUtil(suite.mockController)
	suite.urlShortnerService = NewUrlShortnerService(suite.mockRandomStringGenerator, suite.mockCacheUtil)
}

func (suite *UrlShortnerServiceTestSuite) TestShorten_ShouldReturnRandomStringOfLength6() {
	suite.mockRandomStringGenerator.EXPECT().GetRandString(6).Return("adf3Lk")
	suite.mockCacheUtil.EXPECT().Get("www.google.com").Return(nil, false)
	suite.mockCacheUtil.EXPECT().Set("www.google.com", "http://localhost:8080/adf3Lk").Return()
	suite.mockCacheUtil.EXPECT().Set("http://localhost:8080/adf3Lk", "www.google.com").Return()

	result := suite.urlShortnerService.Shorten("www.google.com")

	suite.Equal("http://localhost:8080/adf3Lk", result.ShortenedUrl)
}

func (suite *UrlShortnerServiceTestSuite) TestFetch_ShouldReturnAlreadyStoredDataInCache() {
	suite.mockCacheUtil.EXPECT().Get("http://localhost:8080/34asd3").Return("www.google.com", true)

	fetch, err := suite.urlShortnerService.Fetch("http://localhost:8080/34asd3")

	suite.Nil(err)
	suite.Equal("www.google.com", fetch)
}

func (suite *UrlShortnerServiceTestSuite) TestFetch_ShouldReturnErrorWhenNoStoredDataInCache() {
	suite.mockCacheUtil.EXPECT().Get("http://localhost:8080/34asd3").Return(nil, false)

	fetch, err := suite.urlShortnerService.Fetch("http://localhost:8080/34asd3")

	suite.NotNil(err)
	suite.Equal("", fetch)
}