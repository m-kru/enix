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
			"end: expected 0 arguments, provided %d", len(sstr),
		)
	}

	tab.Cursors = &cursor.Cursor{Config: tab.Config, Line: tab.Lines.Last()}

	return nil
}
