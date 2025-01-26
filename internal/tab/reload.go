package tab

import (
	"os"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Reload() error {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var lastCur *cursor.Cursor
	if len(tab.Cursors) > 0 {
		lastCur = tab.Cursors[len(tab.Cursors)-1]
	} else {
		lastCur = tab.Selections[len(tab.Selections)-1].GetCursor()
	}

	var actions action.Actions

	bytes, err := os.ReadFile(tab.Path)
	if err != nil {
		return err
	}

	// Delete all lines
	cur := cursor.New(tab.Lines, 1, 0)
	for range tab.LineCount {
		act := cur.DeleteLine()
		actions = append(actions, act)
		tab.handleAction(act)

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	// Insert new lines
	line, lineCount := line.FromString(string(bytes))
	for range lineCount {
		act := cur.InsertLineBelow(line.String())
		actions = append(actions, act)
		tab.handleAction(act)
		cur.Down()
		line = line.Next
	}

	// Remove first extra line
	cur = cursor.New(tab.Lines, 1, 0)
	act := cur.DeleteLine()
	actions = append(actions, act)
	tab.handleAction(act)

	// Create new cursor
	lineNum := lastCur.LineNum
	if lineNum > tab.LineCount {
		lineNum = tab.LineCount
	}
	line = tab.Lines.Get(lineNum)

	rIdx := lastCur.RuneIdx
	lineRC := line.RuneCount()
	if rIdx > lineRC {
		rIdx = lineRC
	}

	newCur := cursor.New(line, lineNum, rIdx)
	tab.Cursors = []*cursor.Cursor{newCur}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	// Reset Undo and RedoCount
	tab.UndoCount = 0
	tab.RedoCount = 0

	return nil
}
