package tab

import (
	"github.com/m-kru/enix/internal/cursor"
)

func (t *Tab) CursorLeft() {
	c := t.Cursors

	for {
		if c == nil {
			break
		}
		c.Left()
		c = c.Next
	}

	t.Cursors.Prune()
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

	t.Cursors.Prune()
}

func (t *Tab) CursorSpawnDown() {
	var newCursors *cursor.Cursor
	var lastNewCursor *cursor.Cursor

	c := t.Cursors
	for {
		nc := c.SpawnDown()

		if nc != nil {
			if newCursors == nil {
				newCursors = nc
				lastNewCursor = nc
			} else {
				lastNewCursor.Next = nc
				nc.Prev = lastNewCursor
				lastNewCursor = nc
			}
		}

		if c.Next == nil {
			break
		}
		c = c.Next
	}

	c.Next = newCursors
	newCursors.Prev = c

	t.Cursors.Prune()
}
