package line

import (
	"github.com/m-kru/enix/internal/util"
)

func (l *Line) InsertRune(r rune, idx int) {
	newLen := len(l.Buf) + 1
	if newLen > cap(l.Buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.Buf...)
		l.Buf = newBuf
	}

	// Append one rune at the end to increase the rune slice length.
	l.Buf = append(l.Buf, ' ')
	for i := l.RuneCount() - 1; i > idx; i-- {
		l.Buf[i] = l.Buf[i-1]
	}
	l.Buf[idx] = r
}

// InsertNewline inserts a newline at index idx and returns the new line.
func (l *Line) InsertNewline(idx int) *Line {
	newLine, _ := FromString(string(l.Buf[idx:l.RuneCount()]))
	l.Buf = l.Buf[0:idx]

	newLine.Prev = l
	newLine.Next = l.Next
	if l.Next != nil {
		l.Next.Prev = newLine
	}
	l.Next = newLine

	return newLine
}

func (l *Line) Insert(s string, idx int) {
	newLen := len(l.Buf) + len(s)
	if newLen > cap(l.Buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.Buf...)
		l.Buf = newBuf
	}

	// TODO: Below code probably needs fix.
	right := l.Buf[idx:len(l.Buf)]
	l.Buf = l.Buf[0:idx]
	l.Buf = append(l.Buf, []rune(s)...)
	l.Buf = append(l.Buf, right...)
}

func (l *Line) Append(s string) {
	newLen := len(l.Buf) + len(s)
	if newLen > cap(l.Buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.Buf...)
		l.Buf = newBuf
	}

	l.Buf = append(l.Buf, []rune(s)...)
}
