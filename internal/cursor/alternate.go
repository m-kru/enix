package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) LineDown() action.Action {
	if c.Line.Next == nil {
		return nil
	}

	nextLine := c.Line.Next
	nextLine.Prev = c.Line.Prev
	if c.Line.Prev != nil {
		c.Line.Prev.Next = nextLine
	}
	if nextLine.Next != nil {
		nextLine.Next.Prev = c.Line
	}
	c.Line.Next = nextLine.Next
	c.Line.Prev = nextLine
	nextLine.Next = c.Line

	c.LineNum++

	return &action.LineDown{Line: c.Line}
}

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
	prevLine.Prev = c.Line

	c.LineNum--

	return &action.LineUp{Line: c.Line}
}
