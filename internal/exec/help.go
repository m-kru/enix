package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/help"
	"github.com/m-kru/enix/internal/tab"
	"strings"
)

func Help(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) > 1 {
		return nil, fmt.Errorf("help: expected 1 arg, provided %d", len(args))
	}

	arg := "help"

	if len(args) > 0 {
		arg = args[0]
	}

	msg, ok := help.Topics[arg]
	if !ok {
		msg, ok = help.Commands[arg]
		if !ok {
			return nil, fmt.Errorf("help: entry for '%s' not found", arg)
		}
		before, after, found := strings.Cut(msg, "\n")
		if found {
			msg = fmt.Sprintf("Synopsis:\n\n%s\n\nDescription:\n\n%s", before, after)
		} else {
			msg = fmt.Sprintf("Synopsis:\n\n%s", msg)
		}
	}

	helpTab := tab.FromString(t.Config, t.Colors, t.Keys, msg, "help-"+arg)
	t.Append(helpTab)

	return helpTab, nil
}
