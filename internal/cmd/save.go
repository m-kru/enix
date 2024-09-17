package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Save(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("save: expected 0 args, provided %d", len(args))
	}

	if tab.Config.SafeFileSave {
		return safeSave(tab)
	} else {
		panic("unsafe save not yet implemented")
	}

	return nil
}

func safeSave(tab *tab.Tab) error {

	return nil
}
