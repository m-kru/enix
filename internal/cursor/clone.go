package cursor

func (c *Cursor) Clone() *Cursor {
	// Should Idx equal c.Idx or c.BufIdx?
	first := &Cursor{
		Config: c.Config,
		Line:   c.Line,
		Idx:    c.Idx,
		BufIdx: c.BufIdx,
	}
	prev := first

	c = c.Next
	for {
		if c == nil {
			break
		}

		c2 := Cursor{
			Config: c.Config,
			Line:   c.Line,
			Idx:    c.Idx,
			BufIdx: c.BufIdx,
		}
		c2.Prev = prev
		prev.Next = &c2
		prev = &c2

		c = c.Next
	}

	return first
}
