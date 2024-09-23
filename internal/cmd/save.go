package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Save(args []string, tab *tab.Tab, trim bool) error {
	if len(args) > 0 {
		return fmt.Errorf("save: expected 0 args, provided %d", len(args))
	}

	if trim {
		tab.Trim()
	}

	if tab.Config.SafeFileSave {
		return safeSave(tab)
	} else {
		return save(tab)
	}
}

func save(tab *tab.Tab) error {

	return nil
}

func safeSave(tab *tab.Tab) error {
	return nil
}
