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

type command struct {
	name string   // Command name
	args []string // Command arguments
}

func parseScript() ([]command, error) {
	var cmds []command

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

		name, args, _ := strings.Cut(line, " ")
		cmds = append(
			cmds,
			command{
				strings.TrimSpace(name),
				strings.Fields(args),
			},
		)
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

func exec(c command, tab *tab.Tab) error {
	var err error

	switch c.name {
	case "add-cursor":
		err = cmd.AddCursor(c.args, tab)
	case "down":
		err = cmd.Down(c.args, tab)
	case "end":
		err = cmd.End(c.args, tab)
	case "esc":
		err = cmd.Esc(c.args, tab)
	case "goto":
		err = cmd.Goto(c.args, tab)
	case "left":
		err = cmd.Left(c.args, tab)
	case "newline":
		err = cmd.Newline(c.args, tab)
	case "right":
		err = cmd.Right(c.args, tab)
	case "rune":
		err = cmd.Rune(c.args, tab)
	case "save":
		err = cmd.Save(c.args, tab, false)
	case "space":
		err = cmd.Space(c.args, tab)
	case "spawn-down":
		err = cmd.SpawnDown(c.args, tab)
	case "spawn-up":
		err = cmd.SpawnUp(c.args, tab)
	case "tab":
		err = cmd.Tab(c.args, tab)
	case "trim":
		err = cmd.Trim(c.args, tab)
	case "up":
		err = cmd.Up(c.args, tab)
	default:
		err = fmt.Errorf("invalid or unimplemented command '%s'", c.name)
	}

	return err
}
