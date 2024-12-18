package cursor

import (
	"unicode"
	"unicode/utf8"

	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/mattn/go-runewidth"
)

type WordPosition int

const (
	InSpace = iota
	AtWordStart
	InWord
)

// Cursors must be stored in order. Thanks to this, only next cursors must be
// informed about line changes.
type Cursor struct {
	Line    *line.Line
	LineNum int
	RuneIdx int // Line rune index
	colIdx  int // Line column index
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

func (c *Cursor) WithinIndent() bool {
	bIdx := 0
	rIdx := 0
	for {
		if rIdx == c.RuneIdx {
			break
		}

		r, rLen := utf8.DecodeRune(c.Line.Buf[bIdx:])
		if !unicode.IsSpace(r) {
			return false
		}

		bIdx += rLen
		rIdx++
	}
	return true
}

/*
func (c *Cursor) WordPosition() WordPosition {
	if c.RuneIdx == c.Line.RuneCount() || unicode.IsSpace(c.Line.Rune(c.RuneIdx)) {
		return InSpace
	}

	if !util.IsWordRune(c.Line.Rune(c.RuneIdx-1)) {
		return AtWordStart
	}

	return InWord
}
*/
