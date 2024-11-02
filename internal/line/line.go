package line

import (
	"fmt"

	"github.com/mattn/go-runewidth"
)

type Line struct {
	Buf []rune

	Prev *Line
	Next *Line
}

func (l *Line) RuneCount() int    { return len(l.Buf) }
func (l *Line) Rune(idx int) rune { return l.Buf[idx] }

func (l *Line) String() string {
	return string(l.Buf)
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
func (l *Line) Columns(tabWidth int) int {
	c := 0
	for _, r := range l.Buf {
		if r == '\t' {
			c += tabWidth
		} else {
			c += runewidth.RuneWidth(r)
		}
	}
	return c
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
func (l *Line) ColumnIdx(runeIdx int, tabWidth int) int {
	if runeIdx > len(l.Buf) {
		panic(fmt.Sprintf("rune idx (%d) > len(l.Buf) (%d)", runeIdx, len(l.Buf)))
	}

	col := 0
	for i, r := range l.Buf {
		if i == runeIdx {
			col += 1
			break
		} else {
			if r == '\t' {
				col += tabWidth - (col % tabWidth)
			} else {
				col += runewidth.RuneWidth(r)
			}
		}
	}

	if runeIdx == len(l.Buf) {
		col++
	}

	return col
}

// RuneIdx returns rune index for provided column index.
// The second return is a rune subcolumn index.
// The third return is false if column c does not exists in line.
func (l *Line) RuneIdx(col int, tabWidth int) (int, int, bool) {
	if col == 0 {
		panic("internal logic error")
	}

	c := 1
	for i, r := range l.Buf {
		rw := runewidth.RuneWidth(r)
		if r == '\t' {
			if c == col {
				return i, 0, true
			}

			width := tabWidth - ((c - 1) % tabWidth)
			if c+width > col {
				return i, col - c, true
			}
			c += width
		} else if rw == 1 {
			if c == col {
				return i, 0, true
			}
			c++
		} else {
			if c == col {
				return i, 0, true
			} else if c+rw > col {
				return i, 1, true
			}
			c += rw
		}
	}

	return 0, 0, false
}

// Count returns number of lines in the list starting from the line l.
// It does not take into account previous lines.
func (l *Line) Count() int {
	cnt := 1
	for {
		if l.Next == nil {
			break
		}
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
