package line

import "github.com/m-kru/enix/internal/util"

func (l *Line) InsertRune(r rune, idx int) {
	newLen := len(l.buf) + 1
	if newLen > cap(l.buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.buf...)
		l.buf = newBuf
	}

	// Append one rune at the end to increase the rune slice length.
	l.buf = append(l.buf, ' ')
	for i := l.Len() - 1; i > idx; i-- {
		l.buf[i] = l.buf[i-1]
	}
	l.buf[idx] = r
}

// InsertNewline inserts a newline at index idx and returns the new line.
func (l *Line) InsertNewline(idx int) *Line {
	newLine := FromString(string(l.buf[idx:l.Len()]))
	l.buf = l.buf[0:idx]

	newLine.Prev = l
	newLine.Next = l.Next
	if l.Next != nil {
		l.Next.Prev = newLine
	}
	l.Next = newLine

	return newLine
}

func (l *Line) Insert(s string, idx int) {
	newLen := len(l.buf) + len(s)
	if newLen > cap(l.buf) {
		newBuf := make([]rune, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.buf...)
		l.buf = newBuf
	}

	// TODO: Below code probably needs fix.
	right := l.buf[idx:len(l.buf)]
	l.buf = l.buf[0:idx]
	l.buf = append(l.buf, []rune(s)...)
	l.buf = append(l.buf, right...)
}
