package tui

import tea "github.com/charmbracelet/bubbletea"

func GoToNextLevel(m Model) tea.Model {
	m.level++
	return m
}

func GoToPrevLevel(m Model) tea.Model {
	m.level--
	return m
}
