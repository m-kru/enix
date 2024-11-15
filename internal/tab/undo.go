package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
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

	// Currently only slice of actions are pushed to the stack.
	as, ok := act.Action.(action.Actions)
	if !ok {
		return
	}

	tab.undoActions(as)

	tab.RedoStack.Push(act.Action.Reverse(), curs)

	tab.Cursors = act.Cursors

	for _, a := range as {
		for _, m := range tab.Marks {
			m.Inform(a)
		}
	}
}

func (tab *Tab) undoActions(acts action.Actions) {
	for _, act := range acts {
		switch a := act.(type) {
		case *action.LineDown:
			tab.undoLineDown(a)
		case *action.LineUp:
			tab.undoLineUp(a)
		case *action.NewlineDelete:
			a.Line.Join(true)
			tab.LineCount--
		case *action.NewlineInsert:
			a.Line.InsertNewline(a.RuneIdx)
			tab.LineCount++
		case *action.RuneDelete:
			a.Line.DeleteRune(a.RuneIdx)
		case *action.RuneInsert:
			a.Line.InsertRune(a.Rune, a.RuneIdx)
		}
	}
}

func (tab *Tab) undoLineDown(ld *action.LineDown) {
	newFirstLine := ld.Line == tab.Lines

	line := ld.Line
	nextLine := line.Next
	if line.Prev != nil {
		line.Prev.Next = nextLine
	}
	nextLine.Prev = line.Prev
	line.Prev = nextLine
	line.Next = nextLine.Next
	nextLine.Next = line

	if newFirstLine {
		tab.Lines = nextLine
	}
}

func (tab *Tab) undoLineUp(ld *action.LineUp) {

}
