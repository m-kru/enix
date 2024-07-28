package line

import (
	"fmt"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

type Line struct {
	Buf string

	Prev *Line
	Next *Line
}

func (l *Line) Len() int { return len(l.Buf) }

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
	return len(l.Buf)
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
	l.Buf = fmt.Sprintf("%s%s", l.Buf, s)
}

func (l *Line) Delete(idx int, size int) {
	l.Buf = l.Buf[0:idx] + l.Buf[idx+1:len(l.Buf)]
}

func (l *Line) InsertRune(r rune, idx int) {
	left := l.Buf[0:idx]
	right := l.Buf[idx:len(l.Buf)]
	l.Buf = fmt.Sprintf("%s%c%s", left, r, right)
}

func (l *Line) Render(colors *cfg.Colorscheme, frame frame.Frame, view view.View) {
	i := 0
	for _, r := range l.Buf[view.Column-1 : len(l.Buf)] {
		if i == frame.Width-1 {
			break
		}

		frame.SetContent(i, 0, r, colors.Default)
		i++
	}

	for i < frame.Width {
		frame.SetContent(i, 0, ' ', colors.Default)
		i++
	}
}
