package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Join() action.Action {
	rc := c.Line.RuneCount()
	ok := c.Line.Join(true)
	if !ok {
		return nil
	}

	return &action.NewlineDelete{Line: c.Line, LineNum: c.LineNum, RuneIdx: rc}
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
