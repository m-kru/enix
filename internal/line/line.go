package line

import (
	"unicode"
	"unicode/utf8"

	"github.com/mattn/go-runewidth"
)

type Line struct {
	Buf []byte

	Prev *Line
	Next *Line
}

func (l *Line) RuneCount() int { return utf8.RuneCount(l.Buf) }

func (l *Line) String() string { return string(l.Buf) }

func (l *Line) BufIdx(runeIdx int) int {
	rIdx := 0
	bIdx := 0
	for {
		_, rLen := utf8.DecodeRune(l.Buf[bIdx:])
		if rIdx == runeIdx {
			return bIdx
		}
		bIdx += rLen
		rIdx++
	}
}

func (l *Line) Rune(runeIdx int) rune {
	rIdx := 0
	bIdx := 0
	for {
		r, rLen := utf8.DecodeRune(l.Buf[bIdx:])
		if rIdx == runeIdx {
			return r
		}
		bIdx += rLen
		rIdx++
	}
}

// Num returns line number in the line list.
func (l *Line) Num() int {
	n := 1
	for {
		if l.Prev == nil {
			return n
		}
		l = l.Prev
		n++
	}
}

// Columns returns number of columns required by the line.
// It doesn't include the end of line character, as it is
// not stored in the line buffer.
func (l *Line) Columns() int {
	cols := 0
	bIdx := 0

	for bIdx < len(l.Buf) {
		r, rLen := utf8.DecodeRune(l.Buf[bIdx:])

		if r == '\t' {
			cols += 8
		} else {
			cols += runewidth.RuneWidth(r)
		}

		bIdx += rLen
	}

	return cols
}

// Get returns nth line.
// If there is less than n lines, it returns nil.
func (l *Line) Get(n int) *Line {
	i := n

	for {
		if i == 1 {
			return l
		}

		if l.Next == nil {
			break
		}

		l = l.Next
		i--
	}

	return nil
}

func (l *Line) GetNextNonEmpty() *Line {
	l = l.Next

	for l != nil {
		if len(l.Buf) > 0 {
			return l
		}

		l = l.Next
	}

	return nil
}

func (l *Line) GetPrevNonEmpty() *Line {
	l = l.Prev

	for l != nil {
		if len(l.Buf) > 0 {
			return l
		}

		l = l.Prev
	}

	return nil
}

// Last returns last line in the list.
func (l *Line) Last() *Line {
	for {
		if l.Next == nil {
			return l
		}
		l = l.Next
	}
}

// ColumnIdx returns first column index for provided rune index.
func (l *Line) ColumnIdx(runeIdx int) int {
	if len(l.Buf) == 0 {
		return 1
	}

	cIdx := 0
	bIdx := 0
	rIdx := 0

	for {
		if rIdx == runeIdx {
			cIdx += 1
			break
		}

		r, rLen := utf8.DecodeRune(l.Buf[bIdx:])

		if r == '\t' {
			cIdx += 8 - (cIdx % 8)
		} else {
			cIdx += runewidth.RuneWidth(r)
		}

		bIdx += rLen
		rIdx++
	}

	return cIdx
}

// RuneIdx returns rune index for provided column index.
// The second return is a rune subcolumn index.
// The third return is false if column c does not exists in line.
func (l *Line) RuneIdx(colIdx int) (int, int, bool) {
	bIdx := 0
	cIdx := 1
	rIdx := 0

	for bIdx < len(l.Buf) {
		r, rLen := utf8.DecodeRune(l.Buf[bIdx:])
		rWidth := runewidth.RuneWidth(r)
		if r == '\t' {
			if cIdx == colIdx {
				return rIdx, 0, true
			}

			width := 8 - ((cIdx - 1) % 8)
			if cIdx+width > colIdx {
				return rIdx, colIdx - cIdx, true
			}
			cIdx += width
		} else if rWidth == 1 {
			if cIdx == colIdx {
				return rIdx, 0, true
			}
			cIdx++
		} else {
			if cIdx == colIdx {
				return rIdx, 0, true
			} else if cIdx+rWidth > colIdx {
				return rIdx, 1, true
			}
			cIdx += rWidth
		}

		bIdx += rLen
		rIdx++
	}

	return rIdx, 0, false
}

// Count returns number of lines in the list starting from the line l.
// It does not take into account previous lines.
func (l *Line) Count() int {
	cnt := 1
	for l.Next != nil {
		l = l.Next
		cnt++
	}
	return cnt
}

// TrimRight trims trailing whitespaces and returns number of removed runes.
func (l *Line) TrimRight() int {
	trimCount := 0

	for i := len(l.Buf) - 1; i >= 0; i-- {
		r := l.Buf[i]
		if r == ' ' || r == '\t' {
			trimCount++
		} else {
			break
		}
	}
	l.Buf = l.Buf[0 : len(l.Buf)-trimCount]

	return trimCount
}

// Indent returns line indent as a string.
// An indent is a sequence of whitespace runes at the line beginning.
func (l *Line) Indent() string {
	sIdx := 0
	eIdx := 0
	for {
		r, rLen := utf8.DecodeRune(l.Buf[eIdx:])
		if r == utf8.RuneError {
			break
		}
		if !unicode.IsSpace(r) {
			break
		}
		eIdx += rLen
	}
	return string(l.Buf[sIdx:eIdx])
}

func (l *Line) HasOnlySpaces() bool {
	if len(l.Buf) == 0 {
		return false
	}

	bIdx := 0
	for bIdx < len(l.Buf) {
		r, rLen := utf8.DecodeRune(l.Buf[bIdx:])
		if !unicode.IsSpace(r) {
			return false
		}
		bIdx += rLen
	}
	return true
}
