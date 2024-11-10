package tab

import (
	// enixTcell "github.com/m-kru/enix/internal/tcell"

	"github.com/gdamore/tcell/v2"
)

func (tab *Tab) RxEventKeyKeyName(ev *tcell.EventKey) {
	c, _ := tab.Keys.ToCmd(ev)
	switch c.Name {
	case "esc":
		tab.State = "" // Go back to normal mode
		return
	}

	// Uncomment when implemented
	//name := enixTcell.EventKeyName(ev)
	//tab.InsertString(name)
}
