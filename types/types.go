package types

type Problem struct {
	QuestionNo int
	Question   string
	Answer     string
}

// an error made by the user in answering the question
type UserError struct {
	GivenProb Problem
	UserAns   string
	QuesNo    int
}