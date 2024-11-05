package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func SelDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-down: provided %d args, expected 0", len(args),
		)
	}

	tab.SelDown()

	return nil
}

func SelLeft(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-left: provided %d args, expected 0", len(args),
		)
	}

	tab.SelLeft()

	return nil
}

func SelRight(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-right: provided %d args, expected 0", len(args),
		)
	}

	tab.SelRight()

	return nil
}
