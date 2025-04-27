package tab

import (
	"github.com/m-kru/enix/internal/cursor"
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

}
