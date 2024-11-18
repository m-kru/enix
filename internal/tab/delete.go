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
	tab.Cursors = make([]*cursor.Cursor, 0, len(tab.Selections))

	prevSels := sel.Clone(tab.Selections)
	actions := make(action.Actions, 0, len(tab.Selections))

	for i, s := range tab.Selections {
		act := s.Delete()

		if act == nil {
			continue
		}
		actions = append(actions, act)

		tab.handleAction(act)

		for _, c := range tab.Cursors {
			c.Inform(act)
		}

		// Selections are going to be destroyed anyway, so inform only
		// unprocessed selections.
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		// Create cursor from the first selection rune.
		c := &cursor.Cursor{
			Line:    s.Line,
			LineNum: s.LineNum,
			ColIdx:  s.Line.ColumnIdx(s.StartRuneIdx),
			RuneIdx: s.StartRuneIdx,
		}
		tab.Cursors = append(tab.Cursors, c)
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
	tab.Selections = nil

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
