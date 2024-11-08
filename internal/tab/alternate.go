package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
)

func (tab *Tab) Join() {
	if len(tab.Cursors) > 0 {
		tab.joinCursors()
	} else {
		tab.joinSelections()
	}
}

func (tab *Tab) joinCursors() {
	for i := 0; i < len(tab.Cursors); i++ {
		c := tab.Cursors[i]

		act := c.Join()
		if act != nil {
			tab.LineCount--
		}

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	tab.HasChanges = true
}

func (tab *Tab) joinSelections() {
	panic("unimplemented")
}

func (tab *Tab) LineUp() {
	if len(tab.Cursors) > 0 {
		tab.lineUpCursors()
	} else {
		tab.lineUpSelections()
	}
}

func (tab *Tab) lineUpCursors() {
	// Move lines up only once, even if there are multiple cursors in the line.
	curs := make(map[*line.Line]*cursor.Cursor)
	for _, c := range tab.Cursors {
		if _, ok := curs[c.Line]; !ok {
			curs[c.Line] = c
		}
	}

	for _, c := range curs {
		newFirstLine := c.Line.Prev == tab.Lines

		act := c.LineUp()

		if newFirstLine {
			tab.Lines = c.Line
		}

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		tab.Cursors = cursor.Prune(tab.Cursors)
	}

	tab.HasChanges = true
}

func (tab *Tab) lineUpSelections() {
	panic("unimplemented")
}
