package sel

import (
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"
)

type Selection struct {
	Line         *line.Line
	LineNum      int
	StartRuneIdx int
	EndRuneIdx   int
	CursorIdx    int // -1 if given selection doesn't have cursor

	Prev *Selection
	Next *Selection
}

func (s *Selection) Last() *Selection {
	for {
		if s.Next == nil {
			break
		}
		s = s.Next
	}
	return s
}

func (s *Selection) HasCursor() bool {
	return s.CursorIdx >= 0
}

func (s *Selection) CursorOnLeft() bool {
	return s.HasCursor() && s.CursorIdx == s.StartRuneIdx
}

func (s *Selection) CursorOnRight() bool {
	return s.HasCursor() && s.CursorIdx == s.EndRuneIdx
}

func (s *Selection) View() view.View {
	startCol := s.Line.ColumnIdx(s.StartRuneIdx)
	endCol := s.Line.ColumnIdx(s.EndRuneIdx)

	v := view.View{
		Line:   s.LineNum,
		Column: startCol,
		Width:  endCol - startCol + 1,
		Height: 1,
	}

	return v
}

func (s *Selection) FullView() view.View {
	v := view.View{Line: s.LineNum}

	col1 := s.Line.ColumnIdx(s.StartRuneIdx)

	for {
		if s.Next == nil {
			break
		}
		s = s.Next
	}

	v.Height = s.LineNum - v.Line

	col2 := s.Line.ColumnIdx(s.StartRuneIdx)
	if col1 <= col2 {
		v.Column = col1
		v.Width = col2 - col1 + 1
	} else {
		v.Column = col2
		v.Width = col1 - col2 + 1
	}

	return v
}
