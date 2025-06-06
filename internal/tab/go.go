package tab

import (
	"fmt"
	"github.com/m-kru/enix/internal/arg"
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

	cur := cursor.New(line, lineNum, rIdx)
	tab.Cursors = []*cursor.Cursor{cur}
	tab.Selections = nil

	// Don't change the view if running in running in script mode.
	if arg.Script != "" {
		return
	}

	if tab.View.IsVisible(cur.View()) {
		return
	}
	tab.ViewCenter()
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
		return nil
	}

	// Don't change the view if running in running in script mode.
	if arg.Script != "" {
		return nil
	}

	var cur *cursor.Cursor
	if len(tab.Cursors) > 0 {
		cur = tab.Cursors[len(tab.Cursors)-1]
	} else {
		cur = tab.LastSel().GetCursor()
	}

	if tab.View.IsVisible(cur.View()) {
		return nil
	}
	tab.ViewCenter()

	return nil
}
