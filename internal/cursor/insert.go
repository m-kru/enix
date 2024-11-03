package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) InsertRune(r rune) *action.RuneInsert {
	c.Line.InsertRune(r, c.RuneIdx)
	c.RuneIdx++
	return &action.RuneInsert{Line: c.Line, Idx: c.RuneIdx}
}

func (c *Cursor) InsertNewline() *action.NewlineInsert {
	rIdx := c.RuneIdx
	newLine := c.Line.InsertNewline(c.RuneIdx)

	c.Line = newLine
	c.RuneIdx = 0
	c.Idx = 0

	return &action.NewlineInsert{Line: newLine, Idx: rIdx}
}
