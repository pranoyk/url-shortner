package util

import "time"

//go:generate mockgen -source=clock.go -destination=../mocks/mock_clock.go -package=mocks

type Clock interface {
	Now() time.Time
	NowUnixNano() int64
}

type clock struct{}

func (c clock) NowUnixNano() int64 {
	return c.Now().UnixNano()
}

func (c clock) Now() time.Time {
	return time.Now()
}

func NewClock() Clock {
	return clock{}
}
