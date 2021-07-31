package util

import "math/rand"

//go:generate mockgen -source=random_number_generator.go -destination=../mocks/mock_random_number_generator.go -package=mocks

type RandomNumberGenerator interface {
	NewRandomNumber(int64, int) int
}

type randomNumberGenerator struct {}

func NewRandomNumberGenerator() RandomNumberGenerator {
	return randomNumberGenerator{}
}

func (r randomNumberGenerator) NewRandomNumber(timeInNano int64, max int) int {
	rand.Seed(timeInNano)
	return rand.Intn(max)
}
