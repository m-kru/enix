package sel

import (
	"github.com/m-kru/enix/internal/action"
)

func (s *Selection) Inform(act action.Action) {
	switch a := act.(type) {
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
	}
}

func (s *Selection) informRuneDelete(rd *action.RuneDelete) {
	if s.Line == rd.Line && rd.RuneIdx < s.StartRuneIdx {
		s.StartRuneIdx--
		s.EndRuneIdx--
	}
}
