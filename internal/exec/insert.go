package exec

import (
	"fmt"
	"strconv"
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

func InsertRune(args []string, tab *tab.Tab) error {
	if len(args) != 3 {
		return fmt.Errorf(
			"insert-rune: expected 3 args, provided %d", len(args),
		)
	}

	lineNum, err := strconv.Atoi(args[0])
	if err != nil {
		return fmt.Errorf("insert-rune: %v", err)
	}
	if lineNum < 1 {
		return fmt.Errorf(
			"insert-rune: line number must be positive, current value %d", lineNum,
		)
	}

	col, err := strconv.Atoi(args[1])
	if err != nil {
		return fmt.Errorf("insert-rune: %v", err)
	}
	if lineNum < 1 {
		return fmt.Errorf(
			"insert-rune: column number must be positive, current value %d", col,
		)
	}

	runeCount := utf8.RuneCountInString(args[2])
	if runeCount != 1 {
		return fmt.Errorf(
			"insert-rune: expected 1 rune, provided %d", runeCount,
		)
	}

	r, _ := utf8.DecodeRuneInString(args[2])
	if r == utf8.RuneError {
		return fmt.Errorf("insert-rune: invalid rune provided")
	}

	err = tab.InsertRuneAtPosition(lineNum, col, r)
	if err != nil {
		return fmt.Errorf("insert-rune: %v", err)
	}

	return nil
}
