package types

import "time"

type Problem struct {
	Question string
	Answer   string
}

// an error made by the user in answering the question
type UserError struct {
	GivenProb Problem
	UserAns   string
	QuesNo    int
}

type ReaderFunc func() (string, error)

type TimerFunc func(d time.Duration) <-chan time.Time
