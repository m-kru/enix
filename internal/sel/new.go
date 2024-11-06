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
	if c.Line.Next == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
		}
	}

	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       nil,
	}

	c.Down()
	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
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
		c.Left()
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx + 1,
			Cursor:       c,
		}
	}

	if c.Line.Prev == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   0,
			Cursor:       c,
		}
	}

	c.Left()
	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
	}
	second := &Selection{
		Line:         c.Line.Next,
		LineNum:      c.LineNum + 1,
		StartRuneIdx: 0,
		EndRuneIdx:   0,
		Cursor:       nil,
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
		c.Right()
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx - 1,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
		}
	}

	if c.Line.Next == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
		}
	}

	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       nil,
	}
	c.Right()
	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
	}

	first.Next = second
	second.Prev = first

	return first
}
