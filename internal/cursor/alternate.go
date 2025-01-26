package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Join() action.Action {
	l := c.Line
	rc := c.Line.RuneCount()
	newLine, trimmedCount := c.Line.Join(true)
	if newLine == nil {
		return nil
	}

	c.Line = newLine

	return &action.NewlineDelete{
		Line:         l,
		LineNum:      c.LineNum,
		RuneIdx:      rc,
		TrimmedCount: trimmedCount,
		NewLine:      newLine,
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
