package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Join() action.Action {
	l1 := c.Line
	l2 := l1.Next
	rc := c.Line.RuneCount()
	newLine := c.Line.Join(true)
	if newLine == nil {
		return nil
	}

	return &action.NewlineDelete{
		Line1:    l1,
		Line1Num: c.LineNum,
		RuneIdx:  rc,
		Line2:    l2,
		NewLine:  newLine,
	}
}

func (c *Cursor) LineDown() action.Action {
	ok := c.Line.Down()
	if !ok {
		return nil
	}

	c.LineNum++

	return &action.LineDown{Line: c.Line}
}

func (c *Cursor) LineUp() action.Action {
	ok := c.Line.Up()
	if !ok {
		return nil
	}

	c.LineNum--

	return &action.LineUp{Line: c.Line}
}
