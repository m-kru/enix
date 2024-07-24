package line

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
)

type Line struct {
	Screen tcell.Screen

	Colors *cfg.Colorscheme

	Buf string

	Prev *Line
	Next *Line
}

func (l *Line) Len() int { return len(l.Buf) }

// Count returns number of lines in the list starting from line l.
// It does not take int account previous lines.
func (l *Line) Count() int {
	c := 1
	for {
		if l.Next == nil {
			break
		}
		l = l.Next
		c++
	}
	return c
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

func (l *Line) Render(x, y, width, offset int) {
	i := 0
	for _, r := range l.Buf[offset:len(l.Buf)] {
		if i == width-1 {
			break
		}

		l.Screen.SetContent(i+x, y, r, nil, l.Colors.Default)
		i++
	}

	for i < width {
		l.Screen.SetContent(i+x, y, ' ', nil, l.Colors.Default)
		i++
	}
}
