package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strings"
)

func Esc(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 0 {
		return fmt.Errorf(
			"esc: expected 0 args, provided %d", len(sstr),
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
