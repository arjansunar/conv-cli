package tui

import (
	"fmt"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

type Level int

const (
	CommitLevel Level = iota
	ScopeLevel
	Desc
	Exit
)

type Model struct {
	level            Level
	cursor           int
	commitType       string
	scope            textinput.Model
	desc             textarea.Model
	isBreakingChange bool
}

func InitialModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Add Scope"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	ta := textarea.New()
	ta.Placeholder = "Describe the changes"
	ta.Focus()

	return Model{
		scope:            ti,
		level:            CommitLevel,
		desc:             ta,
		isBreakingChange: false,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {

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
			return GoToNextLevel(m), nil
		case "ctrl+p":
			return GoToPrevLevel(m), nil
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

func (m Model) currentMode() string {
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

func (m Model) View() string {
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
	s += HelpStyle.Render(fmt.Sprintf("\nMode: %s\t\nCtrl + n: Go to next • Ctrl + p: Go to prev • Ctrl + b: Toggle breaking change  • q: exit • \n", currentMode))

	return s
}
