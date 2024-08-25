package cursor

func (c *Cursor) SpawnDown() *Cursor {
	if c.Line.Next == nil {
		return nil
	}

	nc := &Cursor{
		Config: c.Config,
		Line:   c.Line.Next,
		Idx:    c.Idx,
		BufIdx: c.BufIdx,
	}

	if nc.BufIdx > nc.Line.Len() {
		nc.BufIdx = nc.Line.Len()
	}

	return nc
}

func (c *Cursor) SpawnUp() *Cursor {
	if c.Line.Prev == nil {
		return nil
	}

	nc := &Cursor{
		Config: c.Config,
		Line:   c.Line.Prev,
		Idx:    c.Idx,
		BufIdx: c.BufIdx,
	}

	if nc.BufIdx > nc.Line.Len() {
		nc.BufIdx = nc.Line.Len()
	}

	return nc
}
