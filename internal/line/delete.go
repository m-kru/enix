package line

import "unicode/utf8"

func (l *Line) DeleteRune(runeIdx int) *Line {
	if runeIdx == l.RuneCount() {
		if l.Next == nil {
			return nil
		}

		// Newline deletion
		return l.Join(false)
	}

	bIdx := l.BufIdx(runeIdx)
	_, rLen := utf8.DecodeRune(l.Buf[bIdx:])

	l.Buf = append(l.Buf[:bIdx], l.Buf[bIdx+rLen:]...)

	return nil
}
