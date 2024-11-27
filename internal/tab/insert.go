package tab

import (
	"unicode"

	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Insert() {
	tab.State = "insert"
	tab.InsertActions = make(action.Actions, 0, 16)
	tab.PrevInsertCursors = cursor.Clone(tab.Cursors)
	tab.PrevInsertSelections = sel.Clone(tab.Selections)
}

func shouldInsertUndo(actions action.Action) bool {
	switch a := actions.(type) {
	case action.Actions:
		for _, subA := range a {
			undo := shouldInsertUndo(subA)
			if undo {
				return true
			}
		}
	case *action.NewlineInsert, *action.NewlineDelete:
		return true
	case *action.LineDown, *action.LineUp:
		return true
	case *action.RuneDelete:
		return unicode.IsSpace(a.Rune)
	case *action.RuneInsert:
		return unicode.IsSpace(a.Rune)
	}
	return false
}

func (tab *Tab) RxEventKeyInsert(ev *tcell.EventKey) {
	var act action.Action = nil
	updateView := true

	switch ev.Key() {
	case tcell.KeyRune:
		act = tab.insertRune(ev.Rune())
	case tcell.KeyTab:
		act = tab.insertRune('\t')
	case tcell.KeyBackspace2:
		act = tab.backspace()
	case tcell.KeyDelete:
		act = tab.delete()
	case tcell.KeyEnter:
		act = tab.insertNewline()
	default:
		c, _ := tab.Keys.ToCmd(ev)
		switch c.Name {
		case "esc":
			tab.State = "" // Go back to normal mode
			updateView = false
		case "view-down":
			tab.ViewDown()
			updateView = false
		case "view-left":
			tab.ViewLeft()
			updateView = false
		case "view-right":
			tab.ViewRight()
			updateView = false
		case "view-up":
			tab.ViewUp()
			updateView = false
		}
	}

	if act != nil {
		tab.InsertActions = append(tab.InsertActions, act)
	}

	if (shouldInsertUndo(act) || tab.State == "") && len(tab.InsertActions) > 0 {
		tab.UndoStack.Push(
			tab.InsertActions.Reverse(),
			tab.PrevInsertCursors,
			tab.PrevInsertSelections,
		)

		tab.InsertActions = make(action.Actions, 0, 16)
		tab.PrevInsertCursors = cursor.Clone(tab.Cursors)
		tab.PrevInsertSelections = sel.Clone(tab.Selections)
	}

	if updateView {
		tab.UpdateView()
	}
}

func (tab *Tab) InsertRune(r rune) {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.insertRune(r)

	if actions != nil {
		tab.UndoStack.Push(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) insertRune(r rune) action.Action {
	var a action.Action

	if tab.Cursors != nil {
		a = tab.insertRuneCursors(r)
	} else {
		a = nil
		// insert rune for selections unimplemented
	}

	if a != nil {
		tab.HasChanges = true
	}

	return a
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
		// insert newline for selections unimplemented
		return nil
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
