package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func fromCursorRight(c *cursor.Cursor) *Selection {
	if c.RuneIdx < c.Line.RuneCount() {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx + 1,
			CursorIdx:    c.RuneIdx + 1,
		}
	}

	panic("unimplemented")
}

func FromCursorsRight(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorRight(c))
	}

	return sels
}
