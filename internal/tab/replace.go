package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"

	"github.com/gdamore/tcell/v2"
)

func (tab *Tab) RxEventKeyReplace(ev *tcell.EventKey) {
	c, _ := cfg.InsertKeys.ToCmd(ev)
	switch c.Name {
	case "esc":
		tab.State = "" // Go back to normal mode
		return
	}

	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.delete()

	// Preserve cursors position
	var curs []*cursor.Cursor
	if tab.Cursors != nil {
		curs = cursor.Clone(tab.Cursors)
	}

	var insertActions action.Action
	switch ev.Key() {
	case tcell.KeyRune:
		insertActions = tab.insertRune(ev.Rune())
	case tcell.KeyTab:
		insertActions = tab.insertRune('\t')
	case tcell.KeyEnter:
		insertActions = tab.insertNewline()
	}

	if insertActions != nil {
		actions = append(actions, insertActions)
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	tab.Cursors = curs
	tab.State = "" // Go back to normal mode
}
