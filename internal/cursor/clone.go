package cursor

func (c *Cursor) Clone() *Cursor {
	newC := &Cursor{
		Line:    c.Line,
		LineNum: c.LineNum,
		RuneIdx: c.RuneIdx,
		colIdx:  c.colIdx,
	}

	return newC
}

func Clone(cursors []*Cursor) []*Cursor {
	if cursors == nil {
		return nil
	}

	cs := make([]*Cursor, 0, len(cursors))

	for _, c := range cursors {
		cs = append(cs, c.Clone())
	}

	return cs
}
