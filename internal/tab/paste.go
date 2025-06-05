package tab

import (
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/clip"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/sel"
	"github.com/m-kru/enix/internal/util"
)

func (tab *Tab) Paste(text string) {
	if len(text) == 0 {
		text = clip.Read()
		if len(text) == 0 {
			return
		}
	}

	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions
	if len(tab.Cursors) > 0 {
		actions = tab.pasteCursors(text, false)
	} else {
		actions = tab.pasteSelections(text)
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) pasteLineBased(text string, addIndent bool, curs []*cursor.Cursor) action.Actions {
	newSels := make([]*sel.Selection, 0, len(curs))
	actions := make(action.Actions, 0, len(curs))

	for curIdx, cur := range curs {
		startCur, endCur, acts := cur.Paste(text, addIndent, true)

		tab.handleAction(acts)

		for _, c := range curs[curIdx+1:] {
			c.Inform(acts)
		}

		for _, m := range tab.Marks {
			m.Inform(acts)
		}

		actions = append(actions, acts)

		newSels = append(newSels, sel.FromTo(startCur, endCur))
	}

	tab.Cursors = nil
	tab.Selections = newSels

	return actions
}

func (tab *Tab) pasteCursors(text string, addIndent bool) action.Actions {
	var actions action.Actions

	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteCursorsLineBased(text, addIndent)
	} else {
		actions = tab.pasteCursorsRegular(text, addIndent, true)
	}

	tab.SearchCtx.Modified = true

	return actions
}

func (tab *Tab) pasteCursorsLineBased(text string, addIndent bool) action.Actions {
	curs := cursor.Uniques(tab.Cursors, true)

	return tab.pasteLineBased(text, addIndent, curs)
}

func (tab *Tab) pasteCursorsRegular(text string, addIndent bool, after bool) action.Actions {
	actions := make(action.Actions, 0, len(tab.Cursors))
	newSels := make([]*sel.Selection, 0, len(tab.Cursors))

	for curIdx, cur := range tab.Cursors {
		startCur, endCur, acts := cur.Paste(text, addIndent, after)

		for _, cur2 := range tab.Cursors[curIdx+1:] {
			cur2.Inform(acts)
		}

		for _, m := range tab.Marks {
			m.Inform(acts)
		}

		for _, s := range newSels {
			s.Inform(acts, true)
		}

		actions = append(actions, acts)

		newSels = append(newSels, sel.FromTo(startCur, endCur))
	}

	tab.Cursors = nil
	tab.Selections = newSels

	return actions
}

func (tab *Tab) pasteSelections(text string) action.Actions {
	var actions action.Actions

	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteSelectionsLineBased(text)
	} else {
		return nil
	}

	tab.SearchCtx.Modified = true

	return actions
}

func (tab *Tab) pasteSelectionsLineBased(text string) action.Actions {
	selCurs := make([]*cursor.Cursor, 0, len(tab.Selections))
	for _, s := range tab.Selections {
		selCurs = append(selCurs, s.SpawnCursorOnRight())
	}

	curs := cursor.Uniques(selCurs, true)

	return tab.pasteLineBased(text, false, curs)
}

func (tab *Tab) PasteBefore(text string) {
	if len(text) == 0 {
		text = clip.Read()
		if len(text) == 0 {
			return
		}
	}

	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions
	if len(tab.Cursors) > 0 {
		actions = tab.pasteBeforeCursors(text, false)
	} else {
		actions = tab.pasteBeforeSelections(text)
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) pasteBeforeCursors(text string, addIndent bool) action.Actions {
	var actions action.Actions

	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteBeforeCursorsLineBased(text, addIndent)
	} else {
		actions = tab.pasteBeforeCursorsRegular(text, addIndent)
	}

	tab.SearchCtx.Modified = true

	return actions
}

func (tab *Tab) pasteBeforeLineBased(text string, addIndent bool, curs []*cursor.Cursor) action.Actions {
	newSels := make([]*sel.Selection, 0, len(curs))
	actions := make(action.Actions, 0, len(curs))

	var lines *line.Line
	var lineCount int
	for curIdx, cur := range curs {
		// Each cursors might be in a line with a different indent.
		if lines == nil || addIndent {
			t := text
			if addIndent {
				indent := cur.Line.Indent()
				t = util.AddIndent(text, indent, true)
			}
			lines, lineCount = line.FromString(t[0 : len(t)-1])
		}

		var endCur *cursor.Cursor

		// Create cursor at the line start
		cur = cursor.New(cur.Line, cur.LineNum, 0)

		acts := make(action.Actions, 0, lineCount)

		line := lines.Last()
		for line != nil {
			act := cur.InsertLineAbove(line.String())
			acts = append(acts, act)

			tab.handleAction(act)

			cur.Up()

			if endCur == nil {
				endCur = cur.Clone()
				endCur.LineEnd()
			} else {
				endCur.Inform(act)
			}

			for _, c := range curs[curIdx+1:] {
				c.Inform(act)
			}

			for _, m := range tab.Marks {
				m.Inform(act)
			}

			line = line.Prev
		}

		actions = append(actions, acts)

		newSels = append(newSels, sel.FromTo(cur, endCur))
	}

	tab.Cursors = nil
	tab.Selections = newSels

	return actions
}

func (tab *Tab) pasteBeforeCursorsLineBased(text string, addIndent bool) action.Actions {
	curs := cursor.Uniques(tab.Cursors, true)

	return tab.pasteBeforeLineBased(text, addIndent, curs)
}

func (tab *Tab) pasteBeforeCursorsRegular(text string, addIndent bool) action.Actions {
	return tab.pasteCursorsRegular(text, addIndent, false)
}

func (tab *Tab) pasteBeforeSelections(text string) action.Actions {
	var actions action.Actions

	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteBeforeSelectionsLineBased(text)
	} else {
		return nil
	}

	tab.SearchCtx.Modified = true

	return actions
}

func (tab *Tab) pasteBeforeSelectionsLineBased(text string) action.Actions {
	selCurs := make([]*cursor.Cursor, 0, len(tab.Selections))

	for _, s := range tab.Selections {
		c := cursor.New(s.Line, s.LineNum, 0)
		selCurs = append(selCurs, c)
	}

	curs := cursor.Uniques(selCurs, true)

	return tab.pasteBeforeLineBased(text, false, curs)
}
