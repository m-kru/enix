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

	if s.Line.Next == nil {
		// Do nothing, this is end of text
		return
	}

	newS := &Selection{
		Line:         s.Line.Next,
		LineNum:      s.LineNum + 1,
		StartRuneIdx: 0,
		EndRuneIdx:   0,
		CursorIdx:    0,
	}

	newS.Prev = s
	s.Next = newS
	s.CursorIdx = -1
}
