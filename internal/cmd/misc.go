package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Esc(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"esc: expected 0 args, provided %d", len(args),
		)
	}

	return esc(tab)
}

func esc(tab *tab.Tab) error {
	if tab.Cursors != nil && tab.Cursors.Count() > 1 {
		tab.Cursors = tab.Cursors.Last()
		tab.Cursors.Prev = nil
	}

	return nil
}
