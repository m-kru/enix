package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Delete() action.Action {
	if c.RuneIdx == c.Line.RuneCount() {
		delLine := c.Line.Join(false)
		if delLine == nil {
			return nil
		}

		return &action.NewlineDelete{Line: delLine, LineNum: c.LineNum}
	}

	c.Line.DeleteRune(c.RuneIdx)
	return &action.RuneDelete{Line: c.Line, RuneIdx: c.RuneIdx}
}

func (c *Cursor) Join() action.Action {
	delLine := c.Line.Join(true)
	if delLine == nil {
		return nil
	}

	return &action.NewlineDelete{Line: delLine, LineNum: c.LineNum}
}

func (c *Cursor) Backspace() action.Action {
	if c.RuneIdx == 0 {
		if c.Line.Prev == nil {
			// Do nothing
			return nil
		} else {
			c.Line = c.Line.Prev
			prevLineLen := c.Line.RuneCount()
			delLine := c.Line.Join(false)
			// delLine is for sure not nil here so do not check for nil.

			c.LineNum--
			c.RuneIdx += prevLineLen
			c.ColIdx = c.Line.ColumnIdx(c.RuneIdx)

			return &action.NewlineDelete{Line: delLine, LineNum: c.LineNum}
		}
	}

	c.Line.DeleteRune(c.RuneIdx - 1)
	c.RuneIdx--

	return &action.RuneDelete{Line: c.Line, RuneIdx: c.RuneIdx - 1}
}
