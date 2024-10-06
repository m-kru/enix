package tab

func (tab *Tab) Join() {
	if tab.Cursors != nil {
		tab.joinCursors()
	} else {
		tab.joinSelections()
	}
}

func (tab *Tab) joinCursors() {
	c0 := tab.Cursors // First cursor

	// Join
	c := c0
	for {
		delLine := c.Join()

		if delLine != nil {
			c2 := c0
			for {
				if c2 == nil {
					break
				}
				if c2 != c {
					c2.InformNewlineDelete(delLine, c.Line)
				}
				c2 = c2.Next
			}
		}

		tab.Cursors = c0.Prune()

		c = c.Next
		if c == nil {
			break
		}
	}

	tab.HasChanges = true
}

func (tab *Tab) joinSelections() {
	panic("unimplemented")
}
