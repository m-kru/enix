package tab

import (
	"github.com/m-kru/enix/internal/cursor"
)

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
