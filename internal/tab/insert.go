package tab

import (
	"fmt"
	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/action"
)

func (tab *Tab) RxEventKeyInsert(ev *tcell.EventKey) {
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

	c, _ := tab.Keys.ToCmd(ev)
	switch c.Name {
	case "esc":
		tab.State = "" // Go back to normal mode
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
	for _, c := range tab.Cursors {
		act := c.InsertRune(r)

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
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
	for _, c := range tab.Cursors {
		act := c.InsertNewline()

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}

		}
	}
}

func (tab *Tab) InsertRuneAtPosition(lineNum int, col int, r rune) error {
	lineCount := tab.Lines.Count()
	if lineNum > lineCount {
		return fmt.Errorf(
			"can't insert rune at line %d, tab has %d lines", lineNum, lineCount,
		)
	}

	line := tab.Lines.Get(lineNum)
	if col-1 > line.RuneCount() {
		return fmt.Errorf(
			"can't insert rune at index %d, line %d has %d runes",
			col, lineNum, line.RuneCount(),
		)
	}
	col--

	line.InsertRune(r, col)

	act := &action.RuneInsert{Line: line, Idx: col}
	for _, c := range tab.Cursors {
		c.Inform(act)
	}

	for _, m := range tab.Marks {
		m.Inform(act)
	}

	return nil
}
