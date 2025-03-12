package sel

import "github.com/m-kru/enix/internal/cursor"

func (s *Selection) adjust(c *cursor.Cursor) *Selection {
	first := s
	last := s.Last()
	oldCurOnLeft := s.CursorOnLeft()

	if oldCurOnLeft {
		first.Cursor = nil
	} else {
		last.Cursor = nil
	}

	if c.LineNum < first.LineNum {
		newFirst := &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.Line.RuneCount(),
			Cursor:       c,
			Prev:         nil,
			Next:         nil,
		}

		s.Prev = newFirst
		s := newFirst

		line := c.Line
		lineNum := c.LineNum
		for range first.LineNum - c.LineNum - 1 {
			line = line.Next
			lineNum++

			nextS := &Selection{
				Line:         line,
				LineNum:      lineNum,
				StartRuneIdx: 0,
				EndRuneIdx:   line.RuneCount(),
				Cursor:       nil,
				Prev:         s,
				Next:         nil,
			}

			s.Next = nextS
			s = nextS
		}

		s.Next = first
		if !oldCurOnLeft {
			first.EndRuneIdx = first.StartRuneIdx
		}
		first.StartRuneIdx = 0

		return newFirst
	} else if c.LineNum == first.LineNum {
		if c.RuneIdx <= first.EndRuneIdx {
			if !oldCurOnLeft {
				if c.RuneIdx >= first.StartRuneIdx {
					first.EndRuneIdx = c.RuneIdx
				} else {
					first.EndRuneIdx = first.StartRuneIdx
					first.StartRuneIdx = c.RuneIdx
				}
				first.Next = nil
			} else {
				first.StartRuneIdx = c.RuneIdx
			}
		} else {
			if oldCurOnLeft {
				first.StartRuneIdx = first.EndRuneIdx
			}
			first.EndRuneIdx = c.RuneIdx
		}

		first.Cursor = c
		return first
	} else if first.LineNum < c.LineNum && c.LineNum < last.LineNum {
		// First find the selection with cursor line
		s := first.Next
		for {
			if s.Line == c.Line {
				break
			}
			s = s.Next
		}

		s.Cursor = c

		if oldCurOnLeft {
			s.Prev = nil
			s.StartRuneIdx = c.RuneIdx
			first = s
		} else {
			s.EndRuneIdx = c.RuneIdx
			s.Next = nil
		}

		return first
	} else if c.LineNum == last.LineNum {
		if last.StartRuneIdx <= c.RuneIdx {
			if oldCurOnLeft {
				if c.RuneIdx <= last.EndRuneIdx {
					last.StartRuneIdx = c.RuneIdx
				} else {
					last.StartRuneIdx = last.EndRuneIdx
					last.EndRuneIdx = c.RuneIdx
				}
				first = last
			} else {
				last.EndRuneIdx = c.RuneIdx
			}
		} else {
			if !oldCurOnLeft {
				last.EndRuneIdx = last.StartRuneIdx
			}
			last.StartRuneIdx = c.RuneIdx
		}

		last.Cursor = c
		return first
	} else if c.LineNum > last.LineNum {
		newFirst := first
		if oldCurOnLeft {
			last.StartRuneIdx = last.EndRuneIdx
			newFirst = last
		}
		last.EndRuneIdx = last.Line.RuneCount()

		s := last
		line := last.Line
		lineNum := last.LineNum

		for range c.LineNum - last.LineNum - 1 {
			line = line.Next
			lineNum++

			nextS := &Selection{
				Line:         line,
				LineNum:      lineNum,
				StartRuneIdx: 0,
				EndRuneIdx:   c.Line.RuneCount(),
				Cursor:       nil,
				Prev:         s,
				Next:         nil,
			}

			s.Next = nextS
			s = nextS
		}

		newLast := &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
			Prev:         s,
			Next:         nil,
		}

		s.Next = newLast

		return newFirst
	}

	return first
}

func (s *Selection) Down() *Selection {
	c := s.GetCursor().Clone()
	c.Down()
	return s.adjust(c)
}

func (s *Selection) Left() *Selection {
	c := s.GetCursor().Clone()
	c.Left()
	return s.adjust(c)
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
		s.Cursor = cursor.New(s.Line, s.LineNum, lineRC)
		return first
	}

	if s.Line.Next == nil {
		return first
	}

	first.Cursor = nil
	lineRC = s.Line.Next.RuneCount()
	c := cursor.New(s.Line.Next, s.LineNum+1, lineRC)

	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   lineRC,
		Cursor:       c,
		Prev:         s,
		Next:         nil,
	}

	s.Next = newS

	return first
}

func (s *Selection) nextLineCursorOnRight() *Selection {
	first := s
	s = s.Last()

	lineRC := s.Line.RuneCount()
	if s.StartRuneIdx != 0 || s.EndRuneIdx != lineRC {
		s.StartRuneIdx = 0
		s.EndRuneIdx = lineRC
		s.Cursor.LineEnd()
		return first
	}

	if s.Line.Next == nil {
		return first
	}

	c := s.Cursor
	s.Cursor = nil
	c.Down()
	c.LineEnd()
	newS := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         s,
		Next:         nil,
	}

	s.Next = newS

	return first
}

func (s *Selection) LineEnd() *Selection {
	c := s.GetCursor().Clone()
	if c.RuneIdx == c.Line.RuneCount() {
		c.Right()
	}
	c.LineEnd()
	return s.adjust(c)
}

func (s *Selection) LineStart() *Selection {
	c := s.GetCursor().Clone()
	if c.RuneIdx == 0 {
		c.Left()
	}
	c.LineStart()
	return s.adjust(c)
}

func (s *Selection) PrevWordStart() *Selection {
	c := s.GetCursor().Clone()
	c.PrevWordStart()
	return s.adjust(c)
}

func (s *Selection) Right() *Selection {
	c := s.GetCursor().Clone()
	c.Right()
	return s.adjust(c)
}

func (s *Selection) Up() *Selection {
	c := s.GetCursor().Clone()
	c.Up()
	return s.adjust(c)
}

func (s *Selection) WordEnd() *Selection {
	c := s.GetCursor().Clone()
	c.WordEnd()
	return s.adjust(c)
}

func (s *Selection) TabEnd() *Selection {
	c := s.GetCursor().Clone()
	c.TabEnd()
	return s.adjust(c)
}
