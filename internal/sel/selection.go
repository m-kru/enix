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

func (s *Selection) RuneCount() int {
	return s.EndRuneIdx - s.StartRuneIdx + 1
}

// Last returns last chunk being part of the selection s.
func (s *Selection) Last() *Selection {
	for s.Next != nil {
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

func (s *Selection) LastColumnNumber() int {
	return s.Line.Columns()
}

func (s *Selection) View() view.View {
	line := s.LineNum

	minCol := s.Line.ColumnIdx(s.StartRuneIdx)
	maxCol := s.Line.ColumnIdx(s.EndRuneIdx)

	for s.Next != nil {
		min := s.Line.ColumnIdx(s.StartRuneIdx)
		if min < minCol {
			minCol = min
		}

		max := s.Line.ColumnIdx(s.EndRuneIdx)
		if max > maxCol {
			maxCol = max
		}

		s = s.Next
	}

	return view.View{
		Line:   line,
		Column: minCol,
		Height: s.LineNum - line + 1,
		Width:  maxCol - minCol + 1,
	}
}

func (s *Selection) Overlaps(s2 *Selection) bool {
	for s != nil {
		subs2 := s2
		for subs2 != nil {
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

// SpawnCursorOnRight creates and returns a cursor on the right side of selection.
func (s *Selection) SpawnCursorOnRight() *cursor.Cursor {
	s = s.Last()
	return cursor.New(s.Line, s.LineNum, s.EndRuneIdx)
}

func (s *Selection) SwitchCursor() {
	last := s.Last()

	if s == last {
		if s.EndRuneIdx == s.Cursor.RuneIdx {
			s.Cursor = cursor.New(s.Line, s.LineNum, s.StartRuneIdx)
		} else {
			s.Cursor = cursor.New(s.Line, s.LineNum, s.EndRuneIdx)
		}
	} else {
		if s.Cursor != nil {
			s.Cursor = nil
			last.Cursor = cursor.New(last.Line, last.LineNum, last.EndRuneIdx)
		} else {
			last.Cursor = nil
			s.Cursor = cursor.New(s.Line, s.LineNum, s.StartRuneIdx)
		}
	}
}

// Lines returns a list of all lines spanned by the selection.
func (s *Selection) Lines() []*line.Line {
	lines := make([]*line.Line, 0, 16)

	for s != nil {
		lines = append(lines, s.Line)

		s = s.Next
	}

	return lines
}

// Lines returns a list of all lines spannded by selections.
// Lines might be placed in an arbitrary order.
// However, each line appears only once in the returned list.
func Lines(sels []*Selection) []*line.Line {
	lines := make([]*line.Line, 0, len(sels))

	for _, s := range sels {
		ls := s.Lines()
		for _, l := range ls {
			found := false

			for _, l2 := range lines {
				if l2 == l {
					found = true
					break
				}
			}

			if found {
				continue
			}
			lines = append(lines, l)
		}
	}

	return lines
}
