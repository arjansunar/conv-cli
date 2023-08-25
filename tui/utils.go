package tui

import "github.com/charmbracelet/lipgloss"

var CommitType = []string{"feat", "fix", "refactor", "test", "chore"}

var (
	HelpStyle  = lipgloss.NewStyle().Foreground(lipgloss.Color("241"))
	ErrorStyle = lipgloss.NewStyle().Foreground(lipgloss.Color("9"))
)
