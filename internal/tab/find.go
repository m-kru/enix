package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/find"
	"github.com/m-kru/enix/internal/search"
	"github.com/m-kru/enix/internal/sel"
)

func getNextFind(finds []find.Find, c *cursor.Cursor) find.Find {
	firstFind := finds[0]
	lastFind := finds[len(finds)-1]

	// First check wrap-around
	if c.LineNum > lastFind.LineNum ||
		(c.LineNum == lastFind.LineNum && c.RuneIdx >= lastFind.StartRuneIdx) ||
		(c.LineNum == firstFind.LineNum && c.RuneIdx < firstFind.StartRuneIdx) {
		return firstFind
	}

	// TODO: Not optimal, O(n) complexity, can be implemented as O(log(n)).
	for i := 0; i < len(finds)-1; i++ {
		prevF := finds[i]
		nextF := finds[i+1]
		// Can below check be simplified?
		if (prevF.LineNum <= c.LineNum && c.LineNum < nextF.LineNum) ||
			(prevF.LineNum == c.LineNum && prevF.StartRuneIdx <= c.RuneIdx && c.RuneIdx < nextF.StartRuneIdx) ||
			(c.LineNum == nextF.LineNum && c.RuneIdx < nextF.StartRuneIdx) {
			return nextF
		}
	}

	return firstFind
}

func (tab *Tab) FindNext() {
	if tab.SearchCtx.Regexp == nil {
		if tab.SearchCtx.PrevRegexp == nil {
			return
		}
		tab.SearchCtx.Regexp = tab.SearchCtx.PrevRegexp
		search.Search(tab.Lines, tab.View.Line, &tab.SearchCtx)
	}

	finds := tab.SearchCtx.Finds
	if len(finds) == 0 {
		return
	}

	var c *cursor.Cursor
	if len(tab.Cursors) > 0 {
		c = tab.Cursors[len(tab.Cursors)-1]
	} else {
		c = tab.Selections[len(tab.Selections)-1].GetCursor()
	}

	f := getNextFind(finds, c)

	line := tab.Lines.Get(f.LineNum)
	s := &sel.Selection{
		Line:         line,
		LineNum:      f.LineNum,
		StartRuneIdx: f.StartRuneIdx,
		EndRuneIdx:   f.EndRuneIdx - 1,
		Cursor: &cursor.Cursor{
			Line:    line,
			LineNum: f.LineNum,
			ColIdx:  line.ColumnIdx(f.StartRuneIdx),
			RuneIdx: f.StartRuneIdx,
		},
	}

	if len(tab.Cursors) > 0 {
		tab.Cursors = nil
	}
	tab.Selections = []*sel.Selection{s}
}

func (tab *Tab) FindSelNext() {
	if len(tab.Cursors) > 0 {
		tab.FindNext()
		return
	}

	if tab.SearchCtx.Regexp == nil {
		if tab.SearchCtx.PrevRegexp == nil {
			return
		}
		tab.SearchCtx.Regexp = tab.SearchCtx.PrevRegexp
		search.Search(tab.Lines, tab.View.Line, &tab.SearchCtx)
	}

	finds := tab.SearchCtx.Finds
	if len(finds) == 0 {
		return
	}

	c := tab.Selections[len(tab.Selections)-1].GetCursor()

	f := getNextFind(finds, c)

	line := tab.Lines.Get(f.LineNum)
	s := &sel.Selection{
		Line:         line,
		LineNum:      f.LineNum,
		StartRuneIdx: f.StartRuneIdx,
		EndRuneIdx:   f.EndRuneIdx - 1,
		Cursor: &cursor.Cursor{
			Line:    line,
			LineNum: f.LineNum,
			ColIdx:  line.ColumnIdx(f.StartRuneIdx),
			RuneIdx: f.StartRuneIdx,
		},
	}

	tab.Selections = append(tab.Selections, s)
	tab.Selections = sel.Prune(tab.Selections)
}
