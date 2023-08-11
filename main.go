package main

import (
	"fmt"
	"log"
	"time"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// The function that launches the program
func main() {
	p := tea.NewProgram(initialModel())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

// A type for our error messages
type (
	errMsg error
)

type model struct {
	frameworks []Framework
	cursor     int
	bibi       *Bibi
}

func initialModel() model {
	return model{
		frameworks: []Framework{
			*NewFramework("Sveltekit", "Web development streamlined"),
			*NewFramework("Nuxt", "The intuitive web framework"),
			*NewFramework("Next.js", "The React framework for the web"),
		},
		bibi: NewBibi("Welcome to Bibimbap!"),
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	// The message is a keypress
	case tea.KeyMsg:
		// What key was pressed
		switch msg.String() {
		// Quitting the program
		case "ctrl+c":
			return m, tea.Quit
		// Going up
		case "up", "k":
			if m.cursor > 0 {
				m.cursor--
			}
		// Going down
		case "down", "j":
			if m.cursor < len(m.frameworks)-1 {
				m.cursor++
			}
		}
	}

	if m.bibi.Text == "Welcome to Bibimbap!" {
		time.Sleep(1250 * time.Millisecond)
		m.bibi.Text = "Choose a framework in the list below"
	}

	return m, nil
}

func (m model) View() string {
	s := "\n\n"
	for i, framework := range m.frameworks {
		if m.cursor == i {
			framework.Selected = true
		}

		s += fmt.Sprintf("%s\n", framework.Render())
	}

	return lipgloss.NewStyle().Padding(1, 4).Render(m.bibi.Render() + s)
}
