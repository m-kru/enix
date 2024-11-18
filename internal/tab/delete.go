package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
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

	for _, c := range tab.Cursors {
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

		tab.handleAction(act)

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
		tab.UndoStack.Push(actions.Reverse(), prevCurs, nil)
	}

	tab.HasChanges = true
}

func (tab *Tab) deleteSelections() {
	prevSels := sel.Clone(tab.Selections)
	actions := make(action.Actions, 0, len(tab.Selections))

	for _, s := range tab.Selections {
		act := s.Delete()

		if act == nil {
			continue
		}
		actions = append(actions, act)

		tab.handleAction(act)

		for _, s2 := range tab.Selections {
			if s2 != s {
				s2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), nil, prevSels)
	}

	tab.HasChanges = true
}

func (tab *Tab) Backspace() {
	if tab.Cursors != nil {
		tab.deleteCursors(true)
	} else {
		tab.deleteSelections()
	}
}
