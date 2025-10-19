package tab

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) prepareExecCmd(
	stdin io.Reader,
	stdout io.Writer,
	stderr io.Writer,
	cmdName string,
	args []string,
) (*exec.Cmd, error) {
	// Try to get a shell
	shell := os.Getenv("SHELL")
	if shell == "" {
		return nil, fmt.Errorf("can't get environment variable $SHELL")
	}

	// Prepare exec arguments
	execArgs := []string{"-c"}
	b := strings.Builder{}
	b.WriteString(cmdName)
	for _, a := range args {
		b.WriteRune(' ')
		b.WriteString(a)
	}
	execArgs = append(execArgs, b.String())

	execCmd := exec.Command(shell, execArgs...)
	execCmd.Stdin = stdin
	execCmd.Stdout = stdout
	execCmd.Stderr = stderr

	// Set environment variables
	execCmd.Env = os.Environ()
	execCmd.Env = append(execCmd.Env, fmt.Sprintf("ENIX_FILETYPE=%s", tab.Filetype))
	path := tab.Path
	if !filepath.IsAbs(path) {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}
		path = filepath.Join(wd, path)
	}
	execCmd.Env = append(execCmd.Env, fmt.Sprintf("ENIX_FILEPATH=%s", path))

	return execCmd, nil
}

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
	var stdout, stderr bytes.Buffer
	cmd, err := tab.prepareExecCmd(nil, &stdout, &stderr, cmdName, args)
	if err != nil {
		return "", err
	}

	err = cmd.Run()
	if err != nil {
		return "", fmt.Errorf("%v: %s", err, stderr.String())
	}

	text := stdout.String()

	prevCurs := cursor.Clone(tab.Cursors)
	cursors := tab.Cursors
	if strings.HasSuffix(text, "\n") {
		cursors = cursor.LineUnique(cursors, true)
	}

	// Paste stdout
	actions := make(action.Actions, 0, 2*len(cursors))
	newSels := make([]*sel.Selection, 0, len(cursors))

	for curIdx, cur := range cursors {
		startCur, endCur, acts := cur.Paste(text, false)

		for _, c := range cursors[curIdx+1:] {
			c.Inform(acts)
		}

		for _, m := range tab.Marks {
			m.Inform(acts)
		}

		for _, s := range newSels {
			s.Inform(acts, true)
		}

		actions = append(actions, acts)
		newSels = append(newSels, sel.FromTo(startCur, endCur))
	}

	tab.undoPush(actions.Reverse(), prevCurs, nil)

	tab.Cursors = nil
	tab.Selections = newSels

	return stderr.String(), nil
}

func (tab *Tab) shSelections(addIndent bool, cmdName string, args []string) (string, error) {
	prevSels := sel.Clone(tab.Selections)

	actions := make(action.Actions, 0, 2*len(tab.Selections))
	newSels := make([]*sel.Selection, 0, len(tab.Selections))

	var stdErrStr string

	for i, s := range tab.Selections {
		str := s.ToString()

		// Execute command
		var stdout, stderr bytes.Buffer
		cmd, err := tab.prepareExecCmd(strings.NewReader(str), &stdout, &stderr, cmdName, args)
		if err != nil {
			return "", err
		}
		err = cmd.Run()
		if err != nil {
			return "", fmt.Errorf("%v: %s", err, stderr.String())
		}

		stdErrStr += stderr.String()

		// Delete selection text
		acts := s.Delete()
		if len(acts) > 0 {
			actions = append(actions, acts)
			tab.handleAction(acts)
		}

		// Inform selections and marks about delete actions
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(acts, true)
		}
		for _, s2 := range newSels {
			s2.Inform(acts, true)
		}
		for _, m := range tab.Marks {
			m.Inform(acts)
		}

		// Create cursor from the first selection rune
		cur := cursor.New(s.Line, s.LineNum, s.StartRuneIdx)

		// Paste stdout text
		startCur, endCur, acts := cur.PasteBefore(stdout.String(), false)
		if len(acts) > 0 {
			actions = append(actions, acts)
			tab.handleAction(acts)
		}

		// Inform selections and marks about paste actions
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(acts, true)
		}
		for _, s2 := range newSels {
			s2.Inform(acts, true)
		}
		for _, m := range tab.Marks {
			m.Inform(acts)
		}

		// Create new selection
		newSels = append(newSels, sel.FromTo(startCur, endCur))
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), nil, prevSels)
		tab.Selections = newSels
	}

	return stdErrStr, nil
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
