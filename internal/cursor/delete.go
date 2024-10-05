package cursor

import (
	"github.com/m-kru/enix/internal/line"
)

func (c *Cursor) Delete() *line.Line {
	delLine := c.Line.DeleteRune(c.BufIdx)

	if delLine == nil {

	} else {
		panic("unimplemented")
	}

	return nil
}

// InformDeletion informs the cursor about content deletion from the line.
func (c *Cursor) InformRuneDelete(l *line.Line, idx int) {
	if l != c.Line || idx >= c.BufIdx {
		return
	}
	c.BufIdx--
	c.Idx--
}
