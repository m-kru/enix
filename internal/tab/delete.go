package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
)

// Delete deletes text under cursors or selections.
func (tab *Tab) Delete() {
	if tab.Cursors != nil {
		tab.deleteCursors(false)
	} else {
		tab.deleteSelections()
	}
}

func (tab *Tab) deleteCursors(backspace bool) {
	prevCurs := cursor.Clone(tab.Cursors)
	actions := make(action.Actions, 0, len(tab.Cursors))

	for i := 0; i < len(tab.Cursors); i++ {
		c := tab.Cursors[i]

		var act action.Action
		if backspace {
			act = c.Backspace()
		} else {
			act = c.Delete()
		}

		if act == nil {
			continue
		}

		actions = append(actions, act)

		if _, ok := act.(*action.NewlineDelete); ok {
			tab.LineCount--
		}

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), prevCurs)
	}

	tab.HasChanges = true
}

func (tab *Tab) deleteSelections() {
	panic("unimplemented")
}

func (tab *Tab) Backspace() {
	if tab.Cursors != nil {
		tab.deleteCursors(true)
	} else {
		tab.deleteSelections()
	}
}
