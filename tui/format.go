package tui

import (
	"conv-cli/git_command"
	"fmt"
	"strings"
)

func Format(m Model) string {
	ctype := m.CommitType
	scope := ""

	if m.IsBreakingChange {
		ctype += "!"
	}

	if m.Scope.Value() != "" {
		scope = fmt.Sprintf("(%s)", strings.TrimSpace(m.Scope.Value()))
	}

	return fmt.Sprintf("%s%s: %s", ctype, scope, strings.TrimSpace(m.Desc.Value()))
}

func CommitBuilder(m Model) *git_commands.GitCommandBuilder {
	commit := Format(m)
	return git_commands.NewGitCmd("commit").Arg("-m", commit)
}

func CommitMsgPreview(m Model) string {
	return CommitBuilder(m).ToString()
}
