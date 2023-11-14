package testUser

import (
	"fmt"
	"quiz/pkg/types"
)

// QuizEvaludation: error class that holds marks scored , collection of UserErrors occoured and attempted questions
type QuizEvaluation struct {
	Attempted            int
	Unattempted          bool
	IncorrectlyAttempted []types.UserError
	UnattemptedQuestions []types.Problem
}

func (e QuizEvaluation) Error() string {
	if len(e.IncorrectlyAttempted) == 0 {
		return "nil"
	}

	return fmt.Sprintf("%d", len(e.IncorrectlyAttempted))
}

func (e QuizEvaluation) PrintErrors() {

	for _, val := range e.IncorrectlyAttempted {
		fmt.Print("Question ", val.QuesNo, ") ", val.GivenProb.Question, " expected answer ", val.GivenProb.Answer, " instead got ", val.UserAns, "\n")
	}
	fmt.Println()
}

// PrintUnattempted questions after time runs out.
func (e QuizEvaluation) PrintUnattempted() {
	fmt.Println("You missed the following questions.")
	for _, val := range e.UnattemptedQuestions {
		e.Attempted++
		fmt.Printf("%d) %s \n", e.Attempted, val.Question)
	}
}
