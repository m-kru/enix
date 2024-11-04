package cursor

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/mattn/go-runewidth"
)

// Cursors must be stored in order. Thanks to this, only next cursors must be
// informed about line changes.
type Cursor struct {
	Line    *line.Line
	LineNum int
	Idx     int
	RuneIdx int // Line rune index
}

// Column returns column number of the cursor within the string in the buffer.
func (c *Cursor) Column() int {
	return c.Line.ColumnIdx(c.RuneIdx)
}

// Width returns width of the rune under the cursor.
func (c *Cursor) Width() int {
	if c.RuneIdx == c.Line.RuneCount() {
		return 1
	}

	r := c.Line.Rune(c.RuneIdx)
	if r == '\t' {
		return 1
	}
	return runewidth.RuneWidth(r)
}

// GetWord returns word under cursor.
func (c *Cursor) GetWord() string {
	return util.GetWord([]rune(c.Line.String()), c.RuneIdx)
}

func (c *Cursor) View() view.View {
	return view.View{
		Line:   c.LineNum,
		Column: c.Column(),
		Width:  c.Width(),
		Height: 1,
	}
}

func (c *Cursor) Inform(act action.Action) {
	switch a := act.(type) {
	case *action.NewlineDelete:
		c.informNewlineDelete(a)
	case *action.NewlineInsert:
		c.informNewlineInsert(a)
	case *action.RuneDelete:
		c.informRuneDelete(a)
	case *action.RuneInsert:
		c.informRuneInsert(a)
	}
}

// Prune function removes duplicates from cursor list.
// A duplicate is a cursor pointing to the same line with equal buffer index.
// It also removes dead cursors pointing to the nil Line.
// Cursors may become dead, for example, when line is removed.
func Prune(cursors []*Cursor) []*Cursor {
	cs := make([]*Cursor, 0, len(cursors))

	for _, c := range cursors {
		// Check if this is dead cursor
		if c.Line == nil {
			continue
		}

		duplicate := false
		for _, c2 := range cs {
			if Equal(c, c2) {
				duplicate = true
				break
			}
		}

		if duplicate {
			continue
		}

		cs = append(cs, c)
	}

	return cs
}
