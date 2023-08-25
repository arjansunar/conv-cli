package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func updateCommitType(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+n", "enter":
			// Send the choice on the channel and exit.
			m.CommitType = CommitType[m.cursor]
			m = GoToNextLevel(m).(Model)
			return m, nil

		case "down", "j":
			m.cursor++
			if m.cursor >= len(CommitType) {
				m.cursor = 0
			}

		case "up", "k":
			m.cursor--
			if m.cursor < 0 {
				m.cursor = len(CommitType) - 1
			}
		}
	}

	return m, nil
}

func updateScope(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+n", "enter":
			if m.Scope.Value() != "" {
				m = GoToNextLevel(m).(Model)
			} else {
				m.Err = "Scope cannot be empty"
			}
			return m, nil
		}

	}

	m.Scope, cmd = m.Scope.Update(msg)
	return m, cmd
}

func updateDesc(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+n", "enter":
			if m.Scope.Value() != "" {
				m = GoToNextLevel(m).(Model)
			} else {
				m.Err = "Scope cannot be empty"
			}

			return m, nil
		}

	}

	m.Desc, cmd = m.Desc.Update(msg)
	return m, cmd

}
