package sel

func (s *Selection) Left() *Selection {
	if s.CursorOnLeft() {
		return s.leftCursorOnLeft()
	} else {
		return s.leftCursorOnRight()
	}
}

func (s *Selection) leftCursorOnLeft() *Selection {
	if s.StartRuneIdx > 0 {
		s.StartRuneIdx--
		s.CursorIdx--
		return s
	}

	if s.Line.Prev == nil {
		// Do nothing, this is start of text
		return s
	}

	newS := &Selection{
		Line:         s.Line.Prev,
		LineNum:      s.LineNum - 1,
		StartRuneIdx: s.Line.Prev.RuneCount(),
		EndRuneIdx:   s.Line.Prev.RuneCount(),
		CursorIdx:    s.Line.Prev.RuneCount(),
	}

	newS.Next = s
	s.Prev = newS
	s.CursorIdx = -1

	return newS
}

func (s *Selection) leftCursorOnRight() *Selection {
	first := s

	s = s.Last()
	if s.EndRuneIdx > 0 {
		s.EndRuneIdx--
		s.CursorIdx--
		return first
	}

	panic("unimplemented")
}

func (s *Selection) Right() *Selection {
	if s.CursorOnLeft() {
		return s.rightCursorOnLeft()
	} else {
		return s.rightCursorOnRight()
	}
}

func (s *Selection) rightCursorOnLeft() *Selection {
	if s.StartRuneIdx < s.Line.RuneCount() {
		if s.StartRuneIdx < s.EndRuneIdx {
			s.StartRuneIdx++
			s.CursorIdx++
		} else {
			s.EndRuneIdx++
			s.CursorIdx++
		}
		return s
	}

	if s.Next != nil {
		s = s.Next
		s.Prev = nil
		s.CursorIdx = 0
		return s
	}

	if s.Line.Next == nil {
		return s
	}

	newS := &Selection{
		Line:         s.Line.Next,
		LineNum:      s.LineNum + 1,
		StartRuneIdx: 0,
		EndRuneIdx:   0,
		CursorIdx:    0,
	}

	s.CursorIdx = -1
	s.Next = newS
	newS.Prev = s

	return s
}

func (s *Selection) rightCursorOnRight() *Selection {
	first := s

	s = s.Last()
	if s.EndRuneIdx < s.Line.RuneCount() {
		s.EndRuneIdx++
		s.CursorIdx++
		return first
	}

	if s.Line.Next == nil {
		// Do nothing, this is end of text
		return first
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

	return first
}
