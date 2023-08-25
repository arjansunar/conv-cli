package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var commitType = []string{"feat", "fix", "refactor", "test", "chore"}

var (
	helpStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
)

type Level int

const (
	CommitLevel Level = iota
	ScopeLevel
	Desc
	Exit
)

type model struct {
	level            Level
	cursor           int
	commitType       string
	scope            textinput.Model
	desc             textarea.Model
	isBreakingChange bool
}

func initialModel() model {
	ti := textinput.New()
	ti.Placeholder = "Add Scope"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	ta := textarea.New()
	ta.Placeholder = "Describe the changes"
	ta.Focus()

	return model{
		scope:            ti,
		level:            CommitLevel,
		desc:             ta,
		isBreakingChange: false,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

	// Make sure these keys always quit
	if msg, ok := msg.(tea.KeyMsg); ok {
		k := msg.String()
		if k == "q" || k == "esc" || k == "ctrl+c" {
			return m, tea.Quit
		}

		switch k {

		case "ctrl+b":
			m.isBreakingChange = !m.isBreakingChange
		case "ctrl+n":
			return goToNextLevel(m), nil
		case "ctrl+p":
			return goToPrevLevel(m), nil
		}

	}

	switch m.level {
	case -1:
		return m, tea.Quit
	case CommitLevel:
		return updateCommitType(msg, m)

	case ScopeLevel:
		return updateScope(msg, m)

	case Desc:
		return updateDesc(msg, m)
	case Exit:
		return m, tea.Quit
	}

	return m, nil
}

func goToNextLevel(m model) tea.Model {
	m.level++
	return m
}

func goToPrevLevel(m model) tea.Model {
	m.level--
	return m
}

func (m model) currentMode() string {
	switch m.level {
	case CommitLevel:
		return "commit"
	case ScopeLevel:
		return "scope"
	case Desc:
		return "desc"
	case Exit:
		return "exit"
	}

	return "spinner"
}

func (m model) View() string {
	var s string

	switch m.level {
	case CommitLevel:
		s += commitTypeView(m)
	case ScopeLevel:
		s += scopeView(m)
	case Desc:
		s += descView(m)
	case Exit:
		s += "Exiting..."
	}

	currentMode := m.currentMode()
	s += helpStyle.Render(fmt.Sprintf("\nMode: %s\t\nCtrl + n: Go to next • Ctrl + p: Go to prev • Ctrl + b: Toggle breaking change  • q: exit • \n", currentMode))

	return s
}

func main() {
	p := tea.NewProgram(initialModel())

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if m, ok := m.(model); ok && m.commitType != "" {
		fmt.Printf("\n---\nYou chose %s!\n", m.commitType)
	}
}
