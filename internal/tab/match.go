package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) MatchParen() {
	if len(tab.Cursors) > 0 {
		tab.matchParenCursors()
	} else {
		tab.matchParenSelections()
	}
}

func (tab *Tab) matchParenCursors() {
	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	for _, cur := range tab.Cursors {
		newC := cur.MatchParen()
		if newC != nil {
			newCurs = append(newCurs, newC)
		}
	}

	if len(newCurs) > 0 {
		tab.Cursors = cursor.Prune(newCurs)
	}
}

func (tab *Tab) matchParenSelections() {
	newSels := make([]*sel.Selection, 0, len(tab.Selections))

	for _, s := range tab.Selections {
		newS := s.MatchParen()
		if newS != nil {
			newSels = append(newSels, newS)
		}
	}

	if len(newSels) > 0 {
		tab.Selections = sel.Prune(newSels)
	}
}
