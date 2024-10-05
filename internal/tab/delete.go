package tab

// Delete deletes text under cursors or selections.
func (tab *Tab) Delete() {
	if tab.Cursors != nil {
		tab.deleteCursors()
	} else {
		tab.deleteSelections()
	}
}

func (tab *Tab) deleteCursors() {
	c0 := tab.Cursors // First cursor

	c := c0
	for {
		delLine := c.Delete()

		if delLine == nil {
			c2 := c0
			for {
				if c2 == nil {
					break
				}
				if c2 != c {
					c2.InformRuneDelete(c.Line, c.BufIdx)
				}
				c2 = c2.Next
			}

			for _, m := range tab.Marks {
				m.InformRuneDelete(c.Line, c.BufIdx)
			}
		} else {

		}

		c0.Prune()

		c = c.Next
		if c == nil {
			break
		}
	}

	tab.HasChanges = true
}

func (tab *Tab) deleteSelections() {
	panic("unimplemented")
}
