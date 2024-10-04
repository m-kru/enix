package tab

import "github.com/m-kru/enix/internal/mark"

func (tab *Tab) Mark(name string) {
	var m mark.Mark

	if tab.Cursors != nil {
		m = mark.NewCursorMark(tab.Cursors)
	} else {

	}

	tab.Marks[name] = m
}
