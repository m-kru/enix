package cmd

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"strings"
	"unicode/utf8"
)

func Space(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"space: expected 0 args, provided %d", len(fields),
		)
	}

	tab.InsertRune(' ')

	return nil
}

func Tab(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"tab: expected 0 args, provided %d", len(fields),
		)
	}

	tab.InsertRune('\t')

	return nil
}

func Newline(args string, tab *tab.Tab) error {
	fields := strings.Fields(args)
	if len(fields) > 0 {
		return fmt.Errorf(
			"newline: expected 0 args, provided %d", len(fields),
		)
	}

	tab.InsertNewline()

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

	tab.InsertRune(r)

	return nil
}
