package tab

import "github.com/m-kru/enix/internal/mark"

func (tab *Tab) Mark(name string) {
	var m mark.Mark

	if tab.Cursors != nil {
		m = mark.NewCursorMark(tab.Cursors)
	} else {
		// Selection mark are currently unimplemented.
		// To implement them correctly it is required that
		// selection Inform supports all actions.
	}

	tab.Marks[name] = m
}
