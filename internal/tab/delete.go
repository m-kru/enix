package tab

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

// Delete deletes text under cursors or selections.
func (tab *Tab) Delete() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.delete()

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) delete() action.Actions {
	var actions action.Actions

	if tab.Cursors != nil {
		actions = tab.deleteCursors(false)
	} else {
		actions = tab.deleteSelections()
	}

	tab.SearchCtx.Modified = true

	return actions
}

func (tab *Tab) deleteCursors(backspace bool) action.Actions {
	actions := make(action.Actions, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		var act action.Action
		if backspace {
			if c.WithinIndent() {
				n := utf8.RuneCountInString(tab.IndentStr)
				as := make(action.Actions, 0, n)
				if c.RuneIdx == 0 {
					n = 1
				} else if c.RuneIdx < n {
					n = c.RuneIdx
				} else if c.RuneIdx%n != 0 {
					n = c.RuneIdx % n
				}
				for range n {
					a := c.Backspace()
					if a != nil {
						as = append(as, a)
					}
				}
				act = as
			} else {
				act = c.Backspace()
			}
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

	return actions
}

func (tab *Tab) deleteSelections() action.Actions {
	tab.Cursors = make([]*cursor.Cursor, 0, len(tab.Selections))

	actions := make(action.Actions, 0, len(tab.Selections))

	for i, s := range tab.Selections {
		act := s.Delete()

		if len(act) > 0 {
			actions = append(actions, act)
			tab.handleAction(act)
		}

		for _, c := range tab.Cursors {
			c.Inform(act)
		}

		// Selections are going to be destroyed anyway, so inform only
		// unprocessed selections.
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(act, true)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		// Create cursor from the first selection rune.
		c := cursor.New(s.Line, s.LineNum, s.StartRuneIdx)
		tab.Cursors = append(tab.Cursors, c)
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
	tab.Selections = nil

	return actions
}

func (tab *Tab) Backspace() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.backspace()

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) backspace() action.Actions {
	var actions action.Actions

	if tab.Cursors != nil {
		actions = tab.deleteCursors(true)
	} else {
		actions = tab.backspaceSelections()
	}

	tab.SearchCtx.Modified = true

	return actions
}

// NOTE: Currently backspace for the selection always deletes a rune before the selection.
// However, it is possible to implement it in such a way, that:
//  1. It deletes a rune before the selection if the cursor is on the left side of the selection.
//  2. It deletes a rune within the selection if the cursor is on the right side of the selection.
//
// So far, there was simply no use case for this more complex behavior.
func (tab *Tab) backspaceSelections() action.Actions {
	// Create cursors for deletions
	curs := make([]*cursor.Cursor, 0, len(tab.Selections))
	for _, s := range tab.Selections {
		line := s.Line
		lineNum := s.LineNum
		runeIdx := s.StartRuneIdx

		if lineNum == 1 && runeIdx == 0 {
			continue
		}

		if runeIdx == 0 {
			// Go to the previous line
			line = line.Prev
			lineNum--
			runeIdx = line.RuneCount()
		} else {
			runeIdx--
		}

		c := cursor.New(line, lineNum, runeIdx)

		curs = append(curs, c)
	}

	actions := make(action.Actions, 0, len(tab.Selections))

	for cIdx, c := range curs {
		act := c.Delete()

		if act == nil {
			continue
		}
		actions = append(actions, act)

		tab.handleAction(act)

		for _, c := range curs[cIdx+1:] {
			c.Inform(act)
		}

		for _, s := range tab.Selections {
			s.Inform(act, true)
		}

	}

	tab.Selections = sel.Prune(tab.Selections)

	return actions
}
