package util

import (
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
	"main/mocks"
	"testing"
)

type RandomStringGeneratorTestSuite struct {
	suite.Suite
	randomStringGenerator     RandomStringGenerator
	mockController            *gomock.Controller
	mockClock                 *mocks.MockClock
	mockRandomNumberGenerator *mocks.MockRandomNumberGenerator
}

func TestNewRandomStringGenerator(t *testing.T) {
	suite.Run(t, new(RandomStringGeneratorTestSuite))
}

func (suite *RandomStringGeneratorTestSuite) SetupTest() {
	suite.mockController = gomock.NewController(suite.T())
	suite.mockRandomNumberGenerator = mocks.NewMockRandomNumberGenerator(suite.mockController)
	suite.mockClock = mocks.NewMockClock(suite.mockController)
	suite.randomStringGenerator = NewRandomStringGenerator(suite.mockClock, suite.mockRandomNumberGenerator)
}

func (suite *RandomStringGeneratorTestSuite) TestGetRandString_ShouldReturnARandomStringForOneStringLength1() {
	nowInNanoSeconds := int64(454)
	suite.mockClock.EXPECT().NowUnixNano().Return(nowInNanoSeconds)
	suite.mockRandomNumberGenerator.EXPECT().NewRandomNumber(nowInNanoSeconds, 62).Return(12)

	result := suite.randomStringGenerator.GetRandString(1)

	suite.Equal("m", result)
}

func (suite *RandomStringGeneratorTestSuite) TestGetRandString_ShouldReturnARandomStringForStrignLengthMoreThan1() {
	nowInNanoSeconds1 := int64(454)
	nowInNanoSeconds2 := int64(655)
	suite.mockClock.EXPECT().NowUnixNano().Return(nowInNanoSeconds1).Times(1)
	suite.mockRandomNumberGenerator.EXPECT().NewRandomNumber(nowInNanoSeconds1, 62).Return(12).Times(1)
	suite.mockClock.EXPECT().NowUnixNano().Return(nowInNanoSeconds2).Times(1)
	suite.mockRandomNumberGenerator.EXPECT().NewRandomNumber(nowInNanoSeconds2, 62).Return(34).Times(1)

	result := suite.randomStringGenerator.GetRandString(2)

	suite.Equal("mI", result)
}
