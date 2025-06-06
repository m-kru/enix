package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strconv"
	"strings"
	"unicode"
)

func Down(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"down: provided %d args, expected 0", len(args),
		)
	}

	tab.Down()

	return nil
}

func Left(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"left: provided %d args, expected 0", len(args),
		)
	}

	tab.Left()

	return nil
}

func Right(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"right: provided %d args, expected 0", len(args),
		)
	}

	tab.Right()

	return nil
}

func Up(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"up: provided %d args, expected 0", len(args),
		)
	}

	tab.Up()

	return nil
}

func Go(args []string, tab *tab.Tab) error {
	var err error

	if len(args) == 0 {
		return fmt.Errorf("go: missing at least line number or mark name")
	} else if len(args) > 2 {
		return fmt.Errorf("go: expected at most 2 args, provided %d", len(args))
	}

	arg0 := args[0]
	arg1 := ""
	if len(args) > 1 {
		arg1 = args[1]
	}

	if len(args) == 1 && !unicode.IsDigit([]rune(arg0)[0]) && arg0[0] != '-' {
		return tab.GoMark(arg0)
	}

	lineStr := arg0
	colStr := ""
	line := 0
	col := 0

	if arg1 == "" {
		if strings.Contains(arg0, ":") {
			lineStr, colStr, _ = strings.Cut(arg0, ":")
		}
	} else {
		lineStr = arg0
		colStr = arg1
	}

	// Prase line
	line, err = strconv.Atoi(lineStr)
	if err != nil {
		return fmt.Errorf("go: %v", err)
	}

	// Parse column
	if colStr != "" {
		col, err = strconv.Atoi(colStr)
		if err != nil {
			return fmt.Errorf("go: %v", err)
		}
	}

	tab.Go(line, col)

	return nil
}

func PrevWordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"prev-word-start: expected 0 args, provided %d", len(args),
		)
	}

	tab.PrevWordStart()

	return nil
}

func WordEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"word-end: expected 0 args, provided %d", len(args),
		)
	}

	tab.WordEnd()

	return nil
}

func WordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"word-start: expected 0 args, provided %d", len(args),
		)
	}

	tab.WordStart()

	return nil
}

func SpawnDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"spawn-down: expected 0 args, provided %d", len(args),
		)
	}

	tab.SpawnDown()

	return nil
}

func SpawnUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"spawn-up: expected 0 args, provided %d", len(args),
		)
	}

	tab.SpawnUp()

	return nil
}

func LineStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-start: expected 0 args, provided %d", len(args),
		)
	}

	tab.LineStart()

	return nil
}

func LineEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-end: expected 0 args, provided %d", len(args),
		)
	}

	tab.LineEnd()

	return nil
}

func AddCursor(args []string, tab *tab.Tab) error {
	if len(args) < 1 || 2 < len(args) {
		return fmt.Errorf(
			"add-cursor: provided %d args, expected 1 or 2", len(args),
		)
	}

	line, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("add-cursor: %v", err)
	}
	if line < 1 {
		return fmt.Errorf(
			"add-cursor: line number must be positive, current value %d", line,
		)
	}

	col := 1
	if len(args) == 2 {
		col, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("add-cursor: %v", err)
		}
	}
	if col < 1 {
		return fmt.Errorf(
			"add-cursor: column number must be positive, current value %d", col,
		)
	}

	tab.AddCursor(line, col)

	return nil
}

func DumpCursor(args []string, tab *tab.Tab) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf(
			"dump-cursor: provided %d args, expected 1", len(args),
		)
	}

	n, err := strconv.Atoi(args[0])
	if err != nil {
		return "", fmt.Errorf("dump-cursor: %v", err)
	}

	if n < 1 {
		return "", fmt.Errorf(
			"dump-cursor: cursor index must be positive, current value %d", n,
		)
	}

	if n > len(tab.Cursors) {
		return "", fmt.Errorf(
			"dump-cursor: can't get %d cursor, there are %d cursors",
			n, len(tab.Cursors),
		)
	}

	c := tab.Cursors[n-1]

	return fmt.Sprintf("%v", *c), nil
}
