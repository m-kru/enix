package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strings"
)

func Space(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"space: expected 0 args, provided %d", len(fields),
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

func Tab(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"tab: expected 0 args, provided %d", len(fields),
		)
	}

	tabCmd(tab)

	return nil
}

func tabCmd(tab *tab.Tab) {
	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertRune('\t')
		c = c.Next
	}
}

func Newline(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"newline: expected 0 args, provided %d", len(fields),
		)
	}

	newline(tab)

	return nil
}

func newline(tab *tab.Tab) {
	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertNewline()
		c = c.Next
	}
}
