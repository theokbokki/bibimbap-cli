package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	selectedStyle    = lipgloss.NewStyle().Foreground(orange)
	titleStyle       = lipgloss.NewStyle().Foreground(darkText)
	descriptionStyle = lipgloss.NewStyle().Foreground(lightText)
	cusrorStyle      = lipgloss.NewStyle().Bold(true)
)

type Framework struct {
	Title       string
	Description string
	Selected    bool
}

func NewFramework(title string, description string) *Framework {
	return &Framework{
		Title:       title,
		Description: description,
	}
}

func (f *Framework) Render() string {
	s := ""
	if f.Selected {
		s += selectedStyle.Render("│ " + f.Title + "\n│ " + f.Description + "\n")
	} else {
		s += titleStyle.Render("│ " + f.Title)
		s += titleStyle.Render("\n│ ")
		s += descriptionStyle.Render(f.Description + "\n")
	}
	return s
}
