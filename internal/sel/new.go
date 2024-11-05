package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func FromCursorsDown(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorDown(c))
	}

	// TODO: Prune selections

	return sels
}

func fromCursorDown(c *cursor.Cursor) *Selection {
	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		CursorIdx:    -1,
	}

	if c.Line.Next == nil {
		first.CursorIdx = c.Line.RuneCount()
		return first
	}

	cIdx := c.RuneIdx
	if c.Line.Next.RuneCount() < cIdx {
		cIdx = c.Line.Next.RuneCount()
	}

	second := &Selection{
		Line:         c.Line.Next,
		LineNum:      c.LineNum + 1,
		StartRuneIdx: 0,
		EndRuneIdx:   cIdx,
		CursorIdx:    cIdx,
	}

	first.Next = second
	second.Prev = first

	return first
}

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

	if c.Line.Prev == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   0,
			CursorIdx:    0,
		}
	}

	first := &Selection{
		Line:         c.Line.Prev,
		LineNum:      c.LineNum - 1,
		StartRuneIdx: c.Line.Prev.RuneCount(),
		EndRuneIdx:   c.Line.Prev.RuneCount(),
		CursorIdx:    c.Line.Prev.RuneCount(),
	}
	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   0,
		CursorIdx:    -1,
	}

	first.Next = second
	second.Prev = first

	return first
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

	if c.Line.Next == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx,
			CursorIdx:    c.RuneIdx,
		}
	}

	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.RuneIdx,
		CursorIdx:    -1,
	}
	second := &Selection{
		Line:         c.Line.Next,
		LineNum:      c.LineNum + 1,
		StartRuneIdx: 0,
		EndRuneIdx:   0,
		CursorIdx:    0,
	}

	first.Next = second
	second.Prev = first

	return first
}
