package line

import (
	"fmt"

	"github.com/m-kru/enix/internal/util"
)

type Line struct {
	buf []rune

	Prev *Line
	Next *Line
}

func (l *Line) Len() int       { return len(l.buf) }
func (l *Line) String() string { return string(l.buf) }

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

// LineNum is an alias to Num() to satisfy Visible interface.
func (l *Line) LineNum() int {
	return l.Num()
}

func (l *Line) Column() int {
	// TODO: Handle tabs.
	return len(l.buf)
}

// Get returns nth line.
// It panics if there is less than n lines.
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

	panic(fmt.Sprintf("cannot get %d ", n))
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

func (l *Line) Append(s string) {
	newLen := len(l.buf) + len(s)
	if newLen > cap(l.buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.buf...)
		l.buf = newBuf
	}

	l.buf = append(l.buf, []rune(s)...)
}

func (l *Line) Delete(idx int, size int) {
	l.buf = append(l.buf[0:idx], l.buf[idx+1:len(l.buf)]...)
}

func (l *Line) InsertRune(r rune, idx int) {
	newLen := len(l.buf) + 1
	if newLen > cap(l.buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.buf...)
		l.buf = newBuf
	}

	right := l.buf[idx:len(l.buf)]
	l.buf = l.buf[0:idx]
	l.buf = append(l.buf, r)
	l.buf = append(l.buf, right...)
}

func (l *Line) Insert(s string, idx int) {
	newLen := len(l.buf) + len(s)
	if newLen > cap(l.buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.buf...)
		l.buf = newBuf
	}

	right := l.buf[idx:len(l.buf)]
	l.buf = l.buf[0:idx]
	l.buf = append(l.buf, []rune(s)...)
	l.buf = append(l.buf, right...)
}
