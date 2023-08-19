package main

import (
	"github.com/charmbracelet/bubbles/textinput"
	"github.com/charmbracelet/lipgloss"
)

var (
	cursorStyle = lipgloss.NewStyle().
			Foreground(orange)

	placeholderStyle = lipgloss.NewStyle().
				Foreground(lightText)

	textStyle = lipgloss.NewStyle().Foreground(darkText)
)

func NewInput() textinput.Model {
	ti := textinput.New()
	ti.Placeholder = "world-domination"
	ti.Focus()
	ti.CharLimit = 0
	ti.Cursor.Style = cursorStyle
	ti.PlaceholderStyle = placeholderStyle
	ti.TextStyle = textStyle

	return ti
}
