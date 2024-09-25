package cmd

import (
	"fmt"
	"strconv"

	"github.com/m-kru/enix/internal/tab"
)

func ViewDown(args []string, tab *tab.Tab) error {
	if len(args) > 1 {
		return fmt.Errorf("view-down: expected at most 1 arg, provided %d", len(args))
	}

	n := 1
	if len(args) > 0 {
		var err error
		n, err = strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("view-down: %v", err)
		}
	}

	tab.View = tab.View.Down(n)

	return nil
}
