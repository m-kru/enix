package sel

import (
	"github.com/m-kru/enix/internal/action"
)

func (s *Selection) Join() (action.Actions, *Selection) {
	if s.Next == nil {
		return s.joinSingleLine()
	}
	return s.joinMultiLine()
}

func (s *Selection) joinSingleLine() (action.Actions, *Selection) {
	l1 := s.Line
	l2 := l1.Next
	rc := s.Line.RuneCount()
	newLine := s.Line.Join(true)
	if newLine == nil {
		return nil, nil
	}

	s.Line = newLine

	return action.Actions{
		&action.NewlineDelete{
			Line1:    l1,
			Line1Num: s.LineNum,
			RuneIdx:  rc,
			Line2:    l2,
			NewLine:  newLine,
		},
	}, s
}

func (s *Selection) joinMultiLine() (action.Actions, *Selection) {
	actions := make(action.Actions, 0, 4)

	s = s.Last()
	//rc := 0 // New selection rune count

	for {
		if s == nil {
			break
		}

		s = s.Prev
	}

	newS := &Selection{}

	return actions, newS
}
