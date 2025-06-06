package cursor

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/line"

	"github.com/mattn/go-runewidth"
)

func (c *Cursor) InsertNewline(indent bool) action.Actions {
	rIdx := c.RuneIdx
	lineNum := c.LineNum
	line := c.Line
	indentStr := line.Indent()
	newLine := c.Line.InsertNewline(c.RuneIdx)

	c.Line = newLine.Next
	c.LineNum++
	c.RuneIdx = 0
	c.colIdx = 1

	actions := make(action.Actions, 1, 2)
	actions[0] = &action.NewlineInsert{
		Line:    line,
		LineNum: lineNum,
		RuneIdx: rIdx,
		NewLine: newLine,
	}

	if indentStr == "" || !indent {
		return actions
	}

	si := c.InsertString(indentStr)
	actions = append(actions, si)

	return actions
}

func (c *Cursor) InsertRune(r rune) *action.RuneInsert {
	c.Line.InsertRune(r, c.RuneIdx)
	c.RuneIdx++
	if r == '\t' {
		c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
	} else {
		c.colIdx += runewidth.RuneWidth(r)
	}
	return &action.RuneInsert{Line: c.Line, Rune: r, RuneIdx: c.RuneIdx - 1}
}

func (c *Cursor) InsertString(s string) *action.StringInsert {
	c.Line.InsertString(s, c.RuneIdx)
	rIdx := c.RuneIdx
	rc := utf8.RuneCountInString(s)
	c.RuneIdx += rc
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
	return &action.StringInsert{
		Line:         c.Line,
		Str:          s,
		StartRuneIdx: rIdx,
		RuneCount:    rc,
	}
}

func (c *Cursor) InsertLineAbove(s string) *action.LineInsert {
	newLine, _ := line.FromString(s)

	if c.Line.Prev != nil {
		c.Line.Prev.Next = newLine
	}
	newLine.Next = c.Line
	newLine.Prev = c.Line.Prev
	c.Line.Prev = newLine

	c.LineNum++

	return &action.LineInsert{Line: newLine, LineNum: c.LineNum - 1}
}

func (c *Cursor) InsertLineBelow(s string) *action.LineInsert {
	newLine, _ := line.FromString(s)

	if c.Line.Next != nil {
		c.Line.Next.Prev = newLine
	}
	newLine.Next = c.Line.Next
	c.Line.Next = newLine
	newLine.Prev = c.Line

	return &action.LineInsert{Line: newLine, LineNum: c.LineNum + 1}
}
