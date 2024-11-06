package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) InsertRune(r rune) *action.RuneInsert {
	c.Line.InsertRune(r, c.RuneIdx)
	c.RuneIdx++
	return &action.RuneInsert{Line: c.Line, RuneIdx: c.RuneIdx}
}

func (c *Cursor) InsertNewline() *action.NewlineInsert {
	rIdx := c.RuneIdx
	lineNum := c.LineNum
	newLine := c.Line.InsertNewline(c.RuneIdx)

	c.Line = newLine
	c.LineNum++
	c.RuneIdx = 0
	c.ColIdx = 1

	return &action.NewlineInsert{Line: newLine, LineNum: lineNum, RuneIdx: rIdx}
}
