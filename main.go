package main

import (
	"fmt"
	"log"
	"time"

	"github.com/charmbracelet/bubbles/textinput"
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
	step       string
	textinput  textinput.Model
	frameworks []Framework
	cursor     int
	bibi       *Bibi
	help       []Help
	err        string
}

func initialModel() model {

	return model{
		step:      "initial",
		textinput: NewInput(),
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
	var cmd tea.Cmd

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
			if m.step == "input" {
				if m.textinput.Value() == "" {
					m.err = "Your project cannot have an empty name!"
				} else {
					m.err = ""
					m.textinput.Blur()
					m.bibi.Text = "Choose a framework in the list below"
					m.step = "framework"
				}
			} else if m.step == "framework" {
				m.bibi.Text = "Ok great! I'll set things up for you, wait a sec..."
				m.step = "setup"
				m.Setup(m.frameworks[m.cursor])
			}
		}
	}

	if m.step == "initial" {
		time.Sleep(1250 * time.Millisecond)
		m.bibi.Text = "What's the name of your project?"
		m.step = "input"
	}

	m.textinput, cmd = m.textinput.Update(msg)

	return m, cmd
}

func (m model) View() string {
	s := "\n"
	h := ""

	if m.step == "input" {
		s += m.textinput.View()
		h += "\n"

		m.help = append(
			m.help,
			*NewHelp("Validate", []string{"enter"}),
		)
	}

	if m.step == "framework" {
		for i, framework := range m.frameworks {
			if m.cursor == i {
				framework.Selected = true
			}

			s += fmt.Sprintf("%s\n", framework.Render())
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

	return lipgloss.NewStyle().
		Padding(1, 4).
		Render(lipgloss.JoinVertical(
			0,
			m.bibi.Render(),
			s,
			func() string {
				if m.err != "" {
					return "\n" + lipgloss.NewStyle().Foreground(orange).Render(m.err)
				}
				return ""
			}(),
			h))
}
