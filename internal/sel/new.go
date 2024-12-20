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
			Prev:         nil,
			Next:         nil,
		}
	}

	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}

	c.Down()
	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         first,
		Next:         nil,
	}

	first.Next = second

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
			Prev:         nil,
			Next:         nil,
		}
	}

	if c.Line.Prev == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   0,
			Cursor:       c,
			Prev:         nil,
			Next:         nil,
		}
	}

	c.Left()
	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         nil,
		Next:         nil,
	}
	second := &Selection{
		Line:         c.Line.Next,
		LineNum:      c.LineNum + 1,
		StartRuneIdx: 0,
		EndRuneIdx:   0,
		Cursor:       nil,
		Prev:         first,
		Next:         nil,
	}

	first.Next = second

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
	// Call LineEnd two times, because there might be spaces at the line end.
	c.LineEnd()
	c.LineEnd()
	return &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         nil,
		Next:         nil,
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
			Prev:         nil,
			Next:         nil,
		}
	}

	if c.Line.Next == nil {
		return &Selection{
			Line:         c.Line,
			LineNum:      c.LineNum,
			StartRuneIdx: c.RuneIdx,
			EndRuneIdx:   c.RuneIdx,
			Cursor:       c,
			Prev:         nil,
			Next:         nil,
		}
	}

	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}
	c.Right()
	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         first,
		Next:         nil,
	}

	first.Next = second

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
			Prev:         nil,
			Next:         nil,
		}
	}

	second := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}

	c.Up()
	first := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: c.RuneIdx,
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       c,
		Prev:         nil,
		Next:         second,
	}

	second.Prev = first

	return first
}

func FromCursorsWord(curs []*cursor.Cursor) []*Selection {
	sels := make([]*Selection, 0, len(curs))

	for _, c := range curs {
		sels = append(sels, fromCursorWord(c))
	}

	sels = Prune(sels)

	return sels
}

func fromCursorWord(c *cursor.Cursor) *Selection {
	wordPos := c.WordPosition()

	switch wordPos {
	case cursor.InSpace:
		c.WordStart()
	case cursor.InWord:
		c.PrevWordStart()
	}

	srIdx := c.RuneIdx
	c.WordEnd()
	c.Left()

	return &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: srIdx,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         nil,
		Next:         nil,
	}
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
		EndRuneIdx:   c.Line.RuneCount(),
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}

	c.WordEnd()

	if c.Line == s.Line {
		s.EndRuneIdx = c.RuneIdx
		s.Cursor = c
		return s
	}

	first := s

	line := s.Line
	lineNum := s.LineNum

	for range c.LineNum - first.LineNum - 1 {
		line = line.Next
		lineNum++

		nextS := &Selection{
			Line:         line,
			LineNum:      lineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   line.RuneCount(),
			Cursor:       nil,
			Prev:         s,
			Next:         nil,
		}

		s.Next = nextS
		s = nextS
	}

	last := &Selection{
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       c,
		Prev:         s,
		Next:         nil,
	}

	s.Next = last

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
		Line:         c.Line,
		LineNum:      c.LineNum,
		StartRuneIdx: 0,
		EndRuneIdx:   c.RuneIdx,
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}

	c.PrevWordStart()

	if c.Line == s.Line {
		s.StartRuneIdx = c.RuneIdx
		s.Cursor = c
		return s
	}

	last := s

	line := s.Line
	lineNum := s.LineNum
	for range last.LineNum - c.LineNum - 1 {
		line = line.Prev
		lineNum--

		prevS := &Selection{
			Line:         line,
			LineNum:      lineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   line.RuneCount(),
			Cursor:       nil,
			Prev:         nil,
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
		Prev:         nil,
		Next:         s,
	}

	return first
}

// FromTo creates a selection spanning from the start cursor to the end cursor.
func FromTo(startC, endC *cursor.Cursor) *Selection {
	cursorOnLeft := true
	if startC.LineNum < endC.LineNum || (startC.LineNum == endC.LineNum && startC.RuneIdx < endC.RuneIdx) {
		cursorOnLeft = false
	}

	if cursorOnLeft {
		tmpC := startC
		startC = endC
		endC = tmpC
	}

	first := &Selection{
		Line:         startC.Line,
		LineNum:      startC.LineNum,
		StartRuneIdx: startC.RuneIdx,
		EndRuneIdx:   0,
		Cursor:       nil,
		Prev:         nil,
		Next:         nil,
	}

	if startC.LineNum == endC.LineNum {
		first.EndRuneIdx = endC.RuneIdx
		c := endC
		if cursorOnLeft {
			c = startC
		}
		first.Cursor = c.Clone()
		return first
	}

	first.EndRuneIdx = startC.Line.RuneCount()

	line := startC.Line.Next
	lineNum := startC.LineNum + 1
	prevS := first
	for {
		s := &Selection{
			Line:         line,
			LineNum:      lineNum,
			StartRuneIdx: 0,
			EndRuneIdx:   0,
			Cursor:       nil,
			Prev:         prevS,
			Next:         nil,
		}

		prevS.Next = s
		prevS = s

		if lineNum == endC.LineNum {
			s.EndRuneIdx = endC.RuneIdx
			break
		} else {
			s.EndRuneIdx = s.Line.RuneCount()
			line = line.Next
			lineNum++
		}
	}

	if cursorOnLeft {
		first.Cursor = startC.Clone()
	} else {
		prevS.Cursor = endC.Clone()
	}

	return first
}
