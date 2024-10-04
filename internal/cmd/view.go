package cmd

import (
	"fmt"

	"github.com/m-kru/enix/internal/tab"
)

func ViewDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-down: expected 0 args, provided %d", len(args))
	}

	if tab.View.LastLine() >= tab.Lines.Count() {
		return nil
	}

	tab.View = tab.View.Down(1)

	return nil
}

func ViewUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-up: expected 0 args, provided %d", len(args))
	}

	tab.View = tab.View.Up(1)

	return nil
}

func ViewRight(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-right: expected 0 args, provided %d", len(args))
	}

	// - 3 because of:
	// 1. Space between line number and first line character.
	// 2. End of line character,
	// 3. One extra column, it simply looks better.
	lastCol := tab.View.LastColumn() + tab.LineNumWidth() - 3
	if lastCol >= tab.LastColumnIdx() {
		return nil
	}

	tab.View = tab.View.Right(1)

	return nil
}

func ViewLeft(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-right: expected 0 args, provided %d", len(args))
	}

	tab.View = tab.View.Left(1)

	return nil
}
