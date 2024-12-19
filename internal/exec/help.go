package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/help"
	"github.com/m-kru/enix/internal/tab"
	"sort"
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

	var msg string

	if arg == "commands" {
		cmds := make([]string, 0, len(help.Commands))
		for c := range help.Commands {
			cmds = append(cmds, c)
		}
		sort.Strings(cmds)
		for _, c := range cmds {
			cmdHelp := help.Commands[c]
			before, _, _ := strings.Cut(cmdHelp, "\n")
			msg += before + "\n"
		}
	} else {
		var ok bool
		msg, ok = help.Topics[arg]
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
	}

	helpTab := tab.FromString(t.Keys, t.Frame, msg, "help-"+arg)
	t.Append(helpTab)

	return helpTab, nil
}
