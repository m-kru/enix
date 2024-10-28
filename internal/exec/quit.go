package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Quit(args []string, tab *tab.Tab, force bool) (*tab.Tab, error) {
	if len(args) > 0 {
		return tab, fmt.Errorf("quit: expected 0 args, provided %d", len(args))
	}

	if tab.HasChanges && !force {
		return tab, fmt.Errorf("quit: tab has unsaved changes")
	}

	newTab := tab.Prev
	if newTab != nil {
		newTab.Next = tab.Next
		if tab.Next != nil {
			tab.Next.Prev = newTab
		}
	} else if tab.Next != nil {
		newTab = tab.Next
		newTab.Prev = tab.Prev
		if tab.Prev != nil {
			tab.Prev.Next = newTab
		}
	}

	return newTab, nil
}
