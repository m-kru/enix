package tab

import (
	"fmt"
	"unicode"

	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Insert() {
	tab.State = "insert"
	tab.InsertActions = make(action.Actions, 0, 16)
	tab.PrevInsertCursors = cursor.Clone(tab.Cursors)
	tab.PrevInsertSelections = sel.Clone(tab.Selections)
}

func (tab *Tab) InsertLineBelow() error {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions

	if len(tab.Cursors) > 0 {
		actions = tab.insertLineBelowCursors()
	} else {
		return fmt.Errorf("unimplemented for selections")
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	tab.Insert()

	return nil
}

func (tab *Tab) insertLineBelowCursors() action.Actions {
	actions := make(action.Actions, 0, 4)
	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	curs := cursor.Uniques(tab.Cursors, true)

	for i, c := range curs {
		indent := c.Line.Indent()
		nnel := c.Line.GetNextNonEmpty()
		nnei := ""
		if nnel != nil {
			nnei = nnel.Indent()
		}
		if len(nnei) > len(indent) {
			indent = nnei
		}

		act := c.InsertLineBelow(indent)
		// act can't be nil here
		actions = append(actions, act)
		tab.handleAction(act)

		// Inform only remaining cursors, as we create new cursors anyway.
		for _, c2 := range curs[i+1:] {
			c2.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		newLine := c.Line.Next
		rIdx := newLine.RuneCount()
		newC := cursor.New(newLine, c.LineNum+1, rIdx)

		newCurs = append(newCurs, newC)
	}

	tab.Cursors = newCurs

	return actions
}

func (tab *Tab) InsertLineAbove() error {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions

	if len(tab.Cursors) > 0 {
		actions = tab.insertLineAboveCursors()
	} else {
		return fmt.Errorf("unimplemented for selections")
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	tab.Insert()

	return nil
}

func (tab *Tab) insertLineAboveCursors() action.Actions {
	actions := make(action.Actions, 0, 4)
	newCurs := make([]*cursor.Cursor, 0, len(tab.Cursors))

	curs := cursor.Uniques(tab.Cursors, true)

	for i, c := range curs {
		indent := c.Line.Indent()
		pnel := c.Line.GetPrevNonEmpty()
		pnei := ""
		if pnel != nil {
			pnei = pnel.Indent()
		}
		if len(pnei) > len(indent) {
			indent = pnei
		}

		act := c.InsertLineAbove(indent)
		// act can't be nil here
		actions = append(actions, act)
		tab.handleAction(act)

		// Inform only remaining cursors, as we create new cursors anyway.
		for _, c2 := range curs[i+1:] {
			c2.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		newLine := c.Line.Prev
		rIdx := newLine.RuneCount()
		newC := cursor.New(newLine, c.LineNum, rIdx)

		newCurs = append(newCurs, newC)
	}

	tab.Cursors = newCurs

	return actions
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

func (tab *Tab) RxEventKeyInsert(ev *tcell.EventKey) string {
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
		c, _ := cfg.InsertKeys.ToCmd(ev)
		switch c.Name {
		case "esc":
			// Trim spaces from empty lines.
			curs := cursor.Uniques(tab.Cursors, true)
			for _, c := range curs {
				if c.Line.HasOnlySpaces() {
					c.Line.Clear()
					c.LineStart()
				}
			}
			tab.Cursors = curs

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
		case "sh":
			return c.String()
		}
	}

	if act != nil {
		tab.InsertActions = append(tab.InsertActions, act)
	}

	if (shouldInsertUndo(act) || tab.State == "") && len(tab.InsertActions) > 0 {
		tab.undoPush(
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

	return ""
}

func (tab *Tab) InsertRune(r rune) {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.insertRune(r)

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) insertRune(r rune) action.Actions {
	if len(tab.Cursors) > 0 {
		return tab.insertRuneCursors(r)
	} else {
		return tab.insertRuneSelections(r)
	}
}

func (tab *Tab) insertRuneCursors(r rune) action.Actions {
	actions := make(action.Actions, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		var act action.Action
		if r == '\t' && c.WithinIndent() {
			act = c.InsertString(tab.IndentStr)
		} else {
			act = c.InsertRune(r)
		}

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

func (tab *Tab) insertRuneSelections(r rune) action.Actions {
	actions := make(action.Actions, 0, len(tab.Selections))

	for _, s := range tab.Selections {
		var act action.Action

		c := s.GetCursor()

		if r == '\t' && c.WithinIndent() {
			act = c.InsertString(tab.IndentStr)
		} else {
			act = c.InsertRune(r)
		}

		actions = append(actions, act)

		for _, s2 := range tab.Selections {
			s2.Inform(act, s2 != s)
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

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) insertNewline() action.Actions {
	if len(tab.Cursors) > 0 {
		return tab.insertNewlineCursors()
	} else {
		return tab.insertNewlineSelections()
	}
}

func (tab *Tab) insertNewlineCursors() action.Actions {
	actions := make(action.Actions, 0, len(tab.Cursors))

	for _, c := range tab.Cursors {
		act := c.InsertNewline(true)
		actions = append(actions, act)
		tab.handleAction(act)

		for _, c2 := range tab.Cursors {
			if c2 != c {
				c2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		// Remove spaces from "empty" lines.
		ni := act[0].(*action.NewlineInsert)
		skip := false
		for _, c2 := range tab.Cursors {
			if c2.Line == ni.NewLine1 {
				skip = true
				break
			}
		}
		if skip || !ni.NewLine1.HasOnlySpaces() {
			continue
		}
		ni.NewLine1.Clear()
	}

	return actions
}

func (tab *Tab) insertNewlineSelections() action.Actions {
	actions := make(action.Actions, 0, len(tab.Selections))

	for _, s := range tab.Selections {
		c := s.GetCursor()

		act := c.InsertNewline(true)
		actions = append(actions, act)
		tab.handleAction(act)

		for _, s2 := range tab.Selections {
			s2.Inform(act, s2 != s)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}

		// Remove spaces from "empty" lines.
		ni := act[0].(*action.NewlineInsert)
		skip := false
		for _, c2 := range tab.Cursors {
			if c2.Line == ni.NewLine1 {
				skip = true
				break
			}
		}
		if skip || !ni.NewLine1.HasOnlySpaces() {
			continue
		}
		ni.NewLine1.Clear()
	}

	return actions
}
