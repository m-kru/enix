package tab

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
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

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	// TODO: Delete all lines

	// TODO: Insert new lines

	// Create new cursor
	lineNum := lastCur.LineNum
	if lineNum > tab.LineCount {
		lineNum = tab.LineCount
	}
	line := tab.Lines.Get(lineNum)

	rIdx := lastCur.RuneIdx
	lineRC := line.RuneCount()
	if rIdx > lineRC {
		rIdx = lineRC
	}

	newCur := cursor.New(line, lineNum, rIdx)
	tab.Cursors = []*cursor.Cursor{newCur}

	return nil
}
