package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Delete() action.Action {
	rc := c.Line.RuneCount()
	if c.RuneIdx == rc {
		ok := c.Line.Join(false)
		if !ok {
			return nil
		}

		return &action.NewlineDelete{Line: c.Line, LineNum: c.LineNum, RuneIdx: rc}
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
			c.Line = c.Line.Prev
			prevLineLen := c.Line.RuneCount()
			_ = c.Line.Join(false)
			// cursor is in the next line here, so Join must success

			c.LineNum--
			c.RuneIdx += prevLineLen
			c.ColIdx = c.Line.ColumnIdx(c.RuneIdx)

			return &action.NewlineDelete{Line: c.Line, LineNum: c.LineNum, RuneIdx: prevLineLen}
		}
	}

	r := c.Line.DeleteRune(c.RuneIdx - 1)
	c.RuneIdx--

	return &action.RuneDelete{Line: c.Line, Rune: r, RuneIdx: c.RuneIdx - 1}
}
