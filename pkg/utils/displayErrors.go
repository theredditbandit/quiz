package utils

import "github.com/charmbracelet/huh"

// PrintErrorsAndWarnings checks if the user wants to print erros and warnings found while validating
func PrintErrorsAndWarnings() bool {
	var confirm bool
	err := huh.NewConfirm().
		Title("Would you like to review validation errors").
		Affirmative("Yes").
		Negative("No").
		Value(&confirm).
		Run()
	if err != nil {
		confirm = false
	}
	return confirm
}
