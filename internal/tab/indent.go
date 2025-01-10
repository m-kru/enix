package tab

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Indent() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions
	if len(tab.Cursors) > 0 {
		actions = tab.indentCursors()
	} else {
		actions = tab.indentSelections()
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) indentCursors() action.Actions {
	// Create temporary cursors at the beginning of all lines with cursors.
	curs := cursor.Uniques(tab.Cursors, true)
	tmpCurs := make([]*cursor.Cursor, len(curs))
	for i, c := range curs {
		tmpCurs[i] = cursor.New(c.Line, c.LineNum, 0)
	}

	actions := make(action.Actions, 0, len(tmpCurs))

	for _, c := range tmpCurs {
		act := c.InsertString(tab.IndentStr)
		actions = append(actions, act)

		tab.handleAction(act)

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	// Create new cursors
	indentRC := utf8.RuneCountInString(tab.IndentStr)
	newCurs := make([]*cursor.Cursor, len(tab.Cursors))
	for i, c := range tab.Cursors {
		newCurs[i] = cursor.New(c.Line, c.LineNum, c.RuneIdx+indentRC)
	}
	tab.Cursors = newCurs

	return actions
}

func (tab *Tab) indentSelections() action.Actions {
	// Create temporary cursors at the beginning of all lines with selections.
	lines := sel.Lines(tab.Selections)
	tmpCurs := make([]*cursor.Cursor, len(lines))
	for i, l := range lines {
		// NOTE: The line number is incorrect as it is set to 0.
		// However, string insert action doesn't require line number.
		tmpCurs[i] = cursor.New(l, 0, 0)
	}

	actions := make(action.Actions, 0, len(tmpCurs))

	for _, c := range tmpCurs {
		act := c.InsertString(tab.IndentStr)
		actions = append(actions, act)

		tab.handleAction(act)

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	// Inform selections about actionns
	for _, s := range tab.Selections {
		s.Inform(actions, true)
	}

	return actions
}
