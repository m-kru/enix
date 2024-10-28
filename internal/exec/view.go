package exec

import (
	"fmt"

	"github.com/m-kru/enix/internal/tab"
)

func ViewDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-down: expected 0 args, provided %d", len(args))
	}

	tab.ViewDown()

	return nil
}

func ViewUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-up: expected 0 args, provided %d", len(args))
	}

	tab.ViewUp()

	return nil
}

func ViewRight(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-right: expected 0 args, provided %d", len(args))
	}

	tab.ViewRight()

	return nil
}

func ViewLeft(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-right: expected 0 args, provided %d", len(args))
	}

	tab.ViewLeft()

	return nil
}
