package controller

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"main/mocks"
	"main/model"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UrlShortnerControllerTestSuite struct {
	suite.Suite
	context                *gin.Context
	recorder               *httptest.ResponseRecorder
	mockController         *gomock.Controller
	mockUrlShortnerService *mocks.MockUrlShortnerService
	urlShortnerController  UrlShortnerController
}

func TestUrlShortnerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UrlShortnerControllerTestSuite))
}

func (suite *UrlShortnerControllerTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.mockUrlShortnerService = mocks.NewMockUrlShortnerService(suite.mockController)
	suite.urlShortnerController = NewUrlShortnerController(suite.mockUrlShortnerService)
}

func (suite *UrlShortnerControllerTestSuite) TestUrlShortner_Shorten_OnSuccess() {
	request := model.ShortenUrlRequestModel{Url: "www.google.com"}
	requestInBytes, _ := json.Marshal(request)
	testUrl := model.ShortenResponseModel{ShortenedUrl: "/1231wE"}
	expectedResponse, _ := json.Marshal(testUrl)

	suite.context.Request, _ = http.NewRequest("GET", "", bytes.NewBufferString(string(requestInBytes)))
	suite.mockUrlShortnerService.EXPECT().Shorten().Return(testUrl)

	suite.urlShortnerController.Shorten(suite.context)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
	assert.Equal(suite.T(), string(expectedResponse), suite.recorder.Body.String())
}

func (suite *UrlShortnerControllerTestSuite) TestUrlShortner_Shorten_OnError_WhenRequestIsMissingBody() {

	suite.context.Request, _ = http.NewRequest("GET", "", http.NoBody)

	suite.urlShortnerController.Shorten(suite.context)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}

func (suite *UrlShortnerControllerTestSuite) TestUrlShortner_Shorten_OnError_WhenUrlIsEmpty() {
	request := model.ShortenUrlRequestModel{Url: ""}
	requestInBytes, _ := json.Marshal(request)

	suite.context.Request, _ = http.NewRequest("GET", "", bytes.NewBufferString(string(requestInBytes)))

	suite.urlShortnerController.Shorten(suite.context)

	assert.Equal(suite.T(), http.StatusBadRequest, suite.recorder.Code)
}