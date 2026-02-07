package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Quit(args []string, tab *tab.Tab, force bool) (*tab.Tab, error) {
	if len(args) > 0 {
		return tab, fmt.Errorf("quit: expected 0 args, provided %d", len(args))
	}

	if tab.HasChanges() && !force {
		return tab, fmt.Errorf(
			"quit: tab has unsaved changes, use quit! or q! to force quit",
		)
	}

	return tab.Quit(), nil
}
