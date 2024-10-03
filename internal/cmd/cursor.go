package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/tab"
	"strconv"
	"strings"
)

func Down(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"down: provided %d args, expected 0", len(args),
		)
	}

	down(tab)

	return nil
}

func down(tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		c.Down()
		c = c.Next
	}

	tab.Cursors.Prune()
}

func Left(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"left: provided %d args, expected 0", len(args),
		)
	}

	left(tab)

	return nil
}

func left(tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		c.Left()
		c = c.Next
	}

	tab.Cursors.Prune()
}

func Right(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"right: provided %d args, expected 0", len(args),
		)
	}

	right(tab)

	return nil
}

func right(tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		c.Right()
		c = c.Next
	}

	tab.Cursors.Prune()
}

func Up(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"up: provided %d args, expected 0", len(args),
		)
	}

	up(tab)

	return nil
}

func up(tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		c.Up()
		c = c.Next
	}

	tab.Cursors.Prune()
}

func End(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"end: expected 0 args, provided %d", len(args),
		)
	}

	tab.Cursors = &cursor.Cursor{Config: tab.Config, Line: tab.Lines.Last()}

	return nil
}

func Goto(args []string, tab *tab.Tab) error {
	var err error
	line := 0
	col := 0

	if len(args) == 0 {
		return fmt.Errorf("goto: missing at least line number")
	} else if len(args) == 1 {
		if strings.Contains(args[0], ":") {
			lineStr, colStr, _ := strings.Cut(args[0], ":")
			line, err = strconv.Atoi(lineStr)
			if err != nil {
				return fmt.Errorf("goto: %v", err)
			}
			col, err = strconv.Atoi(colStr)
			if err != nil {
				return fmt.Errorf("goto: %v", err)
			}
		} else {
			line, err = strconv.Atoi(args[0])
			if err != nil {
				return fmt.Errorf("goto: %v", err)
			}
		}
	} else if len(args) == 2 {
		line, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("goto: %v", err)
		}
		col, err = strconv.Atoi(args[1])
		if err != nil {
			return fmt.Errorf("goto: %v", err)
		}
	} else {
		return fmt.Errorf("goto: expected at most 2 args, provided %d", len(args))
	}

	goTo(line, col, tab)

	return nil
}

func goTo(line, col int, tab *tab.Tab) {
	if line < 1 {
		line = 1
	}
	if col < 1 {
		col = 1
	}
	if line > tab.Lines.Count() {
		line = tab.Lines.Count()
	}

	l := tab.Lines.Get(line)
	if col > l.Len()+1 {
		col = l.Len() + 1
	}

	tab.Cursors = &cursor.Cursor{Config: tab.Config, Line: l, BufIdx: col - 1, Idx: col - 1}
}

func PrevWordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"prev-word-start: expected 0 args, provided %d", len(args),
		)
	}

	return prevWordStart(tab)
}

func prevWordStart(tab *tab.Tab) error {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}

		c.PrevWordStart()

		c = c.Next
	}

	tab.Cursors.Prune()

	return nil
}

func WordEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"word-end: expected 0 args, provided %d", len(args),
		)
	}

	return wordEnd(tab)
}

func wordEnd(tab *tab.Tab) error {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}

		c.WordEnd()

		c = c.Next
	}

	tab.Cursors.Prune()

	return nil
}

func WordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"word-start: expected 0 args, provided %d", len(args),
		)
	}

	return wordStart(tab)
}

func wordStart(tab *tab.Tab) error {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}

		c.WordStart()

		c = c.Next
	}

	tab.Cursors.Prune()

	return nil
}

func SpawnDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"spawn-down: expected 0 args, provided %d", len(args),
		)
	}

	return spawnDown(tab)
}

func spawnDown(tab *tab.Tab) error {
	var newCursors *cursor.Cursor
	var lastNewCursor *cursor.Cursor

	c := tab.Cursors
	for {
		nc := c.SpawnDown()

		if nc != nil {
			if newCursors == nil {
				newCursors = nc
				lastNewCursor = nc
			} else {
				lastNewCursor.Next = nc
				nc.Prev = lastNewCursor
				lastNewCursor = nc
			}
		}

		if c.Next == nil {
			break
		}
		c = c.Next
	}

	c.Next = newCursors
	if newCursors != nil {
		newCursors.Prev = c
	}

	tab.Cursors.Prune()

	return nil
}

func SpawnUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"spawn-up: expected 0 args, provided %d", len(args),
		)
	}

	return spawnUp(tab)
}

func spawnUp(tab *tab.Tab) error {
	var newCursors *cursor.Cursor
	var lastNewCursor *cursor.Cursor

	c := tab.Cursors
	for {
		nc := c.SpawnUp()

		if nc != nil {
			if newCursors == nil {
				newCursors = nc
				lastNewCursor = nc
			} else {
				lastNewCursor.Next = nc
				nc.Prev = lastNewCursor
				lastNewCursor = nc
			}
		}

		if c.Next == nil {
			break
		}
		c = c.Next
	}

	c.Next = newCursors
	if newCursors != nil {
		newCursors.Prev = c
	}

	tab.Cursors.Prune()

	return nil
}

func LineStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-start: expected 0 args, provided %d", len(args),
		)
	}

	return lineStart(tab)
}

func lineStart(tab *tab.Tab) error {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}

		c.LineStart()

		c = c.Next
	}

	tab.Cursors.Prune()

	return nil
}

func LineEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-end: expected 0 args, provided %d", len(args),
		)
	}

	return lineEnd(tab)
}

func lineEnd(tab *tab.Tab) error {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}

		c.LineEnd()

		c = c.Next
	}

	tab.Cursors.Prune()

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
