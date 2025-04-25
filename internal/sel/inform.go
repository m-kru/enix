package sel

import (
	"github.com/m-kru/enix/internal/action"
)

func (s *Selection) Inform(act action.Action, informCursor bool) {
	// There is a one peculiar case, when inserting a newline at the end of selection
	// creates a new subselection at the end. Such a subselection shall not be informed
	// about actions. This is why the last variable is required to track if the previous
	// subselection is the last one that shall be informed about actions.
	last := false
	for !last {
		last = s.Next == nil

		s.inform(act)
		if s.Cursor != nil && informCursor {
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
		s.informNewlineInsert(a)
	case *action.RuneDelete:
		s.informRuneDelete(a)
	case *action.RuneInsert:
		s.informRuneInsert(a)
	case *action.StringDelete:
		s.informStringDelete(a)
	case *action.StringInsert:
		s.informStringInsert(a)
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
	if s.LineNum < nd.LineNum {
		return
	}

	if s.Line == nd.Line {
		s.Line = nd.NewLine
		return
	}

	if s.Line == nd.Line.Next {
		s.Line = nd.NewLine
		s.StartRuneIdx += nd.RuneIdx
		s.EndRuneIdx += nd.RuneIdx
	}

	s.LineNum--
}

func (s *Selection) informNewlineInsert(ni *action.NewlineInsert) {
	if s.LineNum < ni.LineNum {
		return
	}

	if s.LineNum > ni.LineNum {
		s.LineNum++
	}

	// Assume newline insert didn't split the selection.
	// This should not be possible because of the way cursors
	// and selections work.
	// Newline insert can split selection marks. However, in such
	// a case, the mark shall be destroyed.
	if s.EndRuneIdx < ni.RuneIdx {
		s.Line = ni.NewLine
	} else if s.EndRuneIdx == ni.RuneIdx {
		// This is possbile only if this is the last subselection.
		s.Line = ni.NewLine
		c := s.Cursor
		s.Cursor = nil
		newS := &Selection{
			Line:         ni.NewLine.Next,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
			Prev:         s,
			Next:         nil,
		}
		s.Next = newS
	} else {
		s.Line = ni.NewLine.Next
		s.StartRuneIdx -= ni.RuneIdx
		s.EndRuneIdx -= ni.RuneIdx
		s.LineNum++
	}
}

func (s *Selection) informRuneDelete(rd *action.RuneDelete) {
	if s.Line == rd.Line && rd.RuneIdx < s.StartRuneIdx {
		s.StartRuneIdx--
		s.EndRuneIdx--
	}
}

func (s *Selection) informRuneInsert(ri *action.RuneInsert) {
	if s.Line != ri.Line {
		return
	}

	if ri.RuneIdx <= s.StartRuneIdx {
		s.StartRuneIdx++
		s.EndRuneIdx++
	}

	if ri.RuneIdx == s.EndRuneIdx {
		s.EndRuneIdx++
	}
}

func (s *Selection) informStringDelete(sd *action.StringDelete) {
	if s.Line == sd.Line && sd.StartRuneIdx < s.StartRuneIdx {
		s.StartRuneIdx -= sd.RuneCount
		s.EndRuneIdx -= sd.RuneCount
	}
}

func (s *Selection) informStringInsert(si *action.StringInsert) {
	if s.Line != si.Line {
		return
	}

	if si.StartRuneIdx < s.StartRuneIdx {
		s.StartRuneIdx += si.RuneCount
		s.EndRuneIdx += si.RuneCount
	} else if si.StartRuneIdx == s.StartRuneIdx {
		if s.Prev == nil {
			s.StartRuneIdx += si.RuneCount
		}
		s.EndRuneIdx += si.RuneCount
	}
}
