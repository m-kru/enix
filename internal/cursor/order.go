package cursor

func Equal(c1, c2 *Cursor) bool {
	if c1.Line == c2.Line && c1.RuneIdx == c2.RuneIdx {
		return true
	}
	return false
}
