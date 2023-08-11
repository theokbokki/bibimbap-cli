package main

import (
	"github.com/charmbracelet/lipgloss"
)

var (
	bibiBodyStyle = lipgloss.NewStyle().Foreground(orange)
	bibiNameStyle = lipgloss.NewStyle().Bold(true).Foreground(darkText)
	bibiTextStyle = lipgloss.NewStyle().Foreground(lightText)
)

// Bibi is a structure that holds the Bibi-related styles and text.
type Bibi struct {
	Text string
}

// NewBibi creates a new Bibi instance with the provided text.
func NewBibi(text string) *Bibi {
	return &Bibi{Text: text}
}

// Render renders the Bibi's text and styles.
func (b *Bibi) Render() string {
	s := bibiBodyStyle.Render("ʕ ·ᴥ·ʔ")
	s += bibiNameStyle.Render("  Bibi:")
	s += bibiBodyStyle.Render("\n│ ⍵ ⍵│  ")
	s += bibiTextStyle.Render(b.Text + "\n")
	return s
}
