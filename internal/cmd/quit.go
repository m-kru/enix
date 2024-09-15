package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Quit(tab *tab.Tab, force bool) error {
	if tab.HasChanges && !force {
		return fmt.Errorf("tab has unsaved changes")
	}

	return nil
}
