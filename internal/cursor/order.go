package cursor

func Equal(c1, c2 *Cursor) bool {
	if c1.Line == c2.Line && c1.Idx == c2.Idx && c1.BufIdx == c2.BufIdx {
		return true
	}
	return false
}
