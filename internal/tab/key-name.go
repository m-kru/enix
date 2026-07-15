package tab

import (
	enixTcell "github.com/m-kru/enix/internal/tcell"

	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/cfg"
)

func (tab *Tab) RxEventKeyKeyName(ev *tcell.EventKey) {
	cmd, _ := cfg.KeysInsert.ToCmd(ev)
	switch cmd.Name {
	case "esc":
		tab.State = "" // Go back to normal mode
		return
	}

	name := enixTcell.EventKeyName(ev)

	c := tab.Cursors[0]
	_ = c.InsertString(name)
	c.InsertNewline(false)
	tab.LineCount++
	tab.UpdateView()
}
