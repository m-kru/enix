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
	panic("unimplemented")
}

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
