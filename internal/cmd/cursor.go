package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/tab"
	"strconv"
	"strings"
)

func Down(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 1 {
		return fmt.Errorf(
			"down: provided %d args, expected at most 1", len(sstr),
		)
	}

	n := 1
	var err error
	if len(sstr) > 0 {
		n, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("down: %v", err)
		}
	}

	down(n, tab)

	return nil
}

func down(n int, tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		for i := 0; i < n; i++ {
			c.Down()
		}
		c = c.Next
	}

	tab.Cursors.Prune()
}

func Left(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 1 {
		return fmt.Errorf(
			"left: provided %d args, expected at most 1", len(sstr),
		)
	}

	n := 1
	var err error
	if len(sstr) > 0 {
		n, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("left: %v", err)
		}
	}

	left(n, tab)

	return nil
}

func left(n int, tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		for i := 0; i < n; i++ {
			c.Left()
		}
		c = c.Next
	}

	tab.Cursors.Prune()
}

func Right(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 1 {
		return fmt.Errorf(
			"right: provided %d args, expected at most 1", len(sstr),
		)
	}

	n := 1
	var err error
	if len(sstr) > 0 {
		n, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("right: %v", err)
		}
	}

	right(n, tab)

	return nil
}

func right(n int, tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		for i := 0; i < n; i++ {
			c.Right()
		}
		c = c.Next
	}

	tab.Cursors.Prune()
}

func Up(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 1 {
		return fmt.Errorf(
			"up: provided %d args, expected at most 1", len(sstr),
		)
	}

	n := 1
	var err error
	if len(sstr) > 0 {
		n, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("up: %v", err)
		}
	}

	up(n, tab)

	return nil
}

func up(n int, tab *tab.Tab) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}
		for i := 0; i < n; i++ {
			c.Up()
		}
		c = c.Next
	}

	tab.Cursors.Prune()
}

func End(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 0 {
		return fmt.Errorf(
			"end: expected 0 args, provided %d", len(sstr),
		)
	}

	tab.Cursors = &cursor.Cursor{Config: tab.Config, Line: tab.Lines.Last()}

	return nil
}

func Goto(args string, tab *tab.Tab) error {
	var err error
	line := 0
	col := 0

	sstr := strings.Fields(args)
	if len(sstr) == 0 {
		return fmt.Errorf("goto: missing at least line number")
	} else if len(sstr) == 1 {
		if strings.Contains(sstr[0], ":") {
			lineStr, colStr, _ := strings.Cut(sstr[0], ":")
			line, err = strconv.Atoi(lineStr)
			if err != nil {
				return fmt.Errorf("goto: %v", err)
			}
			col, err = strconv.Atoi(colStr)
			if err != nil {
				return fmt.Errorf("goto: %v", err)
			}
		} else {
			line, err = strconv.Atoi(sstr[0])
			if err != nil {
				return fmt.Errorf("goto: %v", err)
			}
		}
	} else if len(sstr) == 2 {
		line, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("goto: %v", err)
		}
		col, err = strconv.Atoi(sstr[1])
		if err != nil {
			return fmt.Errorf("goto: %v", err)
		}
	} else {
		return fmt.Errorf("goto: expected at most 2 args, provided %d", len(sstr))
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
	if col > l.Len() {
		col = l.Len()
	}

	tab.Cursors = &cursor.Cursor{Config: tab.Config, Line: l, BufIdx: col - 1, Idx: col - 1}
}
