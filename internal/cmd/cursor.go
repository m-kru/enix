package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strconv"
	"strings"
)

func CursorDown(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 1 {
		return fmt.Errorf(
			"cursor-down: provided %d args, expected at most 1", len(sstr),
		)
	}

	n := 1
	var err error
	if len(sstr) > 0 {
		n, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("cursor-down: %v", err)
		}
	}

	cursorDown(n, tab)

	return nil
}

func cursorDown(n int, tab *tab.Tab) {
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

func CursorUp(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 1 {
		return fmt.Errorf(
			"cursor-up: provided %d args, expected at most 1", len(sstr),
		)
	}

	n := 1
	var err error
	if len(sstr) > 0 {
		n, err = strconv.Atoi(sstr[0])
		if err != nil {
			return fmt.Errorf("cursor-up: %v", err)
		}
	}

	cursorUp(n, tab)

	return nil
}

func cursorUp(n int, tab *tab.Tab) {
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
