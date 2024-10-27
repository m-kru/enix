package tab

import (
	"github.com/m-kru/enix/internal/cursor"

	"github.com/gdamore/tcell/v2"
)

func (tab *Tab) RxEventKeyReplace(ev *tcell.EventKey) {
	c, _ := tab.Keys.ToCmd(ev)
	switch c.Name {
	case "esc":
		tab.State = "" // Go back to normal mode
		return
	}

	tab.Delete()

	// Preserve cursors position
	var curs *cursor.Cursor
	if tab.Cursors != nil {
		curs = tab.Cursors.Clone()
	}

	switch ev.Key() {
	case tcell.KeyRune:
		tab.InsertRune(ev.Rune())
	case tcell.KeyTab:
		tab.InsertRune('\t')
	case tcell.KeyEnter:
		tab.InsertNewline()
	}

	tab.Cursors = curs

	tab.State = "" // Go back to normal mode
}
