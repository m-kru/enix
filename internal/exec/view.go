package exec

import (
	"fmt"

	"github.com/m-kru/enix/internal/tab"
)

func ViewCenter(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-center: expected 0 args, provided %d", len(args))
	}

	tab.ViewCenter()

	return nil
}

func ViewDown(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-down: expected 0 args, provided %d", len(args))
	}

	tab.ViewDown()

	return nil
}

func ViewDownHalf(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-down-half: expected 0 args, provided %d", len(args))
	}

	tab.ViewDownHalf()

	return nil
}

func ViewEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-end: expected 0 args, provided %d", len(args))
	}

	tab.ViewEnd()

	return nil
}

func ViewUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-up: expected 0 args, provided %d", len(args))
	}

	tab.ViewUp()

	return nil
}

func ViewUpHalf(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-up-half: expected 0 args, provided %d", len(args))
	}

	tab.ViewUpHalf()

	return nil
}

func ViewRight(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-right: expected 0 args, provided %d", len(args))
	}

	tab.ViewRight()

	return nil
}

func ViewStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-start: expected 0 args, provided %d", len(args))
	}

	tab.ViewStart()

	return nil
}

func ViewLeft(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf("view-right: expected 0 args, provided %d", len(args))
	}

	tab.ViewLeft()

	return nil
}
