package cursor

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

func (c *Cursor) WordStart() {
	if idx, ok := c.Line.WordStart(c.BufIdx); ok {
		c.BufIdx = idx
		return
	}

	line := c.Line.Prev
	for {
		if line == nil {
			return
		}

		if idx, ok := line.WordStart(line.Len()); ok {
			c.Line = line
			c.BufIdx = idx + 1
			c.Idx = c.BufIdx
			return
		}

		line = line.Prev
	}
}

func (c *Cursor) WordEnd() {
	if idx, ok := c.Line.WordEnd(c.BufIdx); ok {
		c.BufIdx = idx + 1 // + 1 as we have found word end index.
		c.Idx = c.BufIdx
		return
	}

	line := c.Line.Next
	for {
		if line == nil {
			return
		}

		if idx, ok := line.WordEnd(0); ok {
			c.Line = line
			c.BufIdx = idx + 1
			c.Idx = c.BufIdx
			return
		}

		line = line.Next
	}
}

func (c *Cursor) LineStart() {
	c.BufIdx = 0
}
