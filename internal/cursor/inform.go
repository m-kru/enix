package cursor

import (
	"github.com/m-kru/enix/internal/action"
)

// informNewlineDelete informs cursor about newline deletion.
// If cursor was pointing the the joined line, then the
// cursor Line pointer is set to point to the previous line.
func (c *Cursor) informNewlineDelete(nd *action.NewlineDelete) {
	if c.Line != nd.Line {
		return
	}

	c.Line = nd.Line.Prev
	c.RuneIdx = c.RuneIdx + c.Line.RuneCount() - nd.Line.RuneCount()
	c.Idx = c.RuneIdx
}

func (c *Cursor) informNewlineInsert(ni *action.NewlineInsert) {
	if c.Line != ni.Line.Prev || c.RuneIdx <= ni.Idx {
		return
	}

	c.Line = ni.Line
	c.RuneIdx -= ni.Idx
	c.Idx = c.RuneIdx
}

// informRuneDelete informs the cursor about single rune deletion from the line.
func (c *Cursor) informRuneDelete(rd *action.RuneDelete) {
	if rd.Line != c.Line || rd.Idx >= c.RuneIdx {
		return
	}
	c.RuneIdx--
	c.Idx--
}

// informRuneInsert informs the cursor about content insertion into the line.
func (c *Cursor) informRuneInsert(ri *action.RuneInsert) {
	if ri.Line != c.Line || ri.Idx > c.RuneIdx {
		return
	}
	c.RuneIdx++
}
