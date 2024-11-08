package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Inform(act action.Action) {
	switch a := act.(type) {
	case *action.LineUp:
		c.informLineUp(a)
	case *action.NewlineDelete:
		c.informNewlineDelete(a)
	case *action.NewlineInsert:
		c.informNewlineInsert(a)
	case *action.RuneDelete:
		c.informRuneDelete(a)
	case *action.RuneInsert:
		c.informRuneInsert(a)
	}
}

func (c *Cursor) informLineUp(lu *action.LineUp) {
	if c.Line == lu.Line {
		c.LineNum--
	}
}

// informNewlineDelete informs cursor about newline deletion.
// If cursor was pointing the the joined line, then the
// cursor Line pointer is set to point to the previous line.
func (c *Cursor) informNewlineDelete(nd *action.NewlineDelete) {
	if nd.LineNum < c.LineNum {
		c.LineNum--
	}

	if c.Line != nd.Line {
		return
	}

	c.Line = nd.Line.Prev
	c.RuneIdx = c.RuneIdx + c.Line.RuneCount() - nd.Line.RuneCount()
	c.ColIdx = c.Line.ColumnIdx(c.RuneIdx)
}

func (c *Cursor) informNewlineInsert(ni *action.NewlineInsert) {
	if (ni.LineNum < c.LineNum) || (ni.LineNum == c.LineNum && ni.RuneIdx < c.RuneIdx) {
		c.LineNum++
	}

	if c.Line != ni.Line.Prev || c.RuneIdx <= ni.RuneIdx {
		return
	}

	c.Line = ni.Line
	c.RuneIdx -= ni.RuneIdx
	c.ColIdx = c.Line.ColumnIdx(c.RuneIdx)
}

// informRuneDelete informs the cursor about single rune deletion from the line.
func (c *Cursor) informRuneDelete(rd *action.RuneDelete) {
	if rd.Line != c.Line || rd.RuneIdx >= c.RuneIdx {
		return
	}
	c.RuneIdx--
	c.ColIdx = c.Line.ColumnIdx(c.RuneIdx)
}

// informRuneInsert informs the cursor about content insertion into the line.
func (c *Cursor) informRuneInsert(ri *action.RuneInsert) {
	if ri.Line != c.Line || ri.RuneIdx > c.RuneIdx {
		return
	}
	c.RuneIdx++
}
