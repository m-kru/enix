package sel

func (s *Selection) Right() {
	if s.CursorOnLeft() {
		s.rightCursorOnLeft()
	} else {
		s.rightCursorOnRight()
	}
}

func (s *Selection) rightCursorOnLeft() {
	panic("unimplemented")
}

func (s *Selection) rightCursorOnRight() {
	s = s.Last()
	if s.EndRuneIdx < s.Line.RuneCount() {
		s.EndRuneIdx++
		s.CursorIdx++
		return
	}

	panic("unimplemented")
}
