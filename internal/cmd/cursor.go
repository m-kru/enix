package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strconv"
	"strings"
)

func CursorDown(args string, tab *tab.Tab) error {
	sstr := strings.Split(args, " ")
	if len(sstr) > 1 {
		return fmt.Errorf(
			"cursor-down: too many args provided (%d), expected at most 1", len(sstr),
		)
	}

	n, err := strconv.Atoi(sstr[0])
	if err != nil {
		return fmt.Errorf("cursor-down: %v", err)
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
