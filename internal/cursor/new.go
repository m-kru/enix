package cursor

import (
	"github.com/m-kru/enix/internal/line"
)

func New(line *line.Line, lineNum int, runeIdx int) *Cursor {
	return &Cursor{
		Line:    line,
		LineNum: lineNum,
		RuneIdx: runeIdx,
		colIdx:  line.ColumnIdx(runeIdx),
	}
}
