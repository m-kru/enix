package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func SelRight(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"sel-right: provided %d args, expected 0", len(args),
		)
	}

	tab.SelRight()

	return nil
}
