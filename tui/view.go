package tui

import (
	"fmt"
	"strings"
)

// View for choosing commit types
func commitTypeView(m Model) string {
	s := strings.Builder{}
	s.WriteString("Select type of commit\n\n")

	for i := 0; i < len(CommitType); i++ {
		if m.cursor == i {
			s.WriteString("(•) ")
		} else {
			s.WriteString("( ) ")
		}
		s.WriteString(CommitType[i])
		// add ! if breaking change
		if m.isBreakingChange {
			s.WriteString("!")
		}
		s.WriteString("\n")
	}
	s.WriteString("\n(press q to quit)\n")

	return s.String()
}

// View for adding scope
func scopeView(m Model) string {
	return fmt.Sprintf(
		"Add scope of the commit\n\n%s",
		m.scope.View(),
	) + "\n"
}

// View for adding desc
func descView(m Model) string {
	return fmt.Sprintf(
		"Add description of the commit\n\n%s",
		m.desc.View(),
	) + "\n"
}
