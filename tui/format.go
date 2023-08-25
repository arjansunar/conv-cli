package tui

import (
	"fmt"
)

func Format(m Model) string {
	ctype := m.CommitType
	scope := ""

	if m.IsBreakingChange {
		ctype += "!"
	}

	if m.Scope.Value() != "" {
		scope = fmt.Sprintf("(%s)", m.Scope.Value())
	}

	return fmt.Sprintf("%s%s: %s", ctype, scope, m.Desc.Value())
}
