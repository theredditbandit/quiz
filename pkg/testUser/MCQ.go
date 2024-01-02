package testUser

import "quiz/pkg/types"

func testMCQ(problem types.Problem) int {

	return problem.MarksIfCorrect
}
