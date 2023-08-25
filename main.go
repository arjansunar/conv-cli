package main

import (
	"cli/tui"
	"fmt"
	tea "github.com/charmbracelet/bubbletea"
	"os"
)

func main() {
	p := tea.NewProgram(tui.InitialModel())

	// Run returns the model as a tea.Model.
	m, err := p.Run()
	if err != nil {
		fmt.Println("Oh no:", err)
		os.Exit(1)
	}

	// Assert the final tea.Model to our local model and print the choice.
	if _, ok := m.(tui.Model); ok {
		fmt.Printf("\n---\nYou chose %s!\n", "RESULT")
	}
}
