package testUser

import (
	"fmt"
	"github.com/charmbracelet/huh"
	"quiz/pkg/types"
	"quiz/pkg/utils"
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
	err := huh.NewForm(formQ...).Run()
	if err != nil {
		utils.ExitWithMessage(fmt.Sprintf("TODO Check the error for 'user aborted' message and then ask whether to retry/review/print result etc. . .\n'%T'\n", err), 1)
	}
}
