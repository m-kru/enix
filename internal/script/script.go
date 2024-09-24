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

func exec(c command, t *tab.Tab) error {
	var err error

	switch c.name {
	case "down":
		err = cmd.Down(c.args, t)
	case "end":
		err = cmd.End(c.args, t)
	case "esc":
		err = cmd.Esc(c.args, t)
	case "goto":
		err = cmd.Goto(c.args, t)
	case "left":
		err = cmd.Left(c.args, t)
	case "newline":
		err = cmd.Newline(c.args, t)
	case "right":
		err = cmd.Right(c.args, t)
	case "rune":
		err = cmd.Rune(c.args, t)
	case "save":
		err = cmd.Save(c.args, t, false)
	case "space":
		err = cmd.Space(c.args, t)
	case "spawn-down":
		err = cmd.SpawnDown(c.args, t)
	case "tab":
		err = cmd.Tab(c.args, t)
	case "trim":
		err = cmd.Trim(c.args, t)
	case "up":
		err = cmd.Up(c.args, t)
	default:
		err = fmt.Errorf("invalid or unimplemented command '%s'", c.name)
	}

	return err
}
