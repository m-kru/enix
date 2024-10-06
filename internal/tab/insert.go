package tab

import (
	"fmt"
	"github.com/gdamore/tcell/v2"
)

func (tab *Tab) RxEventKey(ev *tcell.EventKey) {
	switch ev.Key() {
	case tcell.KeyRune:
		tab.InsertRune(ev.Rune())
	case tcell.KeyTab:
		tab.InsertRune('\t')
	case tcell.KeyBackspace2:
		tab.Backspace()
	case tcell.KeyDelete:
		tab.Delete()
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

		for _, m := range tab.Marks {
			m.InformRuneInsert(c.Line, c.BufIdx)
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

func (tab *Tab) InsertRuneAtPosition(lineNum int, col int, r rune) error {
	lineCount := tab.Lines.Count()
	if lineNum > lineCount {
		return fmt.Errorf(
			"can't insert rune at line %d, tab has %d lines", lineNum, lineCount,
		)
	}

	line := tab.Lines.Get(lineNum)
	if col-1 > line.Len() {
		return fmt.Errorf(
			"can't insert rune at index %d, line %d has %d runes",
			col, lineNum, line.Len(),
		)
	}
	col--

	line.InsertRune(r, col)

	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		c.InformRuneInsert(line, col)
		c = c.Next
	}

	for _, m := range tab.Marks {
		m.InformRuneInsert(line, col)
	}

	return nil
}
