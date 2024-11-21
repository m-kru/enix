package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func FromCursorsDown(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorDown(c))
	}

	sels = Prune(sels)

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

	sels = Prune(sels)

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

func FromCursorsLine(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorLine(c))
	}

	sels = Prune(sels)

	return sels
}

func fromCursorLine(c *cursor.Cursor) *Selection {
	lineRC := c.Line.RuneCount()
	c.RuneIdx = lineRC
	return &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   lineRC,
		Cursor:       c,
	}
}

func FromCursorsRight(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorRight(c))
	}

	sels = Prune(sels)

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

func FromCursorsUp(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorUp(c))
	}

	sels = Prune(sels)

	return sels
}

func fromCursorUp(c *cursor.Cursor) *Selection {
	if c.Line.Prev == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
		}
	}

	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       nil,
	}

	c.Up()
	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       c,
	}

	first.Next = second
	second.Prev = first

	return first
}

func FromCursorsWordEnd(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorWordEnd(c))
	}

	sels = Prune(sels)

	return sels
}

func fromCursorWordEnd(c *cursor.Cursor) *Selection {
	s := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
	}

	c.WordEnd()

	if c.Line == s.Line {
		s.EndRuneIdx = c.RuneIdx
		s.Cursor = c
		return s
	}

	first := s

	s.EndRuneIdx = s.Line.RuneCount()

	line := s.Line
	lineNum := s.LineNum

	for i := 0; i < c.LineNum-first.LineNum-1; i++ {
		line = line.Next
		lineNum++

		nextS := &Selection{
			Line:         line,
			LineNum:      lineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   line.RuneCount(),
		}

		s.Next = nextS
		nextS.Prev = s
		s = nextS
	}

	last := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
	}

	s.Next = last
	last.Prev = s

	return first
}

func FromCursorsPrevWordStart(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorPrevWordStart(c))
	}

	sels = Prune(sels)

	return sels
}

func fromCursorPrevWordStart(c *cursor.Cursor) *Selection {
	s := &Selection{
		Line:       c.Line,
		LineNum:    c.LineNum,
		EndRuneIdx: c.RuneIdx,
	}

	c.PrevWordStart()

	if c.Line == s.Line {
		s.StartRuneIdx = c.RuneIdx
		s.Cursor = c
		return s
	}

	s.StartRuneIdx = 0
	last := s

	line := s.Line
	lineNum := s.LineNum
	for i := 0; i < last.LineNum-c.LineNum-1; i++ {
		line = line.Prev
		lineNum--

		prevS := &Selection{
			Line:         line,
			LineNum:      lineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   line.RuneCount(),
			Next:         s,
		}

		s.Prev = prevS
		s = prevS
	}

	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       c,
		Next:         s,
	}

	return first
}
