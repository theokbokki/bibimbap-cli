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
	p := tea.NewProgram(initialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	frameworks []Framework
	cursor     int
	bibi       *Bibi
	help       []Help
}

func initialModel() model {
	return model{
		frameworks: []Framework{
			*NewFramework("Sveltekit", "Web development streamlined"),
			*NewFramework("Nuxt", "The intuitive web framework"),
			*NewFramework("Next.js", "The React framework for the web"),
		},
		bibi: NewBibi("Welcome to Bibimbap!"),
		help: []Help{
			*NewHelp("Quit", []string{"ctrl-c"}),
		},
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

		case "enter":
			m.bibi.Text = "Ok great! I'll set things up for you, wait a sec..."
			m.Setup(m.frameworks[m.cursor])
		}
	}

	if m.bibi.Text == "Welcome to Bibimbap!" {
		time.Sleep(1250 * time.Millisecond)
		m.bibi.Text = "Choose a framework in the list below"
	}

	return m, nil
}

func (m model) View() string {
	f := "\n"
	h := ""

	if m.bibi.Text == "Choose a framework in the list below" {
		for i, framework := range m.frameworks {
			if m.cursor == i {
				framework.Selected = true
			}

			f += fmt.Sprintf("%s\n", framework.Render())
		}

		m.help = append(
			m.help,
			*NewHelp("Up", []string{"↑", "k"}),
			*NewHelp("Down", []string{"↓", "j"}),
			*NewHelp("Select", []string{"enter"}),
		)

	}

	for i, help := range m.help {
		h += fmt.Sprintf("%s", help.Render())

		if i < len(m.help)-1 {
			h += "\n"
		}
	}

	return lipgloss.NewStyle().Padding(1, 4).Render(lipgloss.JoinVertical(0, m.bibi.Render(), f, h))
}
