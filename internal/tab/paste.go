package tab

import (
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/clip"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
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

	cursors := tab.Cursors
	if len(tab.Selections) > 0 {
		cursors = make([]*cursor.Cursor, 0, len(tab.Selections))
		for _, s := range tab.Selections {
			var c *cursor.Cursor
			if s.CursorOnRight() {
				c = s.GetCursor()
			} else {
				c = s.SpawnCursorOnRight()
			}
			cursors = append(cursors, c)
		}
	}

	if strings.HasSuffix(text, "\n") {
		cursors = cursor.Uniques(cursors, true)
	}

	actions := make(action.Actions, 0, len(cursors))
	newSels := make([]*sel.Selection, 0, len(cursors))

	for curIdx, cur := range cursors {
		startCur, endCur, acts := cur.Paste(text, false)

		tab.handleAction(acts)

		for _, c := range cursors[curIdx+1:] {
			c.Inform(acts)
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

	tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	tab.SearchCtx.Modified = true
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

	cursors := tab.Cursors
	if len(tab.Selections) > 0 {
		cursors = make([]*cursor.Cursor, 0, len(tab.Selections))
		for _, s := range tab.Selections {
			var c *cursor.Cursor
			if s.CursorOnLeft() {
				c = s.GetCursor()
			} else {
				c = s.SpawnCursorOnLeft()
			}
			cursors = append(cursors, c)
		}
	}

	if strings.HasSuffix(text, "\n") {
		cursors = cursor.Uniques(cursors, true)
	}

	actions := make(action.Actions, 0, len(cursors))
	newSels := make([]*sel.Selection, 0, len(cursors))

	for curIdx, cur := range cursors {
		startCur, endCur, acts := cur.PasteBefore(text, false)

		tab.handleAction(acts)

		for _, c := range cursors[curIdx+1:] {
			c.Inform(acts)
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

	tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	tab.SearchCtx.Modified = true
}
