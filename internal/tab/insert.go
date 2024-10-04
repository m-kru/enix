package tab

import (
	"github.com/gdamore/tcell/v2"
)

func (tab *Tab) RxEventKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyRune:
		tab.InsertRune(ev.Rune())
	case tcell.KeyTab:
		tab.InsertRune('\t')
	case tcell.KeyBackspace2:
		panic("unimplemented backspace2")
	case tcell.KeyDelete:
		panic("unimplemented delete")
	case tcell.KeyEnter:
		tab.InsertNewline()
	}

	name, _ := tab.Keys.ToCmd(ev)
	switch name {
	case "esc":
		tab.InInsertMode = false
	}
}

func (tab *Tab) InsertRune(r rune) {
	if tab.Cursors != nil {
		tab.insertRuneCursors(r)
	}
}

func (tab *Tab) insertRuneCursors(r rune) {
	c0 := tab.Cursors // First cursor
	c := c0

	for {
		c2 := c0
		for {
			if c2 == nil {
				break
			}
			if c2 != c {
				c2.InformRuneInsert(c.Line, c.BufIdx)
			}
			c2 = c2.Next
		}

		c.InsertRune(r)

		c = c.Next
		if c == nil {
			break
		}
	}

	tab.HasChanges = true
}

func (tab *Tab) InsertNewline() {
	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		c.InsertNewline()
		c = c.Next
	}

	tab.HasChanges = true
}
