package testUser

import (
	"fmt"
	"quiz/pkg/types"

	"github.com/charmbracelet/huh"
)

func testMCQ(problem types.Problem) int {
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().Title(problem.Question).
				Options(
					huh.NewOption("A", "val"),
				),
		),
	)
	err := form.Run()
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	return problem.MarksIfCorrect
}
