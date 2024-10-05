package line

func (l *Line) DeleteRune(idx int) *Line {
	if idx == l.Len() {
		if l.Prev == nil {
			return nil
		}
		// Newline deletion
		panic("unimplemented")
	}

	l.Buf = append(l.Buf[:idx], l.Buf[idx+1:]...)

	return nil
}
