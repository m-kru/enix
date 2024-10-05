package cursor

import (
	"github.com/m-kru/enix/internal/line"
)

// InformRuneInsert informs the cursor about content insertion into the line.
func (c *Cursor) InformRuneInsert(l *line.Line, idx int) {
	if l != c.Line || idx > c.BufIdx {
		return
	}
	c.BufIdx++
}

func (c *Cursor) InsertRune(r rune) {
	c.Line.InsertRune(r, c.BufIdx)
	c.BufIdx++
}

func (c *Cursor) InsertNewline() *line.Line {
	newLine := c.Line.InsertNewline(c.BufIdx)

	c.Line = newLine
	c.BufIdx = 0
	c.Idx = 0

	return newLine
}
