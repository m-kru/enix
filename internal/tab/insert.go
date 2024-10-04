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
	} else {
		panic("insert rune for selections unimplemented")
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
	if tab.Cursors != nil {
		tab.insertNewlineCursors()
	} else {
		panic("insert newline for selections unimplemented")
	}
}

func (tab *Tab) insertNewlineCursors() {
	c0 := tab.Cursors // First cursor

	c := c0
	for {
		line := c.Line
		bufIdx := c.BufIdx
		newLine := c.InsertNewline()

		// Update line pointer for all cursors in the same line as c, but after c.
		c2 := c0
		for {
			if c2 == nil {
				break
			}

			if c2.Line == line && c2.BufIdx > bufIdx {
				c2.Line = newLine
				c2.BufIdx -= bufIdx
				c2.Idx = c2.BufIdx
			}

			c2 = c2.Next
		}

		c = c.Next
		if c == nil {
			break
		}
	}

	tab.HasChanges = true
}
