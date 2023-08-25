package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
)

var commitType = []string{"feat", "fix", "refactor", "test", "chore"}

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
	}

	switch m.level {
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

func updateDesc(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "enter":
			m = goToNextLevel(m).(model)
			return m, nil
		}

	}

	m.desc, cmd = m.desc.Update(msg)
	return m, cmd

}

func updateScope(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "enter":
			m = goToNextLevel(m).(model)
			return m, nil
		}

	}

	m.scope, cmd = m.scope.Update(msg)
	return m, cmd
}

func updateCommitType(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q", "esc":
			return m, tea.Quit

		case "enter":
			// Send the choice on the channel and exit.
			m.commitType = commitType[m.cursor]
			m = goToNextLevel(m).(model)
			return m, nil

		case "down", "j":
			m.cursor++
			if m.cursor >= len(commitType) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(commitType) - 1
			}
		}
	}

	return m, nil
}

func (m model) View() string {
	var s string

	switch m.level {
	case CommitLevel:
		return commitTypeView(m)
	case ScopeLevel:
		return scopeView(m)
	case Desc:
		return descView(m)
	}

	return s
}

// View for choosing commit types
func commitTypeView(m model) string {
	s := strings.Builder{}
	s.WriteString("Select type of commit\n\n")

	for i := 0; i < len(commitType); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(commitType[i])
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

// View for adding scope
func scopeView(m model) string {
	return fmt.Sprintf(
		"Add scope of the commit\n\n%s\n\n%s",
		m.scope.View(),
		"(esc to quit)",
	) + "\n"
}

// View for adding desc
func descView(m model) string {
	return fmt.Sprintf(
		"Add description of the commit\n\n%s\n\n%s",
		m.desc.View(),
		"(esc to quit)",
	) + "\n"
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
