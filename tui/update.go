package tui

import (
	tea "github.com/charmbracelet/bubbletea"
)

func updateCommitType(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "enter":
			// Send the choice on the channel and exit.
			m.commitType = CommitType[m.cursor]
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
		case "enter":
			m = GoToNextLevel(m).(Model)
			return m, nil
		}

	}

	m.scope, cmd = m.scope.Update(msg)
	return m, cmd
}

func updateDesc(msg tea.Msg, m Model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+n":
			m = GoToNextLevel(m).(Model)
			return m, nil
		}

	}

	m.desc, cmd = m.desc.Update(msg)
	return m, cmd

}
