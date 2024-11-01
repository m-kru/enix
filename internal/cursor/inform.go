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
	c.BufIdx = c.BufIdx + c.Line.Len() - nd.Line.Len()
	c.Idx = c.BufIdx
}

func (c *Cursor) informNewlineInsert(ni *action.NewlineInsert) {
	if c.Line != ni.Line.Prev || c.BufIdx <= ni.Idx {
		return
	}

	c.Line = ni.Line
	c.BufIdx -= ni.Idx
	c.Idx = c.BufIdx
}

// informRuneDelete informs the cursor about single rune deletion from the line.
func (c *Cursor) informRuneDelete(rd *action.RuneDelete) {
	if rd.Line != c.Line || rd.Idx >= c.BufIdx {
		return
	}
	c.BufIdx--
	c.Idx--
}

// informRuneInsert informs the cursor about content insertion into the line.
func (c *Cursor) informRuneInsert(ri *action.RuneInsert) {
	if ri.Line != c.Line || ri.Idx > c.BufIdx {
		return
	}
	c.BufIdx++
}
