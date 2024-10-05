package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Del(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("del: expected 0 args, provided %d", len(args))
	}

	tab.Delete()

	return nil
}
