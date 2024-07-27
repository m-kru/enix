package tab

func (t *Tab) CursorLeft() {
	c := t.Cursors

	for {
		if c == nil {
			break
		}

		c.Left()
		c = c.Next
	}
}

func (t *Tab) CursorRight() {
	c := t.Cursors

	for {
		if c == nil {
			break
		}

		c.Right()
		c = c.Next
	}
}
