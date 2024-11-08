package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) LineUp() action.Action {
	if c.Line.Prev == nil {
		return nil
	}

	prevLine := c.Line.Prev
	prevLine.Next = c.Line.Next
	if c.Line.Next != nil {
		c.Line.Next.Prev = prevLine
	}
	if prevLine.Prev != nil {
		prevLine.Prev.Next = c.Line
	}
	c.Line.Prev = prevLine.Prev
	c.Line.Next = prevLine

	c.LineNum--

	return &action.LineUp{Line: c.Line}
}
