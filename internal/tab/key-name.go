package tab

import (
	enixTcell "github.com/m-kru/enix/internal/tcell"

	"github.com/gdamore/tcell/v2"
)

func (tab *Tab) RxEventKeyKeyName(ev *tcell.EventKey) {
	cmd, _ := tab.Keys.ToCmd(ev)
	switch cmd.Name {
	case "esc":
		tab.State = "" // Go back to normal mode
		return
	}

	name := enixTcell.EventKeyName(ev)

	c := tab.Cursors[0]
	_ = c.InsertString(name)
	c.InsertNewline()
	tab.LineCount++
	tab.UpdateView()
}
