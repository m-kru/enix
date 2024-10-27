package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Open(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("open: expected at least 1 arg, provided 0")
	}

	var firstTab *tab.Tab

	for i, path := range args {
		newT := tab.Open(t.Config, t.Colors, t.Keys, path)
		t.Append(newT)

		if i == 0 {
			firstTab = newT
		}
	}

	return firstTab, nil
}
