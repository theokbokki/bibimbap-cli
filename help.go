package main

import "github.com/charmbracelet/lipgloss"

var helpStyle = lipgloss.NewStyle().Foreground(lightText)

type Help struct {
	Keys   []string
	Action string
}

func NewHelp(action string, keys []string) *Help {
	return &Help{
		Keys:   keys,
		Action: action,
	}
}

func (h *Help) Render() string {
	s := h.Action + ": "

	for i, key := range h.Keys {
		if i < len(h.Keys)-1 {
			s += "\"" + key + "\"/"
		} else {
			s += "\"" + key + "\""
		}
	}

	return helpStyle.Render(s)
}
