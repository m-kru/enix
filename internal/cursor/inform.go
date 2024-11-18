package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Inform(act action.Action) {
	switch a := act.(type) {
	case *action.LineDown:
		c.informLineDown(a)
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

func (c *Cursor) informLineDown(ld *action.LineDown) {
	if c.Line == ld.Line {
		c.LineNum++
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
	if c.LineNum < nd.Line1Num {
		return
	}

	if c.LineNum == nd.Line1Num {
		c.Line = nd.NewLine
		return
	}

	c.LineNum--
	if c.LineNum > nd.Line1Num {
		return
	}

	c.Line = nd.NewLine
	c.RuneIdx = c.RuneIdx + nd.RuneIdx
	c.ColIdx = c.Line.ColumnIdx(c.RuneIdx)
}

func (c *Cursor) informNewlineInsert(ni *action.NewlineInsert) {
	if c.LineNum < ni.LineNum {
		return
	}

	if c.LineNum == ni.LineNum && c.RuneIdx < ni.RuneIdx {
		c.Line = ni.NewLine1
		return
	}

	c.LineNum++
	if c.LineNum > ni.LineNum+1 {
		return
	}

	c.Line = ni.NewLine2
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
