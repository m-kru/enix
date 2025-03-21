package cursor

import (
	"github.com/m-kru/enix/internal/util"
	"unicode"
)

func (c *Cursor) Down() {
	if c.Line.Next == nil {
		return
	}

	c.Line = c.Line.Next
	c.LineNum++
	rIdx, _, ok := c.Line.RuneIdx(c.colIdx)
	if ok {
		c.RuneIdx = rIdx
		return
	}

	c.RuneIdx = c.Line.RuneCount()
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
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
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
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}

func (c *Cursor) Up() {
	if c.Line.Prev == nil {
		return
	}

	c.Line = c.Line.Prev
	c.LineNum--
	rIdx, _, ok := c.Line.RuneIdx(c.colIdx)
	if ok {
		c.RuneIdx = rIdx
		return
	}

	c.RuneIdx = c.Line.RuneCount()
}

func (c *Cursor) PrevWordStart() {
	if idx, ok := util.PrevWordStart([]rune(c.Line.String()), c.RuneIdx); ok {
		c.RuneIdx = idx
		c.colIdx = c.Line.ColumnIdx(idx)
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
			c.colIdx = c.Line.ColumnIdx(idx)
			return
		}

		line = line.Prev
	}
}

func (c *Cursor) WordEnd() {
	if idx, ok := util.WordEnd([]rune(c.Line.String()), c.RuneIdx); ok {
		c.RuneIdx = idx
		c.colIdx = c.Line.ColumnIdx(idx)
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
			c.RuneIdx = idx
			c.colIdx = c.Line.ColumnIdx(idx)
			return
		}

		line = line.Next
	}
}

func (c *Cursor) WordStart() bool {
	if idx, ok := util.WordStart([]rune(c.Line.String()), c.RuneIdx); ok {
		c.RuneIdx = idx
		c.colIdx = c.Line.ColumnIdx(idx)
		return true
	}

	line := c.Line.Next
	lineNum := c.LineNum
	for {
		if line == nil {
			return false
		}
		lineNum++

		if idx, ok := util.WordStart([]rune(line.String()), -1); ok {
			c.Line = line
			c.LineNum = lineNum
			c.RuneIdx = idx
			c.colIdx = c.Line.ColumnIdx(idx)
			return true
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
			c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
			return
		}
	}
	c.RuneIdx = 0
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}

func (c *Cursor) LineEnd() {
	for i := c.Line.RuneCount() - 1; i > 0; i-- {
		r := c.Line.Rune(i)
		if !unicode.IsSpace(r) {
			if c.RuneIdx == i+1 {
				break
			} else {
				c.RuneIdx = i + 1
				c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
				return
			}
		}
	}

	c.RuneIdx = c.Line.RuneCount()
	c.colIdx = c.Line.ColumnIdx(c.RuneIdx)
}
