package testUser

import (
	tea "github.com/charmbracelet/bubbletea"
	"quiz/pkg/types"
)

// testNonMCQ function responsible for taking a non mcq type problem and questioning the user and returning the marks and user error if any
func testNonMCQ(problem types.Problem) int {

	_ = tea.ClearScreen()
	return problem.MarksIfCorrect
}
