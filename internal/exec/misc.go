package exec

import (
	"fmt"
	"os"

	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/tab"
)

func Align(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"align: expected 0 args, provided %d", len(args),
		)
	}

	err := tab.Align()
	if err != nil {
		return fmt.Errorf("align: %v", err)
	}

	return nil
}

func Change(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"change: expected 0 args, provided %d", len(args),
		)
	}

	tab.Change()

	return nil
}

func Cut(args []string, tab *tab.Tab) (string, error) {
	if len(args) > 0 {
		return "", fmt.Errorf(
			"cut: expected 0 args, provided %d", len(args),
		)
	}

	info := tab.Cut()

	return info, nil
}

func Esc(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"esc: expected 0 args, provided %d", len(args),
		)
	}

	tab.Esc()

	return nil
}

func Indent(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"indent: expected 0 args, provided %d", len(args),
		)
	}

	tab.Indent()

	return nil
}

func Trim(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"trim: expected 0 args, provided %d", len(args),
		)
	}

	tab.Trim()

	return nil
}

func Join(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"join: expected 0 args, provided %d", len(args),
		)
	}

	tab.Join()

	return nil
}

func KeyName(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) > 0 {
		return nil, fmt.Errorf(
			"key-name: expected 0 args, provided %d", len(args),
		)
	}

	path := "key-name"
	idx := 2
	for {
		if !t.Exists(path) {
			break
		}
		path = fmt.Sprintf("key-name-%d", idx)
		idx++
	}

	newT := tab.FromString(
		t.Frame,
		"insert a single key or key combo, press esc to exit\n", path,
	)
	c := newT.Cursors[0]
	c.LineNum = 2
	c.Line = c.Line.Next
	newT.State = "key-name"

	t.Append(newT)

	return newT, nil
}

func LineDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-down: expected 0 args, provided %d", len(args),
		)
	}

	tab.LineDown()

	return nil
}

func LineUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-up: expected 0 args, provided %d", len(args),
		)
	}

	tab.LineUp()

	return nil
}

func Paste(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"paste: expected 0 args, provided %d", len(args),
		)
	}

	tab.Paste()

	return nil
}

func Pwd(args []string) (string, error) {
	if len(args) > 0 {
		return "", fmt.Errorf(
			"pwd: expected 0 args, provided %d", len(args),
		)
	}

	pwd, err := os.Getwd()
	if err != nil {
		return "", fmt.Errorf("pwd: %v", err)
	}

	return pwd, nil
}

func Redo(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"redo: expected 0 args, provided %d", len(args),
		)
	}

	tab.Redo()

	return nil
}

func Search(args []string, tab *tab.Tab) error {
	if len(args) != 1 {
		return fmt.Errorf(
			"search: expected 1 arg, provided %d", len(args),
		)
	}

	return tab.Search(args[0])
}

func Sh(args []string, tab *tab.Tab) error {
	if len(args) == 0 {
		return fmt.Errorf(
			"sh: expected at least 1 arg, provided %d", len(args),
		)
	}

	indent := args[0] == "-i"
	if indent && len(args) == 1 {
		return fmt.Errorf("sh: missing command name")
	}

	var shCmd string
	var cmdArgs []string
	if indent {
		shCmd = args[1]
		cmdArgs = args[2:]
	} else {
		shCmd = args[0]
		cmdArgs = args[1:]
	}

	script, err := tab.Sh(indent, shCmd, cmdArgs)
	if err != nil {
		return fmt.Errorf("sh: %v", err)
	}

	cmds, err := cmd.ParseScript(script)
	if err != nil {
		return fmt.Errorf("sh: %v", err)
	}

	for _, cmd := range cmds {
		err := Exec(cmd, tab)
		if err != nil {
			return fmt.Errorf("sh: %v", err)
		}
	}

	return nil
}

func Undo(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"undo: expected 0 args, provided %d", len(args),
		)
	}

	tab.Undo()

	return nil
}

func Yank(args []string, tab *tab.Tab) (string, error) {
	if len(args) > 0 {
		return "", fmt.Errorf(
			"yank: expected 0 args, provided %d", len(args),
		)
	}

	info := tab.Yank()

	return info, nil
}
