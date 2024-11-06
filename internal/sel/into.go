package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func IntoCursor(s *Selection) *cursor.Cursor {
	if !s.CursorOnLeft() {
		s = s.Last()
	}

	return s.Cursor
}

func IntoCursors(sels []*Selection) []*cursor.Cursor {
	curs := make([]*cursor.Cursor, 0, len(sels))

	for _, s := range sels {
		curs = append(curs, IntoCursor(s))
	}

	return curs
}
