package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Quit(args []string, tab *tab.Tab, force bool) error {
	if len(args) > 0 {
		return fmt.Errorf("quit: expected 0 args, provided %d", len(args))
	}

	if tab.HasChanges && !force {
		return fmt.Errorf("quit: tab has unsaved changes")
	}

	return nil
}
