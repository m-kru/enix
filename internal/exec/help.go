package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/help"
	"github.com/m-kru/enix/internal/tab"
	"strings"
)

func Help(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) != 1 {
		return nil, fmt.Errorf("help: expected 1 arg, provided %d", len(args))
	}

	msg, ok := help.Topics[args[0]]
	if !ok {
		msg, ok = help.Commands[args[0]]
		if !ok {
			return nil, fmt.Errorf("help: entry for '%s' not found", args[0])
		}
		before, after, found := strings.Cut(msg, "\n")
		if found {
			msg = fmt.Sprintf("Synopsis:\n\n%s\n\nDescription:\n\n%s", before, after)
		} else {
			msg = fmt.Sprintf("Synopsis:\n\n%s", msg)
		}
	}

	helpTab := tab.FromString(t.Config, t.Colors, t.Keys, msg, "help-"+args[0])
	t.Append(helpTab)

	return helpTab, nil
}
