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

		return &action.NewlineDelete{Line: delLine}
	}

	c.Line.DeleteRune(c.RuneIdx)
	return &action.RuneDelete{Line: c.Line, RuneIdx: c.RuneIdx}
}

func (c *Cursor) Join() action.Action {
	delLine := c.Line.Join(false)
	if delLine == nil {
		return nil
	}

	return &action.NewlineDelete{Line: delLine}
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

			c.RuneIdx += prevLineLen
			c.Idx = c.RuneIdx

			return &action.NewlineDelete{Line: delLine}
		}
	}

	c.Line.DeleteRune(c.RuneIdx - 1)
	c.RuneIdx--

	return &action.RuneDelete{Line: c.Line, RuneIdx: c.RuneIdx - 1}
}
