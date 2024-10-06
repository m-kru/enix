package cursor

import (
	"github.com/m-kru/enix/internal/line"
)

func (c *Cursor) Delete() *line.Line {
	return c.Line.DeleteRune(c.BufIdx)
}

// InformRuneDelete informs the cursor about single rune deletion from the line.
func (c *Cursor) InformRuneDelete(l *line.Line, idx int) {
	if l != c.Line || idx >= c.BufIdx {
		return
	}
	c.BufIdx--
	c.Idx--
}

func (c *Cursor) Backspace() (*line.Line, bool) {
	if c.BufIdx == 0 {
		if c.Line.Prev == nil {
			// Do nothing
			return nil, false
		} else {
			panic("unimplemented")
		}
	}

	// Inform other cursors about deletion.
	cNext := c.Next
	for {
		if cNext == nil {
			break
		}
		cNext.InformRuneDelete(c.Line, c.BufIdx-1)
		cNext = cNext.Next
	}

	c.Line.DeleteRune(c.BufIdx - 1)
	c.BufIdx--

	return nil, true
}

// InformNewlineDelete informs cursor about newline deletion.
// If cursor was pointing the the joined line, then the
// cursor Line pointer is set to point to the previous line.
func (c *Cursor) InformNewlineDelete(l *line.Line, prevL *line.Line) {
	if c.Line != l {
		return
	}

	c.Line = prevL
	c.BufIdx = c.BufIdx + prevL.Len() - l.Len()
}
