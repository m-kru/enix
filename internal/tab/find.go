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
	for i := range len(finds) - 1 {
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

func getPrevFind(finds []find.Find, c *cursor.Cursor) find.Find {
	firstFind := finds[0]
	lastFind := finds[len(finds)-1]

	// First check wrap-around
	if c.LineNum < firstFind.LineNum ||
		(c.LineNum == firstFind.LineNum && c.RuneIdx <= lastFind.StartRuneIdx) ||
		(c.LineNum == lastFind.LineNum && c.RuneIdx > lastFind.StartRuneIdx) {
		return lastFind
	}

	// TODO: Not optimal, O(n) complexity, can be implemented as O(log(n)).
	for i := range len(finds) - 1 {
		prev := finds[i]
		next := finds[i+1]
		// Can below check be simplified?
		if (prev.LineNum < c.LineNum && c.LineNum < next.LineNum) ||
			(prev.LineNum == c.LineNum && prev.EndRuneIdx < c.RuneIdx && c.RuneIdx <= next.EndRuneIdx) ||
			(c.LineNum == next.LineNum && c.RuneIdx <= next.EndRuneIdx) {
			return prev
		}
	}

	return firstFind
}

// next is true for find-next and false for find-prev.
func (tab *Tab) Find(next bool) {
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

	var f find.Find
	if next {
		f = getNextFind(finds, c)
	} else {
		f = getPrevFind(finds, c)
	}

	line := tab.Lines.Get(f.LineNum)
	s := &sel.Selection{
		Line:         line,
		LineNum:      f.LineNum,
		StartRuneIdx: f.StartRuneIdx,
		EndRuneIdx:   f.EndRuneIdx - 1,
		Cursor:       cursor.New(line, f.LineNum, f.StartRuneIdx),
		Prev:         nil,
		Next:         nil,
	}

	if len(tab.Cursors) > 0 {
		tab.Cursors = nil
	}
	tab.Selections = []*sel.Selection{s}

	if !tab.View.IsVisible(s.View()) {
		tab.ViewCenter()
	}
}

// next is true for find-sel-next and false for find-sel-prev.
func (tab *Tab) FindSel(next bool) {
	if len(tab.Cursors) > 0 {
		tab.Find(next)
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

	var f find.Find
	if next {
		f = getNextFind(finds, c)
	} else {
		f = getPrevFind(finds, c)
	}

	line := tab.Lines.Get(f.LineNum)
	s := &sel.Selection{
		Line:         line,
		LineNum:      f.LineNum,
		StartRuneIdx: f.StartRuneIdx,
		EndRuneIdx:   f.EndRuneIdx - 1,
		Cursor:       cursor.New(line, f.LineNum, f.StartRuneIdx),
		Prev:         nil,
		Next:         nil,
	}

	tab.Selections = append(tab.Selections, s)
	tab.Selections = sel.Prune(tab.Selections)

	if !tab.View.IsVisible(s.View()) {
		tab.ViewCenter()
	}
}
