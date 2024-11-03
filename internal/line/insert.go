package line

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/util"
)

func (l *Line) InsertRune(r rune, runeIdx int) {
	bIdx := l.BufIdx(runeIdx)
	runeLen := utf8.RuneLen(r)
	newLen := len(l.Buf) + runeLen

	if newLen > cap(l.Buf) {
		newBuf := make([]byte, newLen, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.Buf...)
		l.Buf = newBuf
	}

	// Can growing reslicing be done in a more efficient way?
	for i := 0; i < runeLen; i++ {
		l.Buf = append(l.Buf, 0)
	}

	for i := len(l.Buf) - 1; i > bIdx; i-- {
		l.Buf[i] = l.Buf[i-runeLen]
	}
	utf8.EncodeRune(l.Buf[bIdx:], r)
}

// InsertNewline inserts a newline at index idx and returns the new line.
func (l *Line) InsertNewline(runeIdx int) *Line {
	bIdx := l.BufIdx(runeIdx)

	newLine, _ := FromString(string(l.Buf[bIdx:]))
	l.Buf = l.Buf[0:bIdx]

	newLine.Prev = l
	newLine.Next = l.Next
	if l.Next != nil {
		l.Next.Prev = newLine
	}
	l.Next = newLine

	return newLine
}

func (l *Line) InsertString(s string, runeIdx int) {
	bIdx := l.BufIdx(runeIdx)

	newLen := len(l.Buf) + len(s)
	if newLen > cap(l.Buf) {
		newBuf := make([]byte, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.Buf...)
		l.Buf = newBuf
	}

	// TODO: Below code probably needs fix.
	right := l.Buf[bIdx:len(l.Buf)]
	l.Buf = l.Buf[0:bIdx]
	l.Buf = append(l.Buf, []byte(s)...)
	l.Buf = append(l.Buf, right...)
}

func (l *Line) Append(b []byte) {
	newLen := len(l.Buf) + len(b)
	if newLen > cap(l.Buf) {
		newBuf := make([]byte, 0, util.NextPowerOfTwo(newLen))
		newBuf = append(newBuf, l.Buf...)
		l.Buf = newBuf
	}

	l.Buf = append(l.Buf, b...)
}
