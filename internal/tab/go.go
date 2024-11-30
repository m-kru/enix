package tab

import (
	"fmt"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/mark"
)

func (tab *Tab) Go(lineNum int, col int) {
	if lineNum == 0 {
		lineNum = 1
	} else if lineNum < 0 {
		lineNum = tab.LineCount + lineNum + 1
		if lineNum < 1 {
			lineNum = 1
		}
	} else if lineNum > tab.LineCount {
		lineNum = tab.LineCount
	}

	line := tab.Lines.Get(lineNum)
	lineCols := line.Columns()
	if col == 0 {
		col = 1
	} else if col < 0 {
		col = lineCols + col + 2
		if col < 1 {
			col = 1
		}
	} else if col > line.Columns()+1 {
		col = line.Columns() + 1
	}

	rIdx, _, _ := line.RuneIdx(col)

	tab.Cursors = []*cursor.Cursor{
		&cursor.Cursor{
			Line:    line,
			LineNum: lineNum,
			RuneIdx: rIdx,
			ColIdx:  col,
		},
	}
	tab.Selections = nil
}

func (tab *Tab) GoMark(name string) error {
	m, ok := tab.Marks[name]
	if !ok {
		return fmt.Errorf("go: no '%s' mark", name)
	}

	switch m := m.(type) {
	case *mark.CursorMark:
		tab.Cursors = cursor.Clone(m.Cursors)
	default:
		// Going to selection mark unimplemented
	}

	return nil
}
