package exec

import (
	"fmt"
	"unicode/utf8"

	"github.com/m-kru/enix/internal/tab"
)

func Space(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"space: expected 0 args, provided %d", len(args),
		)
	}

	tab.InsertRune(' ')

	return nil
}

func Tab(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"tab: expected 0 args, provided %d", len(args),
		)
	}

	tab.InsertRune('\t')

	return nil
}

func Newline(args []string, tab *tab.Tab) error {
	if len(args) > 0 {
		return fmt.Errorf(
			"newline: expected 0 args, provided %d", len(args),
		)
	}

	tab.InsertNewline()

	return nil
}

func Rune(args []string, tab *tab.Tab) error {
	if len(args) != 1 {
		return fmt.Errorf(
			"rune: expected 1 arg, provided %d", len(args),
		)
	}

	runeCount := utf8.RuneCountInString(args[0])
	if runeCount != 1 {
		return fmt.Errorf(
			"rune: expected 1 rune, provided %d", runeCount,
		)
	}

	r, _ := utf8.DecodeRuneInString(args[0])
	if r == utf8.RuneError {
		return fmt.Errorf("rune: invalid rune provided")
	}

	tab.InsertRune(r)

	return nil
}
