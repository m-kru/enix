package sel

import (
	"github.com/m-kru/enix/internal/action"
)

func (s *Selection) Delete() action.Actions {
	acts := make(action.Actions, 0, 8)

	s = s.Last()

	for {
		if s == nil {
			break
		}

		a := s.delete()
		if a != nil {
			acts = append(acts, a)
		}

		s = s.Prev
	}

	return acts
}

// subselection delete
func (s *Selection) delete() action.Action {
	rc := s.Line.RuneCount()

	if s.StartRuneIdx == s.EndRuneIdx {
		if s.EndRuneIdx < rc {
			return s.deleteRune()
		} else {
			return s.deleteNewline()
		}
	} else if s.StartRuneIdx == 0 && s.EndRuneIdx == rc {
		return s.deleteLine()
	} else if s.EndRuneIdx < rc {
		return s.deleteString()
	} else {
		return s.deleteStringAndNewline()
	}
}

func (s *Selection) deleteLine() action.Action {
	l := s.Line.Delete()
	if l == nil {
		str := s.Line.DeleteString(0, s.EndRuneIdx-1)
		return &action.StringDelete{
			Line:         s.Line,
			Str:          str,
			StartRuneIdx: s.StartRuneIdx,
			RuneCount:    s.EndRuneIdx - s.StartRuneIdx,
		}
	}

	ld := &action.LineDelete{
		Line:    s.Line,
		LineNum: s.LineNum,
		NewLine: l,
	}

	if l == s.Line.Prev {
		s.LineNum--
	}
	s.Line = l

	return ld
}

func (s *Selection) deleteNewline() action.Action {
	l1 := s.Line
	l2 := l1.Next
	rc := s.Line.RuneCount()

	newLine, _ := s.Line.Join(false)
	if newLine == nil {
		return nil
	}

	// XXX: This is nasty. We modify the selection so that it is possible to
	// create cursor in a valid place after the deletion. We can do this, because
	// we assume after the deletion the selection is destroyed and never used again.
	s.Line = newLine

	return &action.NewlineDelete{
		Line1:        l1,
		Line1Num:     s.LineNum,
		RuneIdx:      rc,
		Line2:        l2,
		TrimmedCount: 0,
		NewLine:      newLine,
	}
}

func (s *Selection) deleteRune() *action.RuneDelete {
	r := s.Line.DeleteRune(s.StartRuneIdx)
	return &action.RuneDelete{Line: s.Line, Rune: r, RuneIdx: s.StartRuneIdx}
}

func (s *Selection) deleteString() *action.StringDelete {
	str := s.Line.DeleteString(s.StartRuneIdx, s.EndRuneIdx)
	return &action.StringDelete{
		Line:         s.Line,
		Str:          str,
		StartRuneIdx: s.StartRuneIdx,
		RuneCount:    s.EndRuneIdx - s.StartRuneIdx + 1,
	}
}

func (s *Selection) deleteStringAndNewline() action.Actions {
	acts := make(action.Actions, 0, 2)

	// String delete
	str := s.Line.DeleteString(s.StartRuneIdx, s.EndRuneIdx)
	acts = append(acts,
		&action.StringDelete{
			Line:         s.Line,
			Str:          str,
			StartRuneIdx: s.StartRuneIdx,
			RuneCount:    s.EndRuneIdx - s.StartRuneIdx + 1,
		},
	)

	// Newline delete
	l1 := s.Line
	l2 := l1.Next
	rc := s.Line.RuneCount()

	newLine, _ := s.Line.Join(false)
	if newLine == nil {
		return acts
	}

	// XXX: The same nasty trick as in case of deleteNewline.
	s.Line = newLine

	acts = append(acts,
		&action.NewlineDelete{
			Line1:        l1,
			Line1Num:     s.LineNum,
			RuneIdx:      rc,
			Line2:        l2,
			TrimmedCount: 0,
			NewLine:      newLine,
		},
	)

	return acts
}
