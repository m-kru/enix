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
	// New selections
	sels := make([]*sel.Selection, 0, len(tab.Selections))
	actions := make(action.Actions, 0, len(tab.Cursors))

	for i, s := range tab.Selections {
		act, newS := s.Join()

		if len(act) == 0 {
			sels = append(sels, s.Clone())
			continue
		}

		actions = append(actions, act)

		tab.handleAction(act)

		// Selections are going to be changed, inform only unprocessed selections.
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(act)
		}

		// Inform new selections
		for _, s2 := range sels {
			s2.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		sels = append(sels, newS)
	}

	tab.Selections = sels

	return actions
}

func (tab *Tab) LineDown() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.lineDown()

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, prevSels)
		tab.HasChanges = true
	}
}

func (tab *Tab) lineDown() action.Actions {
	if len(tab.Cursors) > 0 {
		return tab.lineDownCursors()
	} else {
		return tab.lineDownSelections()
	}
}

func (tab *Tab) lineDownCursors() action.Actions {
	actions := make(action.Actions, 0, len(tab.Cursors))

	// Move lines down only once, even if there are multiple cursors in the line.
	curs := cursor.Uniques(tab.Cursors, false)

	for _, c := range curs {
		act := c.LineDown()
		if act == nil {
			continue
		}

		tab.handleAction(act)

		actions = append(actions, act)

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	return actions
}

func (tab *Tab) lineDownSelections() action.Actions {
	// New selections
	sels := make([]*sel.Selection, 0, len(tab.Selections))
	actions := make(action.Actions, 0, len(tab.Cursors))

	for i, s := range tab.Selections {
		act, newS := s.LineDown()

		if len(act) == 0 {
			sels = append(sels, s.Clone())
			continue
		}

		actions = append(actions, act)

		tab.handleAction(act)

		// Selections are going to be changed, inform only unprocessed selections.
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(act)
		}

		// Inform new selections
		for _, s2 := range sels {
			s2.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		sels = append(sels, newS)
	}

	tab.Selections = sels

	return actions
}

func (tab *Tab) LineUp() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.lineUp()

	if len(actions) > 0 {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, prevSels)
		tab.HasChanges = true
	}
}

func (tab *Tab) lineUp() action.Actions {
	if len(tab.Cursors) > 0 {
		return tab.lineUpCursors()
	} else {
		return tab.lineUpSelections()
	}
}

func (tab *Tab) lineUpCursors() action.Actions {
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

	return actions
}

func (tab *Tab) lineUpSelections() action.Actions {
	// New selections
	sels := make([]*sel.Selection, 0, len(tab.Selections))
	actions := make(action.Actions, 0, len(tab.Cursors))

	for i, s := range tab.Selections {
		act, newS := s.LineUp()

		if len(act) == 0 {
			sels = append(sels, s.Clone())
			continue
		}

		actions = append(actions, act)

		tab.handleAction(act)

		// Selections are going to be changed, inform only unprocessed selections.
		for _, s2 := range tab.Selections[i+1:] {
			s2.Inform(act)
		}

		// Inform new selections
		for _, s2 := range sels {
			s2.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		sels = append(sels, newS)
	}

	tab.Selections = sels

	return actions
}
