package cursor

func (c *Cursor) SpawnDown() *Cursor {
	if c.Line.Next == nil {
		return nil
	}

	nc := &Cursor{
		Config:  c.Config,
		Line:    c.Line.Next,
		Idx:     c.Idx,
		RuneIdx: c.RuneIdx,
	}

	if nc.RuneIdx > nc.Line.RuneCount() {
		nc.RuneIdx = nc.Line.RuneCount()
	}

	return nc
}

func (c *Cursor) SpawnUp() *Cursor {
	if c.Line.Prev == nil {
		return nil
	}

	nc := &Cursor{
		Config:  c.Config,
		Line:    c.Line.Prev,
		Idx:     c.Idx,
		RuneIdx: c.RuneIdx,
	}

	if nc.RuneIdx > nc.Line.RuneCount() {
		nc.RuneIdx = nc.Line.RuneCount()
	}

	return nc
}
