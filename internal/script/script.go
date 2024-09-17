package script

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/m-kru/enix/internal/arg"
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

func Exec() error {
	cmds, err := parseScript()
	if err != nil {
		return err
	}

	for _, file := range arg.Files {
		tab := tab.Open(nil, nil, nil, file)
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
	case "goto":
		err = cmd.Goto(c.args, t)
	case "trim":
		err = cmd.Trim(c.args, t)
	case "up":
		err = cmd.Up(c.args, t)
	default:
		err = fmt.Errorf("invalid or unimplemented command '%s'", c.name)
	}

	return err
}
