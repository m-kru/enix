package cursor

import (
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

func (cur *Cursor) Paste(text string, addIndent bool, after bool) (*Cursor, *Cursor, action.Actions) {
	if strings.HasSuffix(text, "\n") {
		return cur.pasteLineBased(text, addIndent)
	}
	return cur.pasteRegular(text, addIndent, after)
}

func (cur *Cursor) pasteLineBased(text string, addIndent bool) (*Cursor, *Cursor, action.Actions) {
	if addIndent {
		indent := cur.Line.Indent()
		text = util.AddIndent(text, indent, true)
	}
	lines, lineCount := line.FromString(text[0 : len(text)-1])

	// Create cursor at the line start
	cur = New(cur.Line, cur.LineNum, 0)
	var startCur *Cursor

	actions := make(action.Actions, 0, lineCount)

	line := lines
	for line != nil {
		act := cur.InsertLineBelow(line.String())
		actions = append(actions, act)

		cur.Down()
		if startCur == nil {
			startCur = cur.Clone()
		}

		line = line.Next
	}

	cur.LineEnd()

	return startCur, cur, actions
}

func (cur *Cursor) pasteRegular(text string, addIndent bool, after bool) (*Cursor, *Cursor, action.Actions) {
	if addIndent {
		indent := cur.Line.Indent()
		text = util.AddIndent(text, indent, false)
	}
	lines, lineCount := line.FromString(text)

	if after {
		cur.Right()
	}

	startRuneIdx := cur.RuneIdx
	startCur := cur.Clone()

	actions := make(action.Actions, 0, 2*lineCount)

	line := lines
	for i := range lineCount {
		str := line.String()
		if str != "" {
			act := cur.InsertString(str)
			actions = append(actions, act)
		}

		if line.Next != nil {
			acts := cur.InsertNewline(false)
			actions = append(actions, acts)

			// Update start cursor if newline was inserted in the paste line.
			if i == 0 {
				ni := acts[0].(*action.NewlineInsert)
				startCur = New(ni.NewLine, cur.LineNum-1, startRuneIdx)
			}
		}

		line = line.Next
	}

	cur.Left()

	return startCur, cur, actions
}
