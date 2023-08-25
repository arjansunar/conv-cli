package main

import (
	tea "github.com/charmbracelet/bubbletea"
)

func updateCommitType(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
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

func updateDesc(msg tea.Msg, m model) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:

		switch msg.String() {
		case "ctrl+n":
			m = goToNextLevel(m).(model)
			return m, nil
		}

	}

	m.desc, cmd = m.desc.Update(msg)
	return m, cmd

}
