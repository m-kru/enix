package cursor

import (
	"unicode/utf8"

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

func (c *Cursor) InsertString(s string) *action.StringInsert {
	c.Line.InsertString(s, c.RuneIdx)
	rIdx := c.RuneIdx
	rc := utf8.RuneCountInString(s)
	c.RuneIdx += rc
	return &action.StringInsert{
		Line:         c.Line,
		Str:          s,
		StartRuneIdx: rIdx,
		RuneCount:    rc,
	}
}
