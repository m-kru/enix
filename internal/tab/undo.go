package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
	"github.com/m-kru/enix/internal/undo"
)

func (tab *Tab) Undo() {
	action := tab.UndoStack.Pop()

	if action == nil {
		return
	}

	tab.undo(action)
}

func (tab *Tab) undo(act *undo.Action) {
	curs := cursor.Clone(tab.Cursors)
	sels := sel.Clone(tab.Selections)

	// Currently only slice of actions are pushed to the stack.
	as, ok := act.Action.(action.Actions)
	if !ok {
		return
	}

	tab.undoActions(as)

	tab.RedoStack.Push(act.Action.Reverse(), curs, sels)

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
		}
	}
}

func (tab *Tab) undoLineDown(ld *action.LineDown) {
	newFirstLine := ld.Line == tab.Lines

	ld.Line.Down()

	if newFirstLine {
		tab.Lines = ld.Line.Prev
	}
}

func (tab *Tab) undoLineUp(lu *action.LineUp) {
	newFirstLine := lu.Line == tab.Lines.Next

	lu.Line.Up()

	if newFirstLine {
		tab.Lines = lu.Line
	}
}

func (tab *Tab) undoNewlineDelete(nd *action.NewlineDelete) {
	if nd.Line1.Prev != nil {
		nd.Line1.Prev.Next = nd.NewLine
	} else {
		tab.Lines = nd.NewLine
	}

	if nd.Line2.Next != nil {
		nd.Line2.Next.Prev = nd.NewLine
	}

	tab.LineCount--
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

	tab.LineCount++
}
