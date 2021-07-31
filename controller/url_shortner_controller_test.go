package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"net/http"
	"net/http/httptest"
	"testing"
)

type UrlShortnerControllerTestSuite struct {
	suite.Suite
	context *gin.Context
	recorder *httptest.ResponseRecorder
	urlShortnerController UrlShortnerController
}

func TestUrlShortnerControllerTestSuite(t *testing.T) {
	suite.Run(t, new(UrlShortnerControllerTestSuite))
}

func (suite *UrlShortnerControllerTestSuite) SetupTest() {
	suite.recorder = httptest.NewRecorder()
	suite.context, _ = gin.CreateTestContext(suite.recorder)
	suite.urlShortnerController = NewUrlShortnerController()
}

func  (suite *UrlShortnerControllerTestSuite) TestUrlShortner_Shorten_OnSuccess() {
	suite.context.Request, _ = http.NewRequest("GET", "", http.NoBody)

	suite.urlShortnerController.Shorten(suite.context)

	assert.Equal(suite.T(), http.StatusOK, suite.recorder.Code)
}