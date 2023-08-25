package tui

import (
	"cli/git_command"
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

func CommitBuilder(m Model) *git_commands.GitCommandBuilder {
	commit := Format(m)
	return git_commands.NewGitCmd("commit").Arg("-m", commit)
}

func CommitMsgPreview(m Model) string {
	return CommitBuilder(m).ToString()
}
