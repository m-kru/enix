package sel

func (s *Selection) Down() *Selection {
	if s.CursorOnLeft() {
		return s.downCursorOnLeft()
	} else {
		return s.downCursorOnRight()
	}
}

func (s *Selection) downCursorOnLeft() *Selection {
	if s.Cursor.Line.Next == nil {
		return s
	}

	if s.Next == nil {
		c := s.Cursor

		s.StartRuneIdx = s.EndRuneIdx
		s.EndRuneIdx = s.Line.RuneCount()
		s.Cursor = nil

		c.Down()

		newS := &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
		}

		s.Next = newS
		newS.Prev = s

		return s
	}

	c := s.Cursor
	c.Down()
	s = s.Next
	s.Prev = nil
	s.Cursor = c

	if c.RuneIdx < s.EndRuneIdx {
		s.StartRuneIdx = c.RuneIdx
	} else {
		s.StartRuneIdx = s.EndRuneIdx
		s.EndRuneIdx = c.RuneIdx
	}

	return s
}

func (s *Selection) downCursorOnRight() *Selection {
	first := s
	s = s.Last()

	c := s.Cursor
	c.Down()

	if c.LineNum == s.LineNum {
		return first
	}

	s.EndRuneIdx = s.Line.RuneCount()

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
	}

	s.Next = newS
	newS.Prev = s
	s.Cursor = nil

	return first
}

func (s *Selection) Left() *Selection {
	if s.CursorOnLeft() {
		return s.leftCursorOnLeft()
	} else {
		return s.leftCursorOnRight()
	}
}

func (s *Selection) leftCursorOnLeft() *Selection {
	c := s.Cursor
	c.Left()

	if c.RuneIdx == s.StartRuneIdx && c.LineNum == s.LineNum {
		return s
	}

	if c.LineNum == s.LineNum {
		s.StartRuneIdx = c.RuneIdx
		return s
	}

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
	}

	newS.Next = s
	s.Prev = newS
	s.Cursor = nil

	return newS
}

func (s *Selection) leftCursorOnRight() *Selection {
	first := s

	s = s.Last()
	c := s.Cursor
	c.Left()

	if c.LineNum == s.LineNum {
		s.EndRuneIdx = c.RuneIdx
		return first
	}

	s = s.Prev
	s.Cursor = c
	s.Next = nil

	return first
}

func (s *Selection) Right() *Selection {
	if s.CursorOnLeft() {
		return s.rightCursorOnLeft()
	} else {
		return s.rightCursorOnRight()
	}
}

func (s *Selection) rightCursorOnLeft() *Selection {
	c := s.Cursor
	c.Right()

	if c.LineNum == s.LineNum {
		if s.StartRuneIdx != s.EndRuneIdx {
			s.StartRuneIdx = c.RuneIdx
		} else {
			s.EndRuneIdx = c.RuneIdx
		}

		return s
	}

	s = s.Next
	s.Cursor = c

	return s
}

func (s *Selection) rightCursorOnRight() *Selection {
	first := s

	s = s.Last()
	c := s.Cursor
	c.Right()

	if c.RuneIdx == s.EndRuneIdx && c.LineNum == s.LineNum {
		return first
	}

	if c.LineNum == s.LineNum {
		s.EndRuneIdx++
		return first
	}

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
	}

	newS.Prev = s
	s.Next = newS
	s.Cursor = nil

	return first
}

func (s *Selection) Up() *Selection {
	if s.CursorOnLeft() {
		return s.upCursorOnLeft()
	} else {
		return s.upCursorOnRight()
	}
}

func (s *Selection) upCursorOnLeft() *Selection {
	c := s.Cursor

	if c.Line.Prev == nil {
		return s
	}

	c.Up()
	s.StartRuneIdx = 0
	s.Cursor = nil

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       c,
	}

	s.Prev = newS
	newS.Next = s

	return newS
}

func (s *Selection) upCursorOnRight() *Selection {
	if s.Next == nil {
		c := s.Cursor

		s.EndRuneIdx = s.StartRuneIdx
		s.StartRuneIdx = 0
		s.Cursor = nil

		c.Up()

		newS := &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.Line.RuneCount(),
			Cursor:       c,
		}

		s.Prev = newS
		newS.Next = s

		return newS
	}

	first := s
	s = s.Last()

	c := s.Cursor
	c.Up()
	s = s.Prev
	s.Next = nil
	s.Cursor = c

	if c.RuneIdx > s.StartRuneIdx {
		s.EndRuneIdx = c.RuneIdx
	} else {
		s.EndRuneIdx = s.StartRuneIdx
		s.StartRuneIdx = c.RuneIdx
	}

	return first
}
