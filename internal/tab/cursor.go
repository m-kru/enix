package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Down() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.IntoCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Down()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) Left() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.IntoCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Left()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) Right() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.IntoCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Right()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) Up() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.IntoCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Up()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}
