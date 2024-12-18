package cursor

// Prune function removes duplicates from cursor list.
// A duplicate is a cursor pointing to the same line with equal buffer index.
// It also removes dead cursors pointing to the nil Line.
// Cursors may become dead, for example, when line is removed.
func Prune(cursors []*Cursor) []*Cursor {
	cs := make([]*Cursor, 0, len(cursors))

	for _, c := range cursors {
		// Check if this is dead cursor
		if c.Line == nil {
			continue
		}

		duplicate := false
		for _, c2 := range cs {
			if Equal(c, c2) {
				duplicate = true
				break
			}
		}

		if duplicate {
			continue
		}

		cs = append(cs, c)
	}

	return cs
}
