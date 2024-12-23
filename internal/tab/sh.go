package tab

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Sh(addIndent bool, cmdName string, args []string) (string, error) {
	var (
		stderr string
		err    error
	)
	if len(tab.Cursors) > 0 {
		stderr, err = tab.shCursors(addIndent, cmdName, args)
	} else {
		stderr, err = tab.shSelections(addIndent, cmdName, args)
	}

	return filterEnixLines(stderr), err
}

func (tab *Tab) shCursors(addIndent bool, cmdName string, args []string) (string, error) {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	// Execute command in shell
	shCmd := exec.Command(cmdName, args...)
	var stdout, stderr bytes.Buffer
	shCmd.Stdout = &stdout
	shCmd.Stderr = &stderr
	err := shCmd.Run()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}

	// Move cursor left for regular paste.
	stdoutStr := stdout.String()
	if !strings.HasSuffix(stdoutStr, "\n") {
		for _, c := range tab.Cursors {
			c.Left()
		}
	}

	// Paste stdout
	actions := tab.pasteCursors(stdoutStr, addIndent)
	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	return stderr.String(), nil
}

func (tab *Tab) shSelections(addIndent bool, cmdName string, args []string) (string, error) {
	return "", nil
}

func filterEnixLines(str string) string {
	b := strings.Builder{}

	lines := strings.Split(str, "\n")
	for _, line := range lines {
		line := strings.Trim(line, " \t\r")

		if !strings.HasPrefix(line, "enix:") {
			continue
		}

		b.WriteString(line[5:])
		b.WriteRune('\n')
	}

	return b.String()
}
