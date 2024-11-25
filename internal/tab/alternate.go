package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Join() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.join()

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, prevSels)
		tab.HasChanges = true
	}
}

func (tab *Tab) join() action.Actions {
	if len(tab.Cursors) > 0 {
		return tab.joinCursors()
	} else {
		return tab.joinSelections()
	}
}

func (tab *Tab) joinCursors() action.Actions {
	// Join lines only once, even if there are multiple cursors in the line.
	curs := cursor.Uniques(tab.Cursors, true)

	actions := make(action.Actions, 0, len(tab.Cursors))

	for _, c := range curs {
		act := c.Join()
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

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	return actions
}

func (tab *Tab) joinSelections() action.Actions {
	return nil
}

func (tab *Tab) LineDown() {
	if len(tab.Cursors) > 0 {
		tab.lineDownCursors()
	} else {
		tab.lineDownSelections()
	}
}

func (tab *Tab) lineDownCursors() {
	prevCurs := cursor.Clone(tab.Cursors)
	actions := make(action.Actions, 0, len(tab.Cursors))

	// Move lines down only once, even if there are multiple cursors in the line.
	curs := cursor.Uniques(tab.Cursors, false)

	for _, c := range curs {
		newFirstLine := c.Line == tab.Lines

		act := c.LineDown()
		if act == nil {
			continue
		}

		actions = append(actions, act)

		if newFirstLine {
			tab.Lines = c.Line.Prev
		}

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, nil)
	}

	tab.HasChanges = true
}

func (tab *Tab) lineDownSelections() {
	panic("unimplemented")
}

func (tab *Tab) LineUp() {
	if len(tab.Cursors) > 0 {
		tab.lineUpCursors()
	} else {
		tab.lineUpSelections()
	}
}

func (tab *Tab) lineUpCursors() {
	prevCurs := cursor.Clone(tab.Cursors)
	actions := make(action.Actions, 0, len(tab.Cursors))

	// Move lines up only once, even if there are multiple cursors in the line.
	curs := cursor.Uniques(tab.Cursors, true)

	for _, c := range curs {
		newFirstLine := c.Line.Prev == tab.Lines

		act := c.LineUp()
		if act == nil {
			continue
		}

		actions = append(actions, act)

		if newFirstLine {
			tab.Lines = c.Line
		}

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, nil)
	}

	tab.HasChanges = true
}

func (tab *Tab) lineUpSelections() {
	panic("unimplemented")
}
