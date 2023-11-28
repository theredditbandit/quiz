package testUser

import (
	"quiz/pkg/types"
)

// Questionuser takes in a problem array and total time limit prints the questions, returns marks and errors
func QuestionUser(questions []types.Problem) {
	testScore := 0
	for _, problem := range questions {
		if problem.IsMCQTypeQuestion {
			marks := testMCQ(problem)
			testScore += marks
		} else {
			marks := testNonMCQ(problem)
			testScore += marks
		}

	}

}
