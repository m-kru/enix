package cmd

import (
	"fmt"
	"strconv"

	"github.com/m-kru/enix/internal/tab"
)

func ViewDown(args []string, tab *tab.Tab) error {
	if len(args) > 1 {
		return fmt.Errorf("view-down: expected at most 1 arg, provided %d", len(args))
	}

	n := 1
	if len(args) > 0 {
		var err error
		n, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("view-down: %v", err)
		}
	}

	if n < 1 {
		return fmt.Errorf("view-down: expected positive value, provided %d", n)
	}

	view := tab.View

	if n+view.LastLine() > tab.Lines.Count() {
		n = tab.Lines.Count() - view.LastLine()
	}

	tab.View = view.Down(n)

	return nil
}

func ViewUp(args []string, tab *tab.Tab) error {
	if len(args) > 1 {
		return fmt.Errorf("view-up: expected at most 1 arg, provided %d", len(args))
	}

	n := 1
	if len(args) > 0 {
		var err error
		n, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("view-up: %v", err)
		}
	}

	if n < 1 {
		return fmt.Errorf("view-up: expected positive value, provided %d", n)
	}

	tab.View = tab.View.Up(n)

	return nil
}

func ViewRight(args []string, tab *tab.Tab) error {
	if len(args) > 1 {
		return fmt.Errorf("view-right: expected at most 1 arg, provided %d", len(args))
	}

	n := 1
	if len(args) > 0 {
		var err error
		n, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("view-right: %v", err)
		}
	}

	if n < 1 {
		return fmt.Errorf("view-right: expected positive value, provided %d", n)
	}

	view := tab.View
	// - 3 because of:
	// 1. Space between line number and first line character.
	// 2. End of line character,
	// 3. One extra column, it simply looks better.
	lastCol := view.LastColumn() + tab.LineNumWidth() - 3
	if n+lastCol > tab.LastColumnIdx() {
		n = tab.LastColumnIdx() - lastCol
	}

	tab.View = view.Right(n)

	return nil
}
