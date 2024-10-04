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
// A duplicate is a cursor pointing to the same line with the same index and buffer index.
func (c *Cursor) Prune() {
	for {
		c2 := c.Next
		if c2 == nil {
			return
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
			return
		}
	}
}

// InformRuneInsert informs the cursor about content insertion into the line.
func (c *Cursor) InformRuneInsert(l *line.Line, idx int) {
	if l != c.Line || idx > c.BufIdx {
		return
	}
	c.BufIdx++
}

// InformDeletion informs the cursor about content deletion from the line.
func (c *Cursor) InformDeletion(l *line.Line, idx int, size int) {
	if l != c.Line || idx > c.BufIdx {
		return
	}
	panic("unimplemented")
}

func (c *Cursor) InsertRune(r rune) {
	c.Line.InsertRune(r, c.BufIdx)
	c.BufIdx++
}

func (c *Cursor) InsertNewline() {
	newLine := c.Line.InsertNewline(c.BufIdx)

	// Update line pointer for all cursors in the same line as c, but after c.
	c2 := c.Last()
	for {
		if c2 == nil {
			break
		}

		if c2.Line == c.Line && c2.BufIdx > c.BufIdx {
			c2.Line = newLine
			c2.BufIdx -= c.BufIdx
			c2.Idx = c2.BufIdx
		}

		c2 = c2.Prev
	}

	c.Line = newLine
	c.BufIdx = 0
	c.Idx = 0
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
