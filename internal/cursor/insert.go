package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) InsertRune(r rune) *action.RuneInsert {
	c.Line.InsertRune(r, c.BufIdx)
	c.BufIdx++
	return &action.RuneInsert{Line: c.Line, Idx: c.BufIdx}
}

func (c *Cursor) InsertNewline() *action.NewlineInsert {
	line := c.Line
	bufIdx := c.BufIdx
	newLine := c.Line.InsertNewline(c.BufIdx)

	c.Line = newLine
	c.BufIdx = 0
	c.Idx = 0

	return &action.NewlineInsert{Line: line, Idx: bufIdx, NewLine: newLine}
}
