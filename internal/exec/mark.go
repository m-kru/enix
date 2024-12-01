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
	r0 := []rune(name)[0]
	if unicode.IsDigit(r0) || r0 == '-' {
		return "", fmt.Errorf(
			"mark: invalid name '%s', name must not start with a digit or '-' rune",
			name,
		)
	}

	err := tab.Mark(name)
	if err != nil {
		return "", fmt.Errorf("mark: %v", err)
	}

	return fmt.Sprintf("mark '%s' created", name), nil
}
