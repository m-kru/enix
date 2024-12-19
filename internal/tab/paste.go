package tab

import (
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/clip"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Paste() {
	text := clip.Read()
	if len(text) == 0 {
		return
	}

	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions
	if len(tab.Cursors) > 0 {
		actions = tab.pasteCursors(text)
	} else {
		actions = tab.pasteSelections(text)
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) pasteLineBased(text string, curs []*cursor.Cursor) action.Actions {
	lines, lineCount := line.FromString(text[0 : len(text)-1])

	newSels := make([]*sel.Selection, 0, len(curs))
	actions := make(action.Actions, 0, len(curs))

	for _, cur := range curs {
		var startCur *cursor.Cursor

		// Move cursor to the first column.
		// LineStart is called 2 times, because there might be spaces at the line start.
		cur.LineStart()
		cur.LineStart()

		acts := make(action.Actions, 0, lineCount)

		line := lines
		for {
			if line == nil {
				break
			}

			act := cur.InsertLineBelow(line.String())
			acts = append(acts, act)

			tab.handleAction(act)

			cur.Down()

			if startCur == nil {
				startCur = cur.Clone()
			}

			for _, cur2 := range curs {
				if cur2 != cur {
					cur2.Inform(act)
				}
			}

			for _, m := range tab.Marks {
				m.Inform(act)
			}

			line = line.Next
		}

		actions = append(actions, acts)

		cur.LineEnd()
		newSels = append(newSels, sel.FromTo(startCur, cur))
	}

	tab.Cursors = nil
	tab.Selections = newSels

	return actions
}

func (tab *Tab) pasteCursors(text string) action.Actions {
	var actions action.Actions
	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteCursorsLineBased(text)
	} else {
		return nil
	}

	return actions
}

func (tab *Tab) pasteCursorsLineBased(text string) action.Actions {
	curs := cursor.Uniques(tab.Cursors, true)

	return tab.pasteLineBased(text, curs)
}

func (tab *Tab) pasteSelections(text string) action.Actions {
	var actions action.Actions
	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteSelectionsLineBased(text)
	} else {
		return nil
	}

	return actions
}

func (tab *Tab) pasteSelectionsLineBased(text string) action.Actions {
	selCurs := make([]*cursor.Cursor, 0, len(tab.Selections))
	for _, s := range tab.Selections {
		selCurs = append(selCurs, s.SpawnCursorOnRight())
	}

	curs := cursor.Uniques(selCurs, true)

	return tab.pasteLineBased(text, curs)
}
