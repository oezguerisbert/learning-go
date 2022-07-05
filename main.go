package main

import (
	"fmt"
	"os"
	time "time"

	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	displayedTime time.Time
}

type TickMsg time.Time

func tickEvery() tea.Cmd {
	return tea.Every(time.Second, func(t time.Time) tea.Msg {
		return TickMsg(t)
	})
}

func initialModel() model {
	return model{
		displayedTime: time.Now(),

	}
}

func (m model) Init() tea.Cmd {
	return tickEvery()
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		}
	case TickMsg:
		m.displayedTime = time.Now()
		return m, tickEvery()
	}
	
	return m, nil
}

func (m model) View() string {
	timeFormat := "3:04:05 PM"
	s := m.displayedTime.Format(timeFormat)
	s += "\nPress q to quit.\n"

	return s
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}
}
