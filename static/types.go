package static

import (
	"fmt"
)

type Problem struct {
	Question string
	Answer   string
}

type UserErr struct {
	GivenProb Problem
	UserAns   string
}

type UserErrors struct {
	Code int
	Msg string
	Errors []UserErr
}

func (e UserErrors) Error() string {
	if e.Code == 0 {
		return "nil"
	}

	return fmt.Sprintf("%s %d",e.Msg,e.Code)
}

func (e UserErrors) PrintErrors() {

	for _, val := range e.Errors {
		fmt.Println("Question ", val.GivenProb.Question, "expected answer ", val.GivenProb.Answer, "instead got ", val.UserAns)
	}
}
