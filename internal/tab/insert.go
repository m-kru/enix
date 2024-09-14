package tab

import (
	"github.com/gdamore/tcell/v2"
)

func (t *Tab) RxEventKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyRune:
		t.InsertRune(ev.Rune())
	case tcell.KeyTab:
		panic("unimplemented tab")
	case tcell.KeyBackspace2:
		panic("unimplemented backspace2")
	case tcell.KeyDelete:
		panic("unimplemented delete")
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
