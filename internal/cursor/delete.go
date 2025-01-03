package cursor

import (
	"github.com/m-kru/enix/internal/action"

	"github.com/mattn/go-runewidth"
)

func (c *Cursor) Delete() action.Action {
	l1 := c.Line
	l2 := l1.Next
	rc := c.Line.RuneCount()
	if c.RuneIdx == rc {
		newLine, _ := c.Line.Join(false)
		if newLine == nil {
			return nil
		}

		c.Line = newLine

		return &action.NewlineDelete{
			Line1:        l1,
			Line1Num:     c.LineNum,
			RuneIdx:      rc,
			Line2:        l2,
			TrimmedCount: 0,
			NewLine:      newLine,
		}
	}

	r := c.Line.DeleteRune(c.RuneIdx)
	return &action.RuneDelete{Line: c.Line, Rune: r, RuneIdx: c.RuneIdx}
}

func (c *Cursor) Backspace() action.Action {
	if c.RuneIdx == 0 {
		if c.Line.Prev == nil {
			// Do nothing
			return nil
		} else {
			l1 := c.Line.Prev
			l2 := c.Line
			l1Len := l1.RuneCount()
			newLine, _ := l1.Join(false)
			// cursor is in the next line here, so Join must success

			c.Line = newLine
			c.LineNum--
			c.RuneIdx = l1Len
			c.colIdx = c.Line.ColumnIdx(c.RuneIdx)

			return &action.NewlineDelete{
				Line1:        l1,
				Line1Num:     c.LineNum,
				RuneIdx:      l1Len,
				Line2:        l2,
				TrimmedCount: 0,
				NewLine:      newLine,
			}
		}
	}

	r := c.Line.DeleteRune(c.RuneIdx - 1)
	c.RuneIdx--
	c.colIdx -= runewidth.RuneWidth(r)

	return &action.RuneDelete{Line: c.Line, Rune: r, RuneIdx: c.RuneIdx}
}

func (c *Cursor) DeleteLine() action.Action {
	delLine := c.Line
	newLine := c.Line.Delete()
	if newLine != nil {
		c.Line = newLine
		// Check if last line was deleted
		if newLine == delLine.Prev {
			c.LineNum--
		}
		c.RuneIdx = 0
		c.colIdx = c.Line.ColumnIdx(0)
		return &action.LineDelete{
			Line:    delLine,
			LineNum: c.LineNum,
			NewLine: c.Line,
		}
	}

	act := &action.StringDelete{
		Line:         c.Line,
		Str:          c.Line.String(),
		StartRuneIdx: 0,
		RuneCount:    c.Line.RuneCount(),
	}

	c.Line.Clear()
	c.RuneIdx = 0
	c.colIdx = c.Line.ColumnIdx(0)

	return act
}
