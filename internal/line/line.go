package line

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
)

type Line struct {
	Screen tcell.Screen
	StartX int
	EndX   int
	StartY int

	Colors *cfg.Colorscheme

	Buf string

	Prev *Line
	Next *Line
}

func (l *Line) Len() int { return len(l.Buf) }

func (l *Line) Delete(idx int, size int) {
	l.Buf = l.Buf[0:idx] + l.Buf[idx+1:len(l.Buf)]
}

func (l *Line) InsertRune(r rune, idx int) {
	left := l.Buf[0:idx]
	right := l.Buf[idx:len(l.Buf)]
	l.Buf = fmt.Sprintf("%s%c%s", left, r, right)
}

func (l *Line) Render() {
	i := 0
	for _, r := range l.Buf {
		if i+l.StartX > l.EndX {
			break
		}

		l.Screen.SetContent(i+l.StartX, l.StartY, r, nil, l.Colors.Default)
		i++
	}

	for i+l.StartX < l.EndX {
		l.Screen.SetContent(i+l.StartX, l.StartY, ' ', nil, l.Colors.Default)
		i++
	}
}
