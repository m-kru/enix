package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
	"unicode"
)

func Mark(args []string, tab *tab.Tab) (string, error) {
	if len(args) != 1 {
		return "", fmt.Errorf(
			"mark: provided %d args, expected 1", len(args),
		)
	}

	name := args[0]
	if unicode.IsDigit([]rune(name)[0]) {
		return "", fmt.Errorf(
			"mark: invalid name '%s', name must not start with a digit",
			name,
		)
	}

	tab.Mark(name)

	return fmt.Sprintf("mark '%s' created", name), nil
}
