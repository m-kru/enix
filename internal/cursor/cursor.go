package cursor

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"
)

// Cursors must be stored in order. Thanks to this, only next cursors must be
// informed about line changes.
type Cursor struct {
	Line   *line.Line
	Idx    int
	BufIdx int // Index into line buffer.

	Prev *Cursor
	Next *Cursor
}

// Col returns column number of the cursor within the string in the buffer.
func (c *Cursor) Column() int { return c.BufIdx + 1 }

func (c *Cursor) LineNum() int { return c.Line.Num() }

// Word returns word under cursor.
func (c *Cursor) Word() string {
	return ""
}

// FarRightBufIdx returns the buf index of the far right cursor.
func (c *Cursor) FarRightBufIdx() int {
	idx := c.BufIdx
	for {
		c = c.Next
		if c == nil {
			return idx
		}
		if c.BufIdx > idx {
			idx = c.BufIdx
		}
	}
}

// Prune function removes duplicates from cursor list.
// A duplicate is a cursor pointing to the same line and buffer index.
func (c *Cursor) Prune() {
	panic("unimplemented")
}

// InformInsertion informs the cursor about content insertion into the line.
func (c *Cursor) InformInsertion(l *line.Line, idx int, size int) {
	if l != c.Line || idx > c.BufIdx {
		return
	}
	panic("unimplemented")
}

// InformDeletion informs the cursor about content deletion from the line.
func (c *Cursor) InformDeletion(l *line.Line, idx int, size int) {
	if l != c.Line || idx > c.BufIdx {
		return
	}
	panic("unimplemented")
}

func (c *Cursor) HandleRune(r rune) {
	// Inform other cursors about deletion.
	cNext := c.Next
	for {
		if cNext == nil {
			break
		}
		cNext.InformDeletion(c.Line, c.BufIdx, 1)
	}

	c.Line.InsertRune(r, c.BufIdx)
	c.BufIdx++
}

func (c *Cursor) Backspace() {
	if c.BufIdx == 0 {
		if c.Line.Prev == nil {
			// Do nothing
			return
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
		cNext.InformDeletion(c.Line, c.BufIdx-1, 1)
		cNext = cNext.Next
	}

	c.Line.Delete(c.BufIdx-1, 1)
	c.BufIdx--
}

func (c *Cursor) Delete() {
	if c.BufIdx == c.Line.Len() {
		if c.Line.Next == nil {
			// Do nothing
			return
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
		cNext.InformDeletion(c.Line, c.BufIdx, 1)
		cNext = cNext.Next
	}

	c.Line.Delete(c.BufIdx, 1)
}

func (c *Cursor) Render(colors *cfg.Colorscheme, frame frame.Frame, view view.View) {
	x := c.BufIdx - view.Column + 1
	r := frame.GetContent(x, 0)
	frame.SetContent(x, 0, r, colors.Cursor)
}
