package sel

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
)

func (s *Selection) Join() (action.Actions, *Selection) {
	if s.Next == nil {
		return s.joinSingleLine()
	}
	return s.joinMultiLine()
}

func (s *Selection) joinSingleLine() (action.Actions, *Selection) {
	l := s.Line
	rc := s.Line.RuneCount()
	newLine, trimmedCount := s.Line.Join(true)
	if newLine == nil {
		return nil, nil
	}

	s.Line = newLine
	s.Cursor.Line = newLine

	return action.Actions{
		&action.NewlineDelete{
			Line:         l,
			LineNum:      s.LineNum,
			RuneIdx:      rc,
			TrimmedCount: trimmedCount,
			NewLine:      newLine,
		},
	}, s
}

func (s *Selection) joinMultiLine() (action.Actions, *Selection) {
	actions := make(action.Actions, 0, 4)

	s = s.Last()
	s = s.Prev
	rc := 0 // New selection rune count

	var newLine *line.Line
	trimmedCount := 0
	for {
		l := s.Line
		lrc := s.Line.RuneCount()
		newLine, trimmedCount = s.Line.Join(true)

		// - 1 because we are in the middle of the selection,
		// so the line end rune is always selected.
		rc += s.Next.RuneCount() - trimmedCount - 1

		actions = append(
			actions,
			&action.NewlineDelete{
				Line:         l,
				LineNum:      s.LineNum,
				RuneIdx:      lrc,
				TrimmedCount: trimmedCount,
				NewLine:      newLine,
			},
		)

		if s.Prev == nil {
			break
		}
		s = s.Prev
	}

	rc += s.RuneCount()

	crIdx := s.StartRuneIdx
	if !s.CursorOnLeft() {
		crIdx += rc - 1
	}

	cur := cursor.New(newLine, s.LineNum, crIdx)

	newS := &Selection{
		Line:         newLine,
		LineNum:      s.LineNum,
		StartRuneIdx: s.StartRuneIdx,
		EndRuneIdx:   s.StartRuneIdx + rc - 1,
		Cursor:       cur,
		Prev:         nil,
		Next:         nil,
	}

	return actions, newS
}

func (s *Selection) LineDown() (action.Actions, *Selection) {
	actions := make(action.Actions, 0, 8)

	s = s.Last()
	for {
		ok := s.Line.Down()
		if !ok {
			break
		}
		actions = append(actions, &action.LineDown{Line: s.Line})

		s.LineNum++
		if s.Cursor != nil {
			s.Cursor.LineNum++
		}

		if s.Prev == nil {
			break
		}
		s = s.Prev
	}

	return actions, s
}

func (s *Selection) LineUp() (action.Actions, *Selection) {
	actions := make(action.Actions, 0, 8)

	for {
		ok := s.Line.Up()
		if !ok {
			break
		}
		actions = append(actions, &action.LineUp{Line: s.Line})

		s.LineNum--
		if s.Cursor != nil {
			s.Cursor.LineNum--
		}

		if s.Next == nil {
			break
		}
		s = s.Next
	}

	return actions, s
}
