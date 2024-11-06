package cursor

import (
	"github.com/m-kru/enix/internal/util"
	"unicode"
)

func (c *Cursor) Down() {
	if c.Line.Next == nil {
		return
	}

	rIdx := c.RuneIdx
	nextLen := c.Line.Next.RuneCount()
	if c.ColIdx != rIdx {
		if c.ColIdx <= nextLen {
			rIdx = c.ColIdx
		} else {
			rIdx = nextLen
		}
	} else if rIdx > nextLen {
		rIdx = nextLen
	}

	c.RuneIdx = rIdx
	c.Line = c.Line.Next
	c.LineNum++
}

func (c *Cursor) Left() {
	if c.RuneIdx == 0 {
		if c.Line.Prev != nil {
			c.Line = c.Line.Prev
			c.LineNum--
			c.RuneIdx = c.Line.RuneCount()
		}
	} else {
		c.RuneIdx--
	}
	c.ColIdx = c.RuneIdx
}

func (c *Cursor) Right() {
	if c.RuneIdx == c.Line.RuneCount() {
		if c.Line.Next != nil {
			c.Line = c.Line.Next
			c.LineNum++
			c.RuneIdx = 0
		}
	} else {
		c.RuneIdx++
	}
	c.ColIdx = c.RuneIdx
}

func (c *Cursor) Up() {
	if c.Line.Prev == nil {
		return
	}

	rIdx := c.RuneIdx
	prevLen := c.Line.Prev.RuneCount()
	if c.ColIdx != rIdx {
		if c.ColIdx <= prevLen {
			rIdx = c.ColIdx
		} else {
			rIdx = prevLen
		}
	} else if rIdx > prevLen {
		rIdx = prevLen
	}

	c.RuneIdx = rIdx
	c.Line = c.Line.Prev
	c.LineNum--
}

func (c *Cursor) PrevWordStart() {
	if idx, ok := util.PrevWordStart([]rune(c.Line.String()), c.RuneIdx); ok {
		c.RuneIdx = idx
		return
	}

	line := c.Line.Prev
	lineNum := c.LineNum
	for {
		if line == nil {
			return
		}
		lineNum--

		if idx, ok := util.PrevWordStart([]rune(line.String()), line.RuneCount()); ok {
			c.Line = line
			c.LineNum = lineNum
			c.RuneIdx = idx
			c.ColIdx = c.RuneIdx
			return
		}

		line = line.Prev
	}
}

func (c *Cursor) WordEnd() {
	if idx, ok := util.WordEnd([]rune(c.Line.String()), c.RuneIdx); ok {
		c.RuneIdx = idx + 1 // + 1 as we have found word end index.
		c.ColIdx = c.RuneIdx
		return
	}

	line := c.Line.Next
	lineNum := c.LineNum
	for {
		if line == nil {
			return
		}
		lineNum++

		if idx, ok := util.WordEnd([]rune(line.String()), 0); ok {
			c.Line = line
			c.LineNum = lineNum
			c.RuneIdx = idx + 1
			c.ColIdx = c.RuneIdx
			return
		}

		line = line.Next
	}
}

func (c *Cursor) WordStart() {
	if idx, ok := util.WordStart([]rune(c.Line.String()), c.RuneIdx); ok {
		c.RuneIdx = idx
		c.ColIdx = c.RuneIdx
		return
	}

	line := c.Line.Next
	lineNum := c.LineNum
	for {
		if line == nil {
			return
		}
		lineNum++

		if idx, ok := util.WordStart([]rune(line.String()), -1); ok {
			c.Line = line
			c.LineNum = lineNum
			c.RuneIdx = idx
			c.ColIdx = c.RuneIdx
			return
		}

		line = line.Next
	}
}

func (c *Cursor) LineStart() {
	for i, r := range []rune(c.Line.String()) {
		if !unicode.IsSpace(r) {
			if c.RuneIdx == i {
				c.RuneIdx = 0
			} else {
				c.RuneIdx = i
			}
			c.ColIdx = c.RuneIdx
			return
		}
	}
	c.RuneIdx = 0
	c.ColIdx = c.RuneIdx
}

func (c *Cursor) LineEnd() {
	for i := c.Line.RuneCount() - 1; i > 0; i-- {
		r := c.Line.Rune(i)
		if !unicode.IsSpace(r) {
			if c.RuneIdx == i {
				c.RuneIdx = c.Line.RuneCount() - 1
			} else {
				c.RuneIdx = i
			}
			c.ColIdx = c.RuneIdx
			return
		}
	}
	c.RuneIdx = c.Line.RuneCount()
	c.ColIdx = c.RuneIdx
}
