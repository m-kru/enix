package cursor

func Clone(cursors []*Cursor) []*Cursor {
	cs := make([]*Cursor, 0, len(cursors))

	for _, c := range cursors {
		// Should Idx equal c.Idx or c.BufIdx?
		c := &Cursor{
			Config: c.Config,
			Line:   c.Line,
			Idx:    c.Idx,
			BufIdx: c.BufIdx,
		}

		cs = append(cs, c)
	}

	return cs
}
