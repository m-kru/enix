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
		(c.LineNum == firstFind.LineNum && c.RuneIdx <= firstFind.StartRuneIdx) ||
		(c.LineNum == lastFind.LineNum && c.RuneIdx > lastFind.StartRuneIdx) {
		return lastFind
	}

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
	finds := tab.SearchCtx.Finds

	if tab.SearchCtx.Regexp == nil {
		if tab.SearchCtx.PrevRegexp == nil {
			return
		}
		tab.SearchCtx.Regexp = tab.SearchCtx.PrevRegexp
		tab.SearchCtx.FirstVisLineNum = tab.View.Line
		finds = search.Search(tab.Lines, &tab.SearchCtx)
	}

	if len(finds) == 0 {
		return
	}

	var c *cursor.Cursor
	if len(tab.Cursors) > 0 {
		c = tab.Cursors[len(tab.Cursors)-1]
	} else {
		c = tab.LastSel().GetCursor()
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

	finds := tab.SearchCtx.Finds

	if tab.SearchCtx.Regexp == nil {
		if tab.SearchCtx.PrevRegexp == nil {
			return
		}
		tab.SearchCtx.Regexp = tab.SearchCtx.PrevRegexp
		tab.SearchCtx.FirstVisLineNum = tab.View.Line
		finds = search.Search(tab.Lines, &tab.SearchCtx)
	}

	if len(finds) == 0 {
		return
	}

	c := tab.LastSel().GetCursor()

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

// next is true for find-desel-next and false for find-sel-prev.
func (tab *Tab) FindDesel(next bool) {
	if len(tab.Selections) == 0 {
		tab.Find(next)
		return
	}

	finds := tab.SearchCtx.Finds
	cur := tab.Selections[len(tab.Selections)-1].GetCursor()

	// Check if last selection is a find result.
	lastSel := tab.LastSel()
	for _, f := range finds {
		if lastSel.EqualsFind(f) {
			tab.Selections = tab.Selections[0 : len(tab.Selections)-1]
			break
		}
	}

	var f find.Find
	if next {
		f = getNextFind(finds, cur)
	} else {
		f = getPrevFind(finds, cur)
	}

	// If seleciton for a given find already exists, then move it to be the last selection.
	selExists := false
	for i, s := range tab.Selections {
		if s.EqualsFind(f) {
			lastIdx := len(tab.Selections) - 1
			tmp := tab.Selections[lastIdx]
			tab.Selections[lastIdx] = s
			tab.Selections[i] = tmp
			selExists = true
			break
		}
	}

	if !selExists {
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
	}

	tab.Selections = sel.Prune(tab.Selections)

	if !tab.View.IsVisible(tab.LastSel().View()) {
		tab.ViewCenter()
	}
}
