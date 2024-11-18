package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) InsertNewline() *action.NewlineInsert {
	rIdx := c.RuneIdx
	lineNum := c.LineNum
	line := c.Line
	newLine1, newLine2 := c.Line.InsertNewline(c.RuneIdx)

	c.Line = newLine2
	c.LineNum++
	c.RuneIdx = 0
	c.ColIdx = 1

	return &action.NewlineInsert{
		Line:     line,
		LineNum:  lineNum,
		RuneIdx:  rIdx,
		NewLine1: newLine1,
		NewLine2: newLine2,
	}
}

func (c *Cursor) InsertRune(r rune) *action.RuneInsert {
	c.Line.InsertRune(r, c.RuneIdx)
	c.RuneIdx++
	return &action.RuneInsert{Line: c.Line, Rune: r, RuneIdx: c.RuneIdx - 1}
}
