package sel

func (s *Selection) Clone() *Selection {
	first := &Selection{
		Line:         s.Line,
		LineNum:      s.LineNum,
		StartRuneIdx: s.StartRuneIdx,
		EndRuneIdx:   s.EndRuneIdx,
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}

	if s.Cursor != nil {
		first.Cursor = s.Cursor.Clone()
	}

	prevS := first
	s = s.Next
	for {
		if s == nil {
			break
		}

		nextS := &Selection{
			Line:         s.Line,
			LineNum:      s.LineNum,
			StartRuneIdx: s.StartRuneIdx,
			EndRuneIdx:   s.EndRuneIdx,
			Cursor:       nil,
			Prev:         prevS,
			Next:         nil,
		}

		if s.Cursor != nil {
			nextS.Cursor = s.Cursor.Clone()
		}

		prevS.Next = nextS
		prevS = nextS

		s = s.Next
	}

	return first
}

func Clone(sels []*Selection) []*Selection {
	if sels == nil {
		return nil
	}

	ss := make([]*Selection, 0, len(sels))

	for _, s := range sels {
		newS := s.Clone()
		ss = append(ss, newS)
	}

	return ss
}
