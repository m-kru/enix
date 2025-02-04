package exec

import (
	"fmt"

	"github.com/m-kru/enix/internal/tab"
)

func Save(args []string, tab *tab.Tab, trim bool) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("save: expected at most 1 arg, provided %d", len(args))
	}

	path := tab.Path
	if len(args) == 1 {
		path = args[0]
	}

	return tab.Save(path, trim)
}
