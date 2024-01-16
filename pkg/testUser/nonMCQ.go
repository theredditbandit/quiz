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
	// TODO : Implement the scoring system with go channels
	input := huh.NewInput().
		Title(p.Question).
		Prompt(">> ").
		Value(&ans).
		Validate(func(ans string) error {
			if strings.TrimSpace(ans) == "" && p.Skippable {
				return nil // this allows to cycle back and forth between questions
			}
			if strings.TrimSpace(ans) != strings.TrimSpace(p.Answer) {
				if !p.Skippable && strings.TrimSpace(ans) == "" {
					return fmt.Errorf("cannot skip this question")
				}
				return fmt.Errorf("%v is incorrect.\nthe correct answer is %v", ans, p.Answer)
			}
			return nil
		})
	return huh.NewGroup(input)
}
