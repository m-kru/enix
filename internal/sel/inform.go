package sel

import (
	"github.com/m-kru/enix/internal/action"
)

func (s *Selection) Inform(act action.Action) {
	for {
		if s == nil {
			break
		}

		s.inform(act)

		s = s.Next
	}
}

func (s *Selection) inform(act action.Action) {
	switch a := act.(type) {
	case action.Actions:
		for _, subA := range a {
			s.inform(subA)
		}
	case *action.LineDown:
		//s.informLineDown(a)
	case *action.LineUp:
		//s.informLineUp(a)
	case *action.NewlineDelete:
		//s.informNewlineDelete(a)
	case *action.NewlineInsert:
		//s.informNewlineInsert(a)
	case *action.RuneDelete:
		s.informRuneDelete(a)
	case *action.RuneInsert:
		//s.informRuneInsert(a)
	case *action.StringDelete:
		s.informStringDelete(a)
	}
}

func (s *Selection) informRuneDelete(rd *action.RuneDelete) {
	if s.Line == rd.Line && rd.RuneIdx < s.StartRuneIdx {
		s.StartRuneIdx--
		s.EndRuneIdx--
	}
}

func (s *Selection) informStringDelete(sd *action.StringDelete) {
	if s.Line == sd.Line && sd.StartRuneIdx < s.StartRuneIdx {
		s.StartRuneIdx -= sd.RuneCount
		s.EndRuneIdx -= sd.RuneCount
	}
}
