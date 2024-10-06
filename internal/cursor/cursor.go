package cursor

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"

	"github.com/mattn/go-runewidth"
)

// Cursors must be stored in order. Thanks to this, only next cursors must be
// informed about line changes.
type Cursor struct {
	Config *cfg.Config
	Line   *line.Line
	Idx    int
	BufIdx int // Index into line buffer.

	Prev *Cursor
	Next *Cursor
}

// Count returns the number of cursors in the list starting from the conter c.
func (c *Cursor) Count() int {
	cnt := 1
	for {
		if c.Next == nil {
			break
		}
		c = c.Next
		cnt++
	}
	return cnt
}

// Get returns nth cursor.
// If there is less than n cursots, it returns nil.
func (c *Cursor) Get(n int) *Cursor {
	i := n

	for {
		if i == 1 {
			return c
		}

		if c.Next == nil {
			break
		}

		c = c.Next
		i--
	}

	return nil
}

// Col returns column number of the cursor within the string in the buffer.
func (c *Cursor) Column() int {
	return c.Line.ColumnIdx(c.BufIdx, c.Config.TabWidth)
}

// Width returns width of the rune under the cursor.
func (c *Cursor) Width() int {
	if c.BufIdx == c.Line.Len() {
		rw := runewidth.RuneWidth(c.Config.NewlineRune)
		if rw == 0 {
			return 1
		}
		return rw
	}

	r := c.Line.Rune(c.BufIdx)
	if r == '\t' {
		return 1
	}
	return runewidth.RuneWidth(r)
}

// Word returns word under cursor.
func (c *Cursor) Word() string {
	return ""
}

func (c *Cursor) View() view.View {
	return view.View{
		Line:   c.Line.Num(),
		Column: c.Column(),
		Width:  c.Width(),
		Height: 1,
	}
}

// Last returns last cursor in the cursor list.
func (c *Cursor) Last() *Cursor {
	for {
		if c.Next == nil {
			return c
		}
		c = c.Next
	}
}

// Prune function removes duplicates from cursor list.
// A duplicate is a cursor pointing to the same line with equal buffer index.
// It also removes dead cursors  pointing to the nil Line.
// Cursors may become dead, for example, when line is removed.
func (c *Cursor) Prune() *Cursor {
	// First remove dead cursors
	c0 := c
	for {
		if c.Line != nil {
			c = c.Next
			if c == nil {
				break
			}
			continue
		}

		deadC := c

		if deadC == c0 && c.Next == nil {
			return nil
		}

		if deadC == c0 {
			c = c.Next
			c0 = c
		} else {
			deadC.Prev.Next = deadC.Next
			if deadC.Next != nil {
				deadC.Next.Prev = deadC.Prev
			}
			c = deadC.Next
		}

		deadC.Prev = nil
		deadC.Next = nil

		if c == nil {
			break
		}
	}

	// Remove cursor duplicates
	c = c0
	for {
		c2 := c.Next
		if c2 == nil {
			break
		}

		for {
			if Equal(c, c2) {
				c2.Prev.Next = c2.Next
				if c2.Next != nil {
					c2.Next.Prev = c2.Prev
				}
			}
			c2 = c2.Next
			if c2 == nil {
				break
			}
		}

		c = c.Next
		if c == nil {
			break
		}
	}

	return c0
}
