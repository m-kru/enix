package cursor

func (c *Cursor) Eq(c2 *Cursor) bool {
	if c.Line == c2.Line && c.Idx == c2.Idx && c.BufIdx == c2.BufIdx {
		return true
	}
	return false
}
