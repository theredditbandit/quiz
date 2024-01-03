package testUser

import (
	"fmt"
	"quiz/pkg/types"
	"strings"

	"github.com/charmbracelet/huh"
)

// testNonMCQ function responsible for taking a non mcq type problem and questioning the user and returning the marks and user error if any
func testNonMCQ(p types.Problem) int {
	var ans string
	answeredCorrectly := true
	huh.NewInput().
		Title(p.Question).
		Prompt(">> ").
		Value(&ans).
		Validate(func(ans string) error {
			if strings.TrimSpace(ans) != strings.TrimSpace(p.Answer) {
				answeredCorrectly = false
				return fmt.Errorf("that's incorrect.\nthe correct answer is %v", p.Answer)
			}
			return nil
		}).
		Run()
	if answeredCorrectly {
		return p.MarksIfCorrect
	}
	return p.MarksIfIncorrect
}
