package tab

import (
	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
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
	default:
		c, _ := tab.Keys.ToCmd(ev)
		switch c.Name {
		case "esc":
			tab.State = "" // Go back to normal mode
		case "view-down":
			tab.ViewDown()
		case "view-left":
			tab.ViewLeft()
		case "view-right":
			tab.ViewRight()
		case "view-up":
			tab.ViewUp()
		}

		return
	}

	tab.UpdateView()
}

func (tab *Tab) InsertRune(r rune) {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.insertRune(r)

	if actions != nil {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, prevSels)
		tab.HasChanges = true
	}
}

func (tab *Tab) insertRune(r rune) action.Action {
	if tab.Cursors != nil {
		return tab.insertRuneCursors(r)
	} else {
		panic("insert rune for selections unimplemented")
	}
}

func (tab *Tab) insertRuneCursors(r rune) action.Action {
	actions := make(action.Actions, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		act := c.InsertRune(r)

		actions = append(actions, act)

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	return actions
}

func (tab *Tab) InsertNewline() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.insertNewline()

	if actions != nil {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, prevSels)
		tab.HasChanges = true
	}
}

func (tab *Tab) insertNewline() action.Action {
	if tab.Cursors != nil {
		return tab.insertNewlineCursors()
	} else {
		panic("insert newline for selections unimplemented")
	}
}

func (tab *Tab) insertNewlineCursors() action.Action {
	actions := make(action.Actions, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		ni := c.InsertNewline()
		actions = append(actions, ni)

		if ni.Line == tab.Lines {
			tab.Lines = ni.NewLine1
		}

		tab.LineCount++

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(ni)
			}

		}
	}

	return actions
}
