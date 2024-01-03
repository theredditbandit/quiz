package colors

import "github.com/charmbracelet/lipgloss"

var GraveError = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("196"))
var GreatSuccess = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("40"))
var MuchWarning = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("184"))
var NormalText = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("86"))
