package types

import "time"

type Problem struct {
	QuestionNumber    int    // Question number
	Question          string // Question text
	IsMCQTypeQuestion bool   // True if MCQ type question
	Answer            string // not needed for MCQ questions
	AllowMultipleAns  bool   // True if multiple answers allowed for MCQ type question
	Options           map[string]string
	MCQAnswers        []string
	IsTimed           bool
	Time              TimeConf
	MarksIfCorrect    int
	MarksIfIncorrect  int
	Skippable         bool
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
