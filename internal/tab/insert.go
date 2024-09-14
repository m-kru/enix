package tab

import (
	"github.com/gdamore/tcell/v2"
)

func (t *Tab) RxEventKey(ev *tcell.EventKey) {
	if ev.Key() == tcell.KeyRune {
		//tab.InsertRune()
	}

	name, _ := t.Keys.ToCmd(ev)
	switch name {
	case "esc":
		t.InInsertMode = false
	}
}

func (t *Tab) InsertRune(r rune) {
	c := t.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertRune(r)
		c = c.Next
	}
}

func (t *Tab) InsertNewline() {
	c := t.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertNewline()
		c = c.Next
	}
}
