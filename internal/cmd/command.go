package cmd

import (
	"strconv"
	"strings"
	"unicode"
)

type Command struct {
	RepCount int
	Name     string
	Args     []string
}

// Parse parses command line string and returns a command.
func Parse(line string) (Command, error) {
	cmd := Command{RepCount: 1, Name: "", Args: nil}

	// The command might be a short version of go command.
	r0 := []rune(line)[0]
	if unicode.IsDigit(r0) || r0 == '-' {
		cmd, ok := parseShortGoto(line)
		if ok {
			return cmd, nil
		}
	}

	fields := strings.Fields(line)

	if len(fields) == 1 {
		cmd.Name = fields[0]
		return cmd, nil
	}

	if unicode.IsDigit([]rune(fields[0])[0]) {
		repCount, err := strconv.Atoi(fields[0])
		if err != nil {
			return cmd, err
		}
		cmd.RepCount = repCount
		cmd.Name = fields[1]
		if len(fields) > 2 {
			cmd.Args = fields[2:]
		}
	} else {
		cmd.Name = fields[0]
		if len(fields) > 1 {
			cmd.Args = fields[1:]
		}
	}

	return cmd, nil
}

// Valid versions of short go command are, for example:
//   - 1
//   - 1:2
//   - 1 2
//   - -1
//   - -1:2
//   - -1 2
func parseShortGoto(line string) (Command, bool) {
	cmd := Command{RepCount: 1, Name: "go", Args: nil}

	for _, r := range line {
		if !unicode.IsDigit(r) && r != ':' && r != ' ' && r != '-' {
			return cmd, false
		}
	}

	cmd.Args = strings.Fields(line)

	return cmd, true
}
