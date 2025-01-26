package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
	"github.com/m-kru/enix/internal/undo"
)

func (tab *Tab) Redo() {
	action := tab.RedoStack.Pop()

	if action == nil {
		return
	}

	curs := cursor.Clone(tab.Cursors)
	sels := sel.Clone(tab.Selections)

	tab.undo(action)
	if tab.RedoCount > 0 {
		tab.RedoCount--
	} else {
		tab.UndoCount++
	}

	tab.UndoStack.Push(action.Action.Reverse(), curs, sels)
}

func (tab *Tab) Undo() {
	action := tab.UndoStack.Pop()

	if action == nil {
		return
	}

	curs := cursor.Clone(tab.Cursors)
	sels := sel.Clone(tab.Selections)

	tab.undo(action)
	if tab.UndoCount > 0 {
		tab.UndoCount--
	} else {
		tab.RedoCount++
	}

	tab.RedoStack.Push(action.Action.Reverse(), curs, sels)
}

func (tab *Tab) undo(act *undo.Action) {
	// Currently only slice of actions are pushed to the stack.
	as, ok := act.Action.(action.Actions)
	if !ok {
		return
	}

	tab.undoActions(as)

	tab.Cursors = act.Cursors
	tab.Selections = act.Selections

	for _, a := range as {
		for _, m := range tab.Marks {
			m.Inform(a)
		}
	}
}

func (tab *Tab) undoActions(acts action.Actions) {
	for _, act := range acts {
		switch a := act.(type) {
		case action.Actions:
			tab.undoActions(a)
		case *action.LineDelete:
			tab.undoLineDelete(a)
		case *action.LineInsert:
			tab.undoLineInsert(a)
		case *action.LineDown:
			tab.undoLineDown(a)
		case *action.LineUp:
			tab.undoLineUp(a)
		case *action.NewlineDelete:
			tab.undoNewlineDelete(a)
		case *action.NewlineInsert:
			tab.undoNewlineInsert(a)
		case *action.RuneDelete:
			a.Line.DeleteRune(a.RuneIdx)
		case *action.RuneInsert:
			a.Line.InsertRune(a.Rune, a.RuneIdx)
		case *action.StringInsert:
			a.Line.InsertString(a.Str, a.StartRuneIdx)
		case *action.StringDelete:
			a.Line.DeleteString(a.StartRuneIdx, a.StartRuneIdx+a.RuneCount-1)
		}

		if _, ok := act.(action.Actions); !ok {
			tab.handleAction(act)
		}
	}
}

func (tab *Tab) undoLineDelete(ld *action.LineDelete) {
	if ld.Line.Prev != nil {
		ld.Line.Prev.Next = ld.Line.Next
	}
	if ld.Line.Next != nil {
		ld.Line.Next.Prev = ld.Line.Prev
	}
}

func (tab *Tab) undoLineInsert(li *action.LineInsert) {
	if li.Line.Prev != nil {
		li.Line.Prev.Next = li.Line
	}
	if li.Line.Next != nil {
		li.Line.Next.Prev = li.Line
	}
}

func (tab *Tab) undoLineDown(ld *action.LineDown) {
	ld.Line.Down()
}

func (tab *Tab) undoLineUp(lu *action.LineUp) {
	lu.Line.Up()
}

func (tab *Tab) undoNewlineDelete(nd *action.NewlineDelete) {
	if nd.Line.Prev != nil {
		nd.Line.Prev.Next = nd.NewLine
	} else {
		tab.Lines = nd.NewLine
	}

	if nd.NextLine.Next != nil {
		nd.NextLine.Next.Prev = nd.NewLine
	}
}

func (tab *Tab) undoNewlineInsert(ni *action.NewlineInsert) {
	if ni.Line.Prev != nil {
		ni.Line.Prev.Next = ni.NewLine1
	} else {
		tab.Lines = ni.NewLine1
	}

	if ni.Line.Next != nil {
		ni.Line.Next.Prev = ni.NewLine2
	}
}
