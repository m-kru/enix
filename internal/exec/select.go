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

func SelLine(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-line: provided %d args, expected 0", len(args),
		)
	}

	tab.SelLine()

	return nil
}

func SelPrevWordStart(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-prev-word-start: provided %d args, expected 0", len(args),
		)
	}

	tab.SelPrevWordStart()

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

func SelToTab(args []string, tab *tab.Tab) (*tab.Tab, error) {
	if len(args) > 1 {
		return tab, fmt.Errorf(
			"sel-to-tab: provided %d args, expected at most 1", len(args),
		)
	}

	path := "no-name"
	if len(args) == 1 {
		path = args[0]
	}

	idx := 1
	pathBase := path
	for {
		if !tab.Exists(path) {
			break
		}
		path = fmt.Sprintf("%s-%d", pathBase, idx)
		idx++
	}

	newTab := tab.SelToTab(path)
	tab.Append(newTab)

	return newTab, nil
}

func SelUp(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-up: provided %d args, expected 0", len(args),
		)
	}

	tab.SelUp()

	return nil
}

func SelWord(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-word: provided &D args, expected 0, len(args)",
		)
	}

	tab.SelWord()

	return nil
}

func SelWordEnd(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-word-end: provided %d args, expected 0", len(args),
		)
	}

	tab.SelWordEnd()

	return nil
}
