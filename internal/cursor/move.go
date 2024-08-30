package cursor

import (
	"github.com/m-kru/enix/internal/util"
	"unicode"
)

func (c *Cursor) Down() {
	if c.Line.Next == nil {
		return
	}

	bufIdx := c.BufIdx
	nextLen := c.Line.Next.Len()
	if c.Idx != bufIdx {
		if c.Idx <= nextLen {
			bufIdx = c.Idx
		} else {
			bufIdx = nextLen
		}
	} else if bufIdx > nextLen {
		bufIdx = nextLen
	}

	c.BufIdx = bufIdx
	c.Line = c.Line.Next
}

func (c *Cursor) Left() {
	if c.BufIdx == 0 {
		if c.Line.Prev != nil {
			c.Line = c.Line.Prev
			c.BufIdx = c.Line.Len()
		}
	} else {
		c.BufIdx--
	}
	c.Idx = c.BufIdx
}

func (c *Cursor) Right() {
	if c.BufIdx == c.Line.Len() {
		if c.Line.Next != nil {
			c.Line = c.Line.Next
			c.BufIdx = 0
		}
	} else {
		c.BufIdx++
	}
	c.Idx = c.BufIdx
}

func (c *Cursor) Up() {
	if c.Line.Prev == nil {
		return
	}

	bufIdx := c.BufIdx
	prevLen := c.Line.Prev.Len()
	if c.Idx != bufIdx {
		if c.Idx <= prevLen {
			bufIdx = c.Idx
		} else {
			bufIdx = prevLen
		}
	} else if bufIdx > prevLen {
		bufIdx = prevLen
	}

	c.BufIdx = bufIdx
	c.Line = c.Line.Prev
}

func (c *Cursor) PrevWordStart() {
	if idx, ok := util.PrevWordStart(c.Line.Buf, c.BufIdx); ok {
		c.BufIdx = idx
		return
	}

	line := c.Line.Prev
	for {
		if line == nil {
			return
		}

		if idx, ok := util.PrevWordStart(line.Buf, line.Len()); ok {
			c.Line = line
			c.BufIdx = idx + 1
			c.Idx = c.BufIdx
			return
		}

		line = line.Prev
	}
}

func (c *Cursor) WordEnd() {
	if idx, ok := util.WordEnd(c.Line.Buf, c.BufIdx); ok {
		c.BufIdx = idx + 1 // + 1 as we have found word end index.
		c.Idx = c.BufIdx
		return
	}

	line := c.Line.Next
	for {
		if line == nil {
			return
		}

		if idx, ok := util.WordEnd(line.Buf, 0); ok {
			c.Line = line
			c.BufIdx = idx + 1
			c.Idx = c.BufIdx
			return
		}

		line = line.Next
	}
}

func (c *Cursor) WordStart() {
	if idx, ok := util.WordStart(c.Line.Buf, c.BufIdx); ok {
		c.BufIdx = idx
		c.Idx = c.BufIdx
		return
	}

	line := c.Line.Next
	for {
		if line == nil {
			return
		}

		if idx, ok := util.WordStart(line.Buf, 0); ok {
			c.Line = line
			c.BufIdx = idx
			c.Idx = c.BufIdx
			return
		}

		line = line.Next
	}
}

func (c *Cursor) LineStart() {
	for i, r := range c.Line.Buf {
		if !unicode.IsSpace(r) {
			if c.BufIdx == i {
				c.BufIdx = 0
			} else {
				c.BufIdx = i
			}
			c.Idx = c.BufIdx
			return
		}
	}
	c.BufIdx = 0
	c.Idx = c.BufIdx
}

func (c *Cursor) LineEnd() {
	for i := c.Line.Len() - 1; i > 0; i-- {
		r := c.Line.Buf[i]
		if !unicode.IsSpace(r) {
			if c.BufIdx == i {
				c.BufIdx = c.Line.Len() - 1
			} else {
				c.BufIdx = i
			}
			c.Idx = c.BufIdx
			return
		}
	}
	c.BufIdx = c.Line.Len()
	c.Idx = c.BufIdx
}
