package cursor

import (
	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/line"
)

// Cursors must be stored in order. Thanks to this, only next cursors must be
// informed about line changes.
type Cursor struct {
	Screen tcell.Screen

	Colors *cfg.Colorscheme

	Line   *line.Line
	BufIdx int // Index into line buffer.

	Prev *Cursor
	Next *Cursor
}

// Col returns column number of the cursor within the string in the buffer.
func (c *Cursor) Col() int { return c.BufIdx + 1 }

// Word returns word under cursor.
func (c *Cursor) Word() string {
	return ""
}

// Prune function removes duplicates from cursor list.
// A duplicate is a cursor pointing to the same line and buffer index.
func (c *Cursor) Prune() {
	panic("unimplemented")
}

func (c *Cursor) Left(n int) {
	if n < c.BufIdx {
		c.BufIdx -= n
		return
	}

	if c.Line.Prev == nil {
		c.BufIdx = 0
		return
	}
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

func (c *Cursor) HandleLeft() {
	if c.BufIdx == 0 {
		if c.Line.Prev == nil {
			// Do nothing
			return
		} else {
			panic("unimplemented")
		}
	}
	c.BufIdx--
}

func (c *Cursor) HandleRight() {
	if c.BufIdx == c.Line.Len() {
		if c.Line.Next == nil {
			// Do nothing
			return
		} else {
			panic("unimplemented")
		}
	}
	c.BufIdx++
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

func (c *Cursor) HandleBackspace() {
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
		cNext.InformDeletion(c.Line, c.BufIdx, 1)
	}

	c.Line.Delete(c.BufIdx-1, 1)
	c.BufIdx--
}

func (c *Cursor) Render(x, y, offset int) {
	x = x + c.BufIdx - offset
	r, _, _, _ := c.Screen.GetContent(x, y)
	c.Screen.SetContent(x, y, r, nil, c.Colors.Cursor)
}
