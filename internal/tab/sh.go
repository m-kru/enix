package tab

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Sh(addIndent bool, cmdName string, args []string) error {
	var err error
	if len(tab.Cursors) > 0 {
		err = tab.shCursors(addIndent, cmdName, args)
	} else {
		err = tab.shSelections(addIndent, cmdName, args)
	}

	return err
}

func (tab *Tab) shCursors(addIndent bool, cmdName string, args []string) error {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	// Execute command in shell
	cmd := exec.Command(cmdName, args...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return fmt.Errorf("%v: %s", err, stderr.String())
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

	// Execute enix commands from stderr

	return nil
}

func (tab *Tab) shSelections(addIndent bool, cmdName string, args []string) error {
	return nil
}
