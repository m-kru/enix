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
		if s.Cursor != nil {
			s.Cursor.Inform(act)
		}

		s = s.Next
	}
}

func (s *Selection) inform(act action.Action) {
	switch a := act.(type) {
	case action.Actions:
		for _, subA := range a {
			s.inform(subA)
		}
	case *action.LineDelete:
		s.informLineDelete(a)
	case *action.LineDown:
		//s.informLineDown(a)
	case *action.LineInsert:
		s.informLineInsert(a)
	case *action.LineUp:
		//s.informLineUp(a)
	case *action.NewlineDelete:
		s.informNewlineDelete(a)
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

func (s *Selection) informLineDelete(ld *action.LineDelete) {
	if ld.LineNum > s.LineNum {
		return
	}
	s.LineNum--
}

func (s *Selection) informLineInsert(li *action.LineInsert) {
	if li.LineNum > s.LineNum {
		return
	}
	s.LineNum++
}

func (s *Selection) informNewlineDelete(nd *action.NewlineDelete) {
	if s.LineNum < nd.Line1Num {
		return
	}

	if s.Line == nd.Line1 {
		s.Line = nd.NewLine
		return
	}

	if s.Line == nd.Line2 {
		s.Line = nd.NewLine
		s.StartRuneIdx += nd.RuneIdx
		s.EndRuneIdx += nd.RuneIdx
	}

	s.LineNum--
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
