package sel

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"
)

type Selection struct {
	Line         *line.Line
	LineNum      int
	StartRuneIdx int
	EndRuneIdx   int
	Cursor       *cursor.Cursor

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

func (s *Selection) CursorOnLeft() bool {
	return s.Cursor != nil && s.Cursor.RuneIdx == s.StartRuneIdx
}

func (s *Selection) CursorOnRight() bool {
	return s.Cursor != nil && s.Cursor.RuneIdx == s.EndRuneIdx
}

func (s *Selection) GetCursor() *cursor.Cursor {
	if !s.CursorOnLeft() {
		s = s.Last()
	}

	return s.Cursor
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

	v.Height = s.LineNum - v.Line + 1

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

func (s *Selection) Overlaps(s2 *Selection) bool {
	for {
		if s == nil {
			break
		}

		subs2 := s2
		for {
			if subs2 == nil {
				break
			}

			if s.LineNum == subs2.LineNum &&
				((s.EndRuneIdx >= subs2.StartRuneIdx && s.StartRuneIdx <= subs2.StartRuneIdx) ||
					(subs2.EndRuneIdx >= s.StartRuneIdx && subs2.StartRuneIdx <= s.StartRuneIdx)) {
				return true
			}

			subs2 = subs2.Next
		}

		s = s.Next
	}

	return false
}

// Merge merges two selections. However, it does not check
// if selections overlap. User must first expliitly check it
// by calling the Overlaps function.
func (s *Selection) Merge(s2 *Selection) *Selection {
	if s2.LineNum < s.LineNum || (s2.LineNum == s.LineNum && s2.StartRuneIdx < s.StartRuneIdx) {
		tmp := s
		s = s2
		s2 = tmp
	}

	first := s

	for {
		if s.LineNum < s2.LineNum {
			s = s.Next
			continue
		}

		if s.Cursor != nil && s.Cursor.RuneIdx >= s2.StartRuneIdx {
			if s2.Cursor != nil {
				s.Cursor = s2.Cursor
			} else {
				s.Cursor = nil
			}
		}
		if s2.Cursor != nil && s2.Cursor.RuneIdx <= s.EndRuneIdx {
			s2.Cursor = nil
		}

		s.EndRuneIdx = s2.EndRuneIdx
		s.Next = s2.Next
		break
	}

	return first
}

func Prune(sels []*Selection) []*Selection {
	newSels := make([]*Selection, 0, len(sels))
	merged := make([]bool, len(sels))

	for i, s := range sels {
		if merged[i] {
			continue
		}

		newS := s
		for j := i + 1; j < len(sels); j++ {
			if merged[j] {
				continue
			}
			s2 := sels[j]
			if newS.Overlaps(s2) {
				newS = newS.Merge(s2)
				merged[j] = true
			}
		}

		newSels = append(newSels, newS)
	}

	return newSels
}
