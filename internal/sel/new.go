package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func FromCursorsLeft(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorLeft(c))
	}

	// TODO: Prune selections

	return sels
}

func fromCursorLeft(c *cursor.Cursor) *Selection {
	if c.RuneIdx > 0 {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx - 1,
			EndRuneIdx:   c.RuneIdx,
			CursorIdx:    c.RuneIdx - 1,
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

	// TODO: Prune selections

	panic("unimplemented")
}
