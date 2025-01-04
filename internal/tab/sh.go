package tab

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) prepareExecCmd(
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
	execCmd.Stdout = stdout
	execCmd.Stderr = stderr

	// Set environment variables
	execCmd.Env = os.Environ()
	execCmd.Env = append(execCmd.Env, fmt.Sprintf("ENIX_FILETYPE=%s", tab.FileType))
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
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var stdout, stderr bytes.Buffer
	cmd, err := tab.prepareExecCmd(&stdout, &stderr, cmdName, args)
	if err != nil {
		return "", err
	}

	err = cmd.Run()
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
