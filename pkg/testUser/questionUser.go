package testUser

import (
	"quiz/pkg/types"

	"github.com/charmbracelet/huh"
)

// Questionuser takes in a problem array and total time limit prints the questions, returns marks and errors
func QuestionUser(questions []types.Problem) {
	testScore := 0
	var formQ []*huh.Group
	for _, problem := range questions {
		if problem.IsMCQTypeQuestion {
			marks := testMCQ(problem)
			testScore += marks
		} else {
			q := testNonMCQ(problem)
			formQ = append(formQ, q)
		}
	}
	huh.NewForm(formQ...).Run()
}
