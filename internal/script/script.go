package script

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cmd"
	"github.com/m-kru/enix/internal/exec"
	"github.com/m-kru/enix/internal/tab"
)

func parseScript() ([]cmd.Command, error) {
	var cmds []cmd.Command

	script, err := os.Open(arg.Script)
	if err != nil {
		return nil, err
	}
	defer script.Close()

	scanner := bufio.NewScanner(script)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		// Skip empty lines and comment lines.
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		c, err := cmd.Parse(line)
		if err != nil {
			return nil, err
		}

		cmds = append(cmds, c)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return cmds, nil
}

func Exec() error {
	cmds, err := parseScript()
	if err != nil {
		return err
	}

	for _, file := range arg.Files {
		tab, err := tab.Open(nil, nil, file)
		if err != nil {
			return err
		}
		for _, cmd := range cmds {
			err := execCmd(cmd, tab)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func execCmd(c cmd.Command, tab *tab.Tab) error {
	var err error

	for range c.RepCount {
		switch c.Name {
		case "add-cursor":
			err = exec.AddCursor(c.Args, tab)
		case "align":
			err = exec.Align(c.Args, tab)
		case "backspace":
			err = exec.Backspace(c.Args, tab)
		case "change":
			err = exec.Change(c.Args, tab)
		case "del":
			err = exec.Del(c.Args, tab)
		case "down":
			err = exec.Down(c.Args, tab)
		case "esc":
			err = exec.Esc(c.Args, tab)
		case "go":
			err = exec.Go(c.Args, tab)
		case "join":
			err = exec.Join(c.Args, tab)
		case "left":
			err = exec.Left(c.Args, tab)
		case "line-down":
			err = exec.LineDown(c.Args, tab)
		case "line-end":
			err = exec.LineEnd(c.Args, tab)
		case "line-up":
			err = exec.LineUp(c.Args, tab)
		case "mark":
			_, err = exec.Mark(c.Args, tab)
		case "newline":
			err = exec.Newline(c.Args, tab)
		case "prev-word-start":
			err = exec.PrevWordStart(c.Args, tab)
		case "right":
			err = exec.Right(c.Args, tab)
		case "rune":
			err = exec.Rune(c.Args, tab)
		case "save":
			_, err = exec.Save(c.Args, tab, false)
		case "sel-line":
			err = exec.SelLine(c.Args, tab)
		case "sel-right":
			err = exec.SelRight(c.Args, tab)
		case "sel-word-end":
			err = exec.SelWordEnd(c.Args, tab)
		case "space":
			err = exec.Space(c.Args, tab)
		case "spawn-down":
			err = exec.SpawnDown(c.Args, tab)
		case "spawn-up":
			err = exec.SpawnUp(c.Args, tab)
		case "tab":
			err = exec.Tab(c.Args, tab)
		case "trim":
			err = exec.Trim(c.Args, tab)
		case "undo":
			err = exec.Undo(c.Args, tab)
		case "up":
			err = exec.Up(c.Args, tab)
		case "word-end":
			err = exec.WordEnd(c.Args, tab)
		case "word-start":
			err = exec.WordStart(c.Args, tab)
		default:
			err = fmt.Errorf("invalid or unimplemented command '%s'", c.Name)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
