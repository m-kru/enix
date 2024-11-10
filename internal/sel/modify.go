package sel

import "github.com/m-kru/enix/internal/cursor"

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

func (s *Selection) NextLine() *Selection {
	if s.CursorOnLeft() {
		return s.nextLineCursorOnLeft()
	} else {
		return s.nextLineCursorOnRight()
	}
}

func (s *Selection) nextLineCursorOnLeft() *Selection {
	lineRC := s.Line.RuneCount()
	if s.StartRuneIdx != 0 || s.EndRuneIdx != lineRC {
		s.StartRuneIdx = 0
		s.EndRuneIdx = lineRC
		s.Cursor.RuneIdx = 0
		return s
	}

	first := s
	s = s.Last()

	lineRC = s.Line.RuneCount()
	if s.StartRuneIdx != 0 || s.EndRuneIdx != lineRC {
		first.Cursor = nil
		s.StartRuneIdx = 0
		s.EndRuneIdx = lineRC
		s.Cursor = &cursor.Cursor{
			Line:    s.Line,
			LineNum: s.LineNum,
			ColIdx:  s.Line.ColumnIdx(lineRC),
			RuneIdx: lineRC,
		}
		return first
	}

	if s.Line.Next == nil {
		return first
	}

	first.Cursor = nil
	lineRC = s.Line.Next.RuneCount()
	c := &cursor.Cursor{
		Line:    s.Line.Next,
		LineNum: s.LineNum + 1,
		ColIdx:  s.Line.Next.ColumnIdx(lineRC),
		RuneIdx: lineRC,
	}

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   lineRC,
		Cursor:       c,
	}

	s.Next = newS
	newS.Prev = s

	return first
}

func (s *Selection) nextLineCursorOnRight() *Selection {
	first := s
	s = s.Last()

	lineRC := s.Line.RuneCount()
	if s.StartRuneIdx != 0 || s.EndRuneIdx != lineRC {
		s.StartRuneIdx = 0
		s.EndRuneIdx = lineRC
		s.Cursor.RuneIdx = lineRC
		s.Cursor.ColIdx = s.Line.ColumnIdx(lineRC)
		return first
	}

	if s.Line.Next == nil {
		return first
	}

	c := s.Cursor
	s.Cursor = nil
	c.Down()
	lineRC = c.Line.RuneCount()
	c.RuneIdx = lineRC
	c.ColIdx = c.Line.ColumnIdx(lineRC)
	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   lineRC,
		Cursor:       c,
	}

	s.Next = newS
	newS.Prev = s

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

	if s.Next != nil {
		s = s.Next
		s.Cursor = c
	}

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         s,
	}
	s.Next = newS
	s.Cursor = nil

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

func (s *Selection) WordEnd() *Selection {
	if s.CursorOnLeft() {
		return s.wordEndCursorOnLeft()
	} else {
		return s.wordEndCursorOnRight()
	}
}

func (s *Selection) wordEndCursorOnLeft() *Selection {
	c := s.Cursor
	c.WordEnd()

	if c.Line == s.Line {
		if c.RuneIdx < s.EndRuneIdx {
			s.StartRuneIdx = c.RuneIdx
		} else {
			s.StartRuneIdx = s.EndRuneIdx
			s.EndRuneIdx = c.RuneIdx
		}
		return s
	}

	first := s
	s.Cursor = nil

	if s.Next == nil {
		line := s.Line.Next
		lines := c.LineNum - s.LineNum - 1
		for i := 0; i < lines; i++ {
			sTmp := &Selection{
				Line:         line,
				LineNum:      s.LineNum + 1,
				StartRuneIdx: 0,
				EndRuneIdx:   line.RuneCount(),
				Prev:         s,
			}
			s.Next = sTmp
			s = sTmp
			line = line.Next
		}

		last := &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
			Prev:         s,
		}

		s.Next = last

		return first
	}

	s = s.Next
	s.Prev = nil

	if c.Line == s.Line {
		s.Cursor = c
		if c.RuneIdx > s.EndRuneIdx {
			s.EndRuneIdx = c.RuneIdx
		} else {
			s.StartRuneIdx = c.RuneIdx
		}
		return s
	}

	return s
}

func (s *Selection) wordEndCursorOnRight() *Selection {
	first := s
	s = s.Last()

	c := s.Cursor
	c.WordEnd()

	if c.Line == s.Line {
		s.EndRuneIdx = c.RuneIdx
		return first
	}

	s.Cursor = nil

	line := s.Line.Next
	lines := c.LineNum - s.LineNum - 1
	for i := 0; i < lines; i++ {
		sTmp := &Selection{
			Line:         line,
			LineNum:      s.LineNum + 1,
			StartRuneIdx: 0,
			EndRuneIdx:   line.RuneCount(),
			Prev:         s,
		}
		s.Next = sTmp
		s = sTmp
		line = line.Next
	}

	last := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         s,
	}

	s.Next = last

	return first
}
