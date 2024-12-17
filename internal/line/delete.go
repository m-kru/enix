package line

import "unicode/utf8"

// Clear empties the line buffer.
func (l *Line) Clear() {
	l.Buf = l.Buf[0:0]
}

func (l *Line) Delete() *Line {
	// First line can't be deleted if it is the only line.
	if l.Prev == nil && l.Next == nil {
		return nil
	}

	if l.Prev != nil {
		l.Prev.Next = l.Next
	}
	if l.Next != nil {
		l.Next.Prev = l.Prev
	}

	if l.Next != nil {
		return l.Next
	}

	return l.Prev
}

func (l *Line) DeleteRune(rIdx int) rune {
	if rIdx == l.RuneCount() {
		if l.Next == nil {
			return 0
		}
	}

	bIdx := l.BufIdx(rIdx)
	r, rLen := utf8.DecodeRune(l.Buf[bIdx:])

	l.Buf = append(l.Buf[:bIdx], l.Buf[bIdx+rLen:]...)

	return r
}

func (l *Line) DeleteString(srIdx int, erIdx int) string {
	sbIdx := l.BufIdx(srIdx)
	ebIdx := l.BufIdx(erIdx + 1)

	str := string(l.Buf[sbIdx:ebIdx])

	l.Buf = append(l.Buf[:sbIdx], l.Buf[ebIdx:]...)

	return str
}
