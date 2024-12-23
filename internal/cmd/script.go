package cmd

import (
	"strings"
)

func ParseScript(script string) ([]Command, error) {
	var cmds []Command

	lines := strings.Split(script, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)

		// Skip empty lines and comment lines.
		if len(line) == 0 || line[0] == '#' {
			continue
		}

		c, err := Parse(line)
		if err != nil {
			return nil, err
		}

		cmds = append(cmds, c)
	}

	return cmds, nil
}
