package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

func (c *Cursor) Inform(act action.Action) {
	switch a := act.(type) {
	case action.Actions:
		for _, subA := range a {
			c.Inform(subA)
		}
	case *action.LineDelete:
		c.informLineDelete(a)
	case *action.LineInsert:
		c.informLineInsert(a)
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
	case *action.StringDelete:
		c.informStringDelete(a)
	case *action.StringInsert:
		c.informStringInsert(a)
	}
}

func (c *Cursor) informLineDelete(ld *action.LineDelete) {
	if ld.LineNum < c.LineNum {
		c.LineNum--
	}
}

func (c *Cursor) informLineInsert(li *action.LineInsert) {
	if li.LineNum < c.LineNum {
		c.LineNum++
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
	c.RuneIdx = c.RuneIdx + nd.RuneIdx - nd.TrimmedCount
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
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
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}

// informRuneDelete informs the cursor about single rune deletion from the line.
func (c *Cursor) informRuneDelete(rd *action.RuneDelete) {
	if rd.Line != c.Line || rd.RuneIdx >= c.RuneIdx {
		return
	}
	c.RuneIdx--
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}

// informRuneInsert informs the cursor about content insertion into the line.
func (c *Cursor) informRuneInsert(ri *action.RuneInsert) {
	if ri.Line != c.Line || ri.RuneIdx > c.RuneIdx {
		return
	}
	c.RuneIdx++
}

func (c *Cursor) informStringDelete(sd *action.StringDelete) {
	if c.Line != sd.Line {
		return
	}

	if c.RuneIdx < sd.StartRuneIdx {
		return
	}

	c.RuneIdx -= sd.RuneCount
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}

func (c *Cursor) informStringInsert(si *action.StringInsert) {
	if c.Line != si.Line {
		return
	}

	if c.RuneIdx < si.StartRuneIdx {
		return
	}

	c.RuneIdx += si.RuneCount
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}
