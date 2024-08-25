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
