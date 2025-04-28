package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func MatchBracket(args []string, tab *tab.Tab) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf(
			"match-bracket: provided %d args, expected 0", len(args),
		)
	}

	tab.MatchBracket()

	return "", nil
}

func MatchCurly(args []string, tab *tab.Tab) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf(
			"match-curly: provided %d args, expected 0", len(args),
		)
	}

	tab.MatchCurly()

	return "", nil
}

func MatchParen(args []string, tab *tab.Tab) (string, error) {
	if len(args) != 0 {
		return "", fmt.Errorf(
			"match-paren: provided %d args, expected 0", len(args),
		)
	}

	tab.MatchParen()

	return "", nil
}
