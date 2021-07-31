package util

import "fmt"

//go:generate mockgen -source=random_string_generator.go -destination=../mocks/mock_random_string_generator.go -package=mocks

type RandomStringGenerator interface {
	GetRandString(int) string
}

type randomStringGenerator struct {
	clock                 Clock
	randomNumberGenerator RandomNumberGenerator
}

func NewRandomStringGenerator(clock Clock, randomNumberGenerator RandomNumberGenerator) RandomStringGenerator {
	return randomStringGenerator{clock: clock, randomNumberGenerator: randomNumberGenerator}
}

func (r randomStringGenerator) GetRandString(length int) string {
	fmt.Println("getting random string")
	return r.stringWithCharset(length, charset)
}

const charset = "abcdefghijklmnopqrstuvwxyz" +
	"ABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func (r randomStringGenerator) stringWithCharset(length int, charset string) string {
	b := make([]byte, length)
	for i := range b {
		randNum := r.randomNumberGenerator.NewRandomNumber(r.clock.NowUnixNano(), len(charset))
		b[i] = charset[randNum]
	}
	return string(b)
}
