package cursor

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
