package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func IntoCursor(s *Selection) *cursor.Cursor {
	if !s.CursorOnLeft() {
		s = s.Last()
	}

	return &cursor.Cursor{
		Line:    s.Line,
		LineNum: s.LineNum,
		Idx:     s.CursorIdx,
		RuneIdx: s.CursorIdx,
	}
}

func IntoCursors(sels []*Selection) []*cursor.Cursor {
	curs := make([]*cursor.Cursor, 0, len(sels))

	for _, s := range sels {
		curs = append(curs, IntoCursor(s))
	}

	return curs
}
