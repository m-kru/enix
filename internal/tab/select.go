package tab

import (
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) SelRight() {
	if len(tab.Cursors) > 0 {
		tab.selRightCursors()
	} else {
		tab.selRightSelections()
	}
}

func (tab *Tab) selRightCursors() {
	tab.Selections = sel.FromCursorsRight(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selRightSelections() {
	for _, s := range tab.Selections {
		s.Right()
	}

	// TODO: Prune selections here
}
