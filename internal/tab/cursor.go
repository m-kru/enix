package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Down() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Down()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) End() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	tab.Cursors = []*cursor.Cursor{
		&cursor.Cursor{
			Line:    tab.Lines.Last(),
			LineNum: tab.LineCount,
		},
	}
}

func (tab *Tab) Left() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Left()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) LineEnd() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.LineEnd()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) LineStart() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.LineStart()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) PrevWordStart() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.PrevWordStart()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) Right() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Right()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) SpawnDown() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		nc := c.SpawnDown()

		if nc == nil {
			continue
		}

		newCurs = append(newCurs, nc)
	}

	if len(newCurs) > 0 {
		tab.Cursors = append(tab.Cursors, newCurs...)
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) SpawnUp() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		nc := c.SpawnUp()

		if nc == nil {
			continue
		}

		newCurs = append(newCurs, nc)
	}

	if len(newCurs) > 0 {
		tab.Cursors = append(tab.Cursors, newCurs...)
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) Up() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.Up()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) WordEnd() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.WordEnd()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) WordStart() {
	if len(tab.Cursors) == 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
	}

	for _, c := range tab.Cursors {
		c.WordStart()
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}
