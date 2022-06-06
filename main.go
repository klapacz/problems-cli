package main

import (
	"fmt"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	problems []string
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {

		case "ctrl+c", "q":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := ""

	for _, problem := range m.problems {
		s += fmt.Sprintf("%s\n", problem)
	}

	return s
}

func main() {
	p := tea.NewProgram(model{
		problems: []string{"Workplace Issues", "Emptiness", "Friendship Issues"},
	})

	if err := p.Start(); err != nil {
		panic(err)
	}
}
