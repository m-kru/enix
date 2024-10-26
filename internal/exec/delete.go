package exec

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

func Backspace(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("backspace: expected 0 args, provided %d", len(args))
	}

	tab.Backspace()

	return nil
}
