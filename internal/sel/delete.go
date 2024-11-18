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
		acts = append(acts, a)

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
		//s.deleteLine()
	} else if s.EndRuneIdx < rc {
		//s.deleteString()
	}

	return nil
}

func (s *Selection) deleteRune() *action.RuneDelete {
	r := s.Line.DeleteRune(s.StartRuneIdx)
	return &action.RuneDelete{Line: s.Line, Rune: r, RuneIdx: s.StartRuneIdx}
}
