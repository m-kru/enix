package mark

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
)

type Mark interface {
	InformRuneInsert(line *line.Line, idx int)
	InformNewlineInsert(line *line.Line, idx int)
	InformRuneDelete(line *line.Line, idx int)
}

type CursorMark struct {
	Cursors *cursor.Cursor
}

func (cm *CursorMark) InformRuneInsert(line *line.Line, idx int) {
	c := cm.Cursors
	for {
		if c == nil {
			break
		}

		c.InformRuneInsert(line, idx)

		c = c.Next
	}
}

func (cm *CursorMark) InformNewlineInsert(line *line.Line, idx int) {

}

func NewCursorMark(c *cursor.Cursor) *CursorMark {
	return &CursorMark{Cursors: c.Clone()}
}

func (cm *CursorMark) InformRuneDelete(line *line.Line, idx int) {
	c := cm.Cursors
	for {
		if c == nil {
			break
		}

		c.InformRuneDelete(line, idx)

		c = c.Next
	}
	cm.Cursors.Prune()
}
