package line

import "unicode/utf8"

func (l *Line) DeleteRune(runeIdx int) rune {
	if runeIdx == l.RuneCount() {
		if l.Next == nil {
			return 0
		}
	}

	bIdx := l.BufIdx(runeIdx)
	r, rLen := utf8.DecodeRune(l.Buf[bIdx:])

	l.Buf = append(l.Buf[:bIdx], l.Buf[bIdx+rLen:]...)

	return r
}
