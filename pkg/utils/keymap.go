package utils

import (
	"github.com/charmbracelet/bubbles/key"
	"github.com/charmbracelet/huh"
)

func NonMCQMap() *huh.KeyMap {
	return &huh.KeyMap{
		Quit: key.NewBinding(key.WithKeys("ctrl+q"), key.WithHelp("ctrl+q", "quit")),
		Input: huh.InputKeyMap{
			Next: key.NewBinding(key.WithKeys("enter"), key.WithHelp("enter", "next")),
		},
	}
}
