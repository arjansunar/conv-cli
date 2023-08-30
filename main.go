package main

import (
	"fmt"
	"github.com/arjansunar/conv-cli/tui"
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
	if m, ok := m.(tui.Model); ok && m.CommitMsg != nil {
		if m.Err != "" || !m.ToCommit {
			fmt.Println("Exiting...\n", m.Err)
			os.Exit(1)
		}
		fmt.Printf("\nCommiting your message....")
		if _, err := m.CommitMsg.Run(); err != nil {
			fmt.Println("Oh no:", err)
			os.Exit(1)
		}
		fmt.Printf("\nCommited your message!")
	}
}
