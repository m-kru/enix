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

	tab.FindNext()

	return nil
}

func FindSelNext(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"find-sel-next: expected 0 args, provided %d", len(args),
		)
	}

	tab.FindSelNext()

	return nil
}
