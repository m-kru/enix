package script

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cmd"
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

func Exec(config *cfg.Config) error {
	cmds, err := parseScript()
	if err != nil {
		return err
	}

	for _, file := range arg.Files {
		tab := tab.Open(config, nil, nil, file)
		for _, cmd := range cmds {
			err := exec(cmd, tab)
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func exec(c cmd.Command, tab *tab.Tab) error {
	var err error

	for i := 0; i < c.RepCount; i++ {
		switch c.Name {
		case "add-cursor":
			err = cmd.AddCursor(c.Args, tab)
		case "down":
			err = cmd.Down(c.Args, tab)
		case "end":
			err = cmd.End(c.Args, tab)
		case "esc":
			err = cmd.Esc(c.Args, tab)
		case "goto":
			err = cmd.Goto(c.Args, tab)
		case "left":
			err = cmd.Left(c.Args, tab)
		case "newline":
			err = cmd.Newline(c.Args, tab)
		case "right":
			err = cmd.Right(c.Args, tab)
		case "rune":
			err = cmd.Rune(c.Args, tab)
		case "save":
			err = cmd.Save(c.Args, tab, false)
		case "space":
			err = cmd.Space(c.Args, tab)
		case "spawn-down":
			err = cmd.SpawnDown(c.Args, tab)
		case "spawn-up":
			err = cmd.SpawnUp(c.Args, tab)
		case "tab":
			err = cmd.Tab(c.Args, tab)
		case "trim":
			err = cmd.Trim(c.Args, tab)
		case "up":
			err = cmd.Up(c.Args, tab)
		default:
			err = fmt.Errorf("invalid or unimplemented command '%s'", c.Name)
		}

		if err != nil {
			return err
		}
	}

	return nil
}
