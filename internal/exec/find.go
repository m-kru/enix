package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func FindNext(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"find-next: expected 0 args, provided %d", len(args),
		)
	}

	tab.Find(true)

	return nil
}

func FindPrev(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"find-prev: expected 0 args, provided %d", len(args),
		)
	}

	tab.Find(false)

	return nil
}

func FindSelNext(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"find-sel-next: expected 0 args, provided %d", len(args),
		)
	}

	tab.FindSel(true)

	return nil
}

func FindSelPrev(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"find-sel-prev: expected 0 args, provided %d", len(args),
		)
	}

	tab.FindSel(false)

	return nil
}
