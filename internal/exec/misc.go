package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Esc(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"esc: expected 0 args, provided %d", len(args),
		)
	}

	tab.Esc()

	return nil
}

func Trim(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"trim: expected 0 args, provided %d", len(args),
		)
	}

	tab.Trim()

	return nil
}

func Join(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"join: expected 0 args, provided %d", len(args),
		)
	}

	tab.Join()

	return nil
}

func KeyName(args []string, tab *tab.Tab) (string, error) {
	if len(args) > 0 {
		return "", fmt.Errorf(
			"key-name: expected 0 args, provided %d", len(args),
		)
	}

	return "insert a single key or key combo", nil
}

func LineDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-down: expected 0 args, provided %d", len(args),
		)
	}

	tab.LineDown()

	return nil
}

func LineUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"line-up: expected 0 args, provided %d", len(args),
		)
	}

	tab.LineUp()

	return nil
}

func Search(args []string, tab *tab.Tab) error {
	if len(args) != 1 {
		return fmt.Errorf(
			"search: expected 1 arg, provided %d", len(args),
		)
	}

	return tab.Search(args[0])
}
