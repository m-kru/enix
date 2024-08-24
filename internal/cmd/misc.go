package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strings"
)

func Space(args string, tab *tab.Tab) error {
	sstr := strings.Fields(args)
	if len(sstr) > 0 {
		return fmt.Errorf(
			"space: provided %d args, expected 0", len(sstr),
		)
	}

	space(tab)

	return nil
}

func space(tab *tab.Tab) {
	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertRune(' ')
		c = c.Next
	}
}
