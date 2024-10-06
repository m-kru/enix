package tab

import (
	"github.com/m-kru/enix/internal/line"
)

// Delete deletes text under cursors or selections.
func (tab *Tab) Delete() {
	if tab.Cursors != nil {
		tab.deleteCursors(false)
	} else {
		tab.deleteSelections()
	}
}

func (tab *Tab) deleteCursors(backspace bool) {
	c0 := tab.Cursors // First cursor

	c := c0
	for {
		var delLine *line.Line
		ok := true
		if backspace {
			delLine, ok = c.Backspace()
		} else {
			delLine = c.Delete()
		}

		if !ok {
			goto nextCursor
		}

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

	nextCursor:
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

func (tab *Tab) Backspace() {
	if tab.Cursors != nil {
		tab.deleteCursors(true)
	} else {
		tab.deleteSelections()
	}
}
