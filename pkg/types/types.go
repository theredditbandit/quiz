package types

import "time"

type Problem struct {
	QuestionNumber   int
	AllowMultipleAns bool
	Question         string
	Answer           string // leave blank if MCQ
	Options          map[string]string
	MCQAnswers       []string
	IsTimed          bool
	Time             TimeConf
	MarksIfCorrect   int
	MarksIfIncorrect int
	Skippable        bool
}

// UserError represents an error made by the user in answering the question
type UserError struct {
	GivenProb Problem
	UserAns   string
	QuesNo    int
}

type ReaderFunc func() (string, error)

type TimerFunc func(t TimeConf) <-chan time.Time

type TimeConf struct {
	Time int
	Unit string
}
