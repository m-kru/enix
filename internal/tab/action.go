package tab

import (
	"github.com/m-kru/enix/internal/action"
)

func (tab *Tab) handleAction(act action.Action) {
	switch a := act.(type) {
	case action.Actions:
		for _, subAct := range a {
			tab.handleAction(subAct)
		}
	case *action.NewlineDelete:
		tab.LineCount--
		if a.Line1 == tab.Lines {
			tab.Lines = a.NewLine
		}
	}
}
