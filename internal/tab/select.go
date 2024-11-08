package tab

import (
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) SelDown() {
	if len(tab.Cursors) > 0 {
		tab.selDownCursors()
	} else {
		tab.selDownSelections()
	}
}

func (tab *Tab) selDownCursors() {
	tab.Selections = sel.FromCursorsDown(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selDownSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Down()
	}

	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelLeft() {
	if len(tab.Cursors) > 0 {
		tab.selLeftCursors()
	} else {
		tab.selLeftSelections()
	}
}

func (tab *Tab) selLeftCursors() {
	tab.Selections = sel.FromCursorsLeft(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selLeftSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Left()
	}

	tab.Selections = sel.Prune(tab.Selections)
}

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
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Right()
	}

	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelUp() {
	if len(tab.Cursors) > 0 {
		tab.selUpCursors()
	} else {
		tab.selUpSelections()
	}
}

func (tab *Tab) selUpCursors() {
	tab.Selections = sel.FromCursorsUp(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selUpSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Up()
	}

	tab.Selections = sel.Prune(tab.Selections)
}
