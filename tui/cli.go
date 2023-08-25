package tui

import (
	git_commands "cli/git_command"
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
	Commiting
	Exit
)

type Model struct {
	level            Level
	cursor           int
	CommitType       string
	Scope            textinput.Model
	Desc             textarea.Model
	IsBreakingChange bool
	Err              string
	CommitMsg        *git_commands.GitCommandBuilder
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
		Scope:            ti,
		level:            CommitLevel,
		Desc:             ta,
		IsBreakingChange: false,
	}
}

func (m Model) currentMode() string {
	switch m.level {
	case CommitLevel:
		return "TYPE"
	case ScopeLevel:
		return "SCOPE"
	case Desc:
		return "DESC"
	case Commiting:
		return "COMMITING"
	case Exit:
		return "EXIT"
	}

	return "LOADING..."
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
			m.IsBreakingChange = !m.IsBreakingChange
		case "ctrl+p":
			return GoToPrevLevel(m), nil
		}

	}

	switch m.level {
	case CommitLevel:
		m.Err = ""
		return updateCommitType(msg, m)
	case ScopeLevel:
		return updateScope(msg, m)
	case Desc:
		return updateDesc(msg, m)
	case Commiting:
		m.CommitMsg = CommitBuilder(m)
		return m, nil
	}

	return m, tea.Quit
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
	case Commiting:
		s += commitView(m)
	}

	currentMode := m.currentMode()
	s += HelpStyle.Render(fmt.Sprintf("\nMode: %s\t\nCtrl + n: Go to next • Ctrl + p: Go to prev • Ctrl + b: Toggle breaking change  • q: exit • \n", currentMode))

	if m.Err != "" {
		s += ErrorStyle.Render(m.Err)
	}

	s += ErrorStyle.Render(CommitMsgPreview(m))

	return s
}
