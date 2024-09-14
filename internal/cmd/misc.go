package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strings"
	"unicode/utf8"
)

func Esc(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"esc: expected 0 args, provided %d", len(fields),
		)
	}

	return esc(tab)
}

func esc(tab *tab.Tab) error {
	if tab.Cursors != nil && tab.Cursors.Count() > 1 {
		tab.Cursors = tab.Cursors.Last()
		tab.Cursors.Prev = nil
	}

	return nil
}

func Rune(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) != 1 {
		return fmt.Errorf(
			"rune: expected 1 arg, provided %d", len(fields),
		)
	}

	runeCount := utf8.RuneCountInString(fields[0])
	if runeCount != 1 {
		return fmt.Errorf(
			"rune: expected 1 rune, provided %d", runeCount,
		)
	}

	r, _ := utf8.DecodeRuneInString(fields[0])
	if r == utf8.RuneError {
		return fmt.Errorf("rune: invalid rune provided")
	}

	insertRune(tab, r)

	return nil
}

func insertRune(tab *tab.Tab, r rune) {
	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertRune(r)
		c = c.Next
	}
}
