package cursor

func Clone(cursors []*Cursor) []*Cursor {
	cs := make([]*Cursor, 0, len(cursors))

	for _, c := range cursors {
		c := &Cursor{
			Line:    c.Line,
			LineNum: c.LineNum,
			ColIdx:  c.ColIdx,
			RuneIdx: c.RuneIdx,
		}

		cs = append(cs, c)
	}

	return cs
}
