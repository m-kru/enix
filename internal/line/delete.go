package line

func (l *Line) DeleteRune(idx int) *Line {
	if idx == l.RuneCount() {
		if l.Next == nil {
			return nil
		}

		// Newline deletion
		return l.Join(false)
	}

	l.Buf = append(l.Buf[:idx], l.Buf[idx+1:]...)

	return nil
}
