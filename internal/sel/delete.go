package sel

import (
	"github.com/m-kru/enix/internal/action"
)

func (s *Selection) Delete() action.Action {
	acts := make(action.Actions, 0, 8)

	for {
		if s == nil {
			break
		}

		a := s.delete()
		if a != nil {
			acts = append(acts, a)
		}

		s = s.Next
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
			//return s.deleteNewline()
		}
	} else if s.StartRuneIdx == 0 && s.EndRuneIdx == rc {
		return s.deleteLine()
	} else if s.EndRuneIdx < rc {
		return s.deleteString()
	}

	return nil
}

func (s *Selection) deleteLine() action.Action {
	l := s.Line.Delete()
	if l == nil {
		return nil
	}
	return &action.LineDelete{Line: s.Line}
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
