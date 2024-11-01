package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Delete() action.Action {
	if c.BufIdx == c.Line.Len() {
		prevLine := c.Line.Prev

		delLine := c.Line.Join(false)
		if delLine == nil {
			return nil
		}

		return &action.NewlineDelete{Line: delLine, PrevLine: prevLine}
	}

	c.Line.DeleteRune(c.BufIdx)
	return &action.RuneDelete{Line: c.Line, Idx: c.BufIdx}
}

func (c *Cursor) Join() action.Action {
	prevLine := c.Line.Prev

	delLine := c.Line.Join(false)
	if delLine == nil {
		return nil
	}

	return &action.NewlineDelete{Line: delLine, PrevLine: prevLine}
}

func (c *Cursor) Backspace() action.Action {
	if c.BufIdx == 0 {
		if c.Line.Prev == nil {
			// Do nothing
			return nil
		} else {
			panic("unimplemented")
		}
	}

	c.Line.DeleteRune(c.BufIdx - 1)
	c.BufIdx--

	return &action.RuneDelete{Line: c.Line, Idx: c.BufIdx - 1}
}
