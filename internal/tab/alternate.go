package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
)

func (tab *Tab) Join() {
	if len(tab.Cursors) > 0 {
		tab.joinCursors()
	} else {
		tab.joinSelections()
	}
}

func (tab *Tab) joinCursors() {
	prevCurs := cursor.Clone(tab.Cursors)
	actions := make(action.Actions, 0, len(tab.Cursors))

	// Join lines only once, even if there are multiple cursors in the line.
	curs := cursor.Uniques(tab.Cursors, true)

	for _, c := range curs {
		act := c.Join()
		if act == nil {
			continue
		}
		actions = append(actions, act)

		tab.LineCount--

		nd := act.(*action.NewlineDelete)
		if nd.Line1 == tab.Lines {
			tab.Lines = nd.NewLine
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

func (tab *Tab) joinSelections() {
	panic("unimplemented")
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
