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
	for rIdx < c.RuneIdx {
		r, rLen := utf8.DecodeRune(c.Line.Buf[bIdx:])
		if !unicode.IsSpace(r) {
			return false
		}

		bIdx += rLen
		rIdx++
	}
	return true
}

func (c *Cursor) WordPosition() WordPosition {
	lineRC := c.Line.RuneCount()
	r := c.Line.Rune(c.RuneIdx)

	if c.RuneIdx == 0 && lineRC > 0 && !unicode.IsSpace(r) {
		return AtWordStart
	}

	if c.RuneIdx == lineRC || unicode.IsSpace(r) {
		return InSpace
	}

	prevR := c.Line.Rune(c.RuneIdx - 1)
	if unicode.IsSpace(prevR) {
		return AtWordStart
	}

	rWordRune := util.IsWordRune(r)
	prevRWordRune := util.IsWordRune(prevR)

	if prevRWordRune && !rWordRune || !prevRWordRune && rWordRune {
		return AtWordStart
	}

	return InWord
}

// TabEnd moves cursor to the tab end.
func (c *Cursor) TabEnd() {
	for c.Line.Next != nil {
		c.Line = c.Line.Next
		c.LineNum++
	}

	c.RuneIdx = c.Line.RuneCount()
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}
