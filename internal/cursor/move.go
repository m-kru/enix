package cursor

import "github.com/m-kru/enix/internal/util"

func (c *Cursor) Left() {
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

func (c *Cursor) Right() {
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

func (c *Cursor) WordStart() {
	if idx, ok := util.WordStart(c.Line.Buf, c.BufIdx); ok {
		c.BufIdx = idx
		return
	}

	if c.Line.Prev == nil {
		// Do nothing
		return
	}

	panic("unimplemented")
}

func (c *Cursor) WordEnd() {
	if idx, ok := util.WordEnd(c.Line.Buf, c.BufIdx); ok {
		c.BufIdx = idx + 1 // + 1 as we have found word end index.
		return
	}

	if c.Line.Next == nil {
		// Do nothing
		return
	}

	panic("unimplemented")
}
