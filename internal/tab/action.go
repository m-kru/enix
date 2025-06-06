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
	case *action.LineDelete:
		tab.LineCount--
		if a.Line == tab.Lines {
			tab.Lines = a.Line.Next
		}
	case *action.LineDown:
		if a.Line == tab.Lines {
			tab.Lines = a.Line.Prev
		}
	case *action.LineInsert:
		tab.LineCount++
		if a.Line.Next == tab.Lines {
			tab.Lines = a.Line
		}
	case *action.LineUp:
		if a.Line.Prev == nil {
			tab.Lines = a.Line
		}
	case *action.NewlineDelete:
		tab.LineCount--
		if a.Line == tab.Lines {
			tab.Lines = a.NewLine
		}
	case *action.NewlineInsert:
		if a.Line == tab.Lines {
			tab.Lines = a.NewLine
		}
		tab.LineCount++
	}
}
