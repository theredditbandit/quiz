package static

import (
	"fmt"
	"strings"
)

type Problem struct {
	Question string
	Answer   string
}

// an error made by the user in answering the question
type UserError struct {
	GivenProb Problem
	UserAns   string
}

// a collection of UserError type , also has number of questions answered incorrectly along with error message if any
type QuizErrors struct {
	Code   int
	Msg    string
	Errors []UserError
}

func (e QuizErrors) Error() string {
	if e.Code == 0 {
		return "nil"
	}

	return fmt.Sprintf("%s %d", e.Msg, e.Code)
}

func (e QuizErrors) PrintErrors() {

	for _, val := range e.Errors {
		fmt.Println("Question ", val.GivenProb.Question, "expected answer ", strings.TrimSpace(val.GivenProb.Answer), "instead got ", val.UserAns)
	}
}
