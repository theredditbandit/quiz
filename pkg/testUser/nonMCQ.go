package testUser

import (
	"fmt"
	"quiz/pkg/types"
	"strings"

	"github.com/charmbracelet/huh"
)

// testNonMCQ function responsible for taking a non mcq type problem and questioning the user and returning the marks and user error if any
func testNonMCQ(p types.Problem) *huh.Group {
	var ans string
	// answeredCorrectly := true // TODO : Reimplement this so marks are written to a channel
	input := huh.NewInput().
		Title(p.Question).
		Prompt(">> ").
		Value(&ans).
		Validate(func(ans string) error {
			if strings.TrimSpace(ans) == "" {
				return nil // this allows to cycle back and forth between questions
			}
			if strings.TrimSpace(ans) != strings.TrimSpace(p.Answer) {
				// answeredCorrectly = false
				return fmt.Errorf("%v is incorrect.\nthe correct answer is %v", ans, p.Answer)
			}
			return nil
		})
	return huh.NewGroup(input)
}
