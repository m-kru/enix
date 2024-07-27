package tab

func (t *Tab) CursorDown() {
	c := t.Cursors

	for {
		if c == nil {
			break
		}
		c.Down()
		c = c.Next
	}
}

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

func (t *Tab) CursorUp() {
	c := t.Cursors

	for {
		if c == nil {
			break
		}
		c.Up()
		c = c.Next
	}
}
