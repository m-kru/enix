package tab

import (
	"unicode"

	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

// PrimaryClick handles mouse primary button click.
// x and y are tab frame coordinates.
func (tab *Tab) PrimaryClick(x, y int) {
	x -= tab.LineNumWidth() + 1
	if x < 0 {
		x = 0
	}

	line := tab.Lines.Get(tab.View.Line + y)
	if line == nil {
		line = tab.Lines.Last()
	}

	rIdx, _, ok := line.RuneIdx(tab.View.Column + x)
	if !ok {
		rIdx = line.RuneCount()
	}

	c := cursor.New(line, line.Num(), rIdx)

	tab.Cursors = make([]*cursor.Cursor, 1, 16)
	tab.Cursors[0] = c
}

// PrimaryClickCtrl handles mouse primary button click with Ctrl modifier.
// x and y are tab frame coordinates.
func (tab *Tab) PrimaryClickCtrl(x, y int) {
	x -= tab.LineNumWidth() + 1
	if x < 0 {
		x = 0
	}

	line := tab.Lines.Get(tab.View.Line + y)
	if line == nil {
		line = tab.Lines.Last()
	}

	idx, _, ok := line.RuneIdx(tab.View.Column + x)
	if !ok {
		idx = line.RuneCount()
	}

	idx, ok = tab.AddCursor(line.Num(), line.ColumnIdx(idx))
	if ok || len(tab.Cursors) == 1 {
		return
	}

	tab.Cursors = append(tab.Cursors[0:idx], tab.Cursors[idx+1:]...)
}

// DoublePrimaryClick handles mouse double primary button click.
// x and y are tab frame coordinates.
func (tab *Tab) DoublePrimaryClick(x, y int) {
	x -= tab.LineNumWidth() + 1
	if x < 0 {
		x = 0
	}

	line := tab.Lines.Get(tab.View.Line + y)
	if line == nil {
		return
	}

	rIdx, _, ok := line.RuneIdx(tab.View.Column + x)
	if !ok {
		return
	}

	r := line.Rune(rIdx)
	if unicode.IsSpace(r) {
		// Select whole whitespace sequence here
		return
	}

	c := cursor.New(line, line.Num(), rIdx)
	s := sel.FromCursorWord(c)

	if s != nil {
		tab.Cursors = nil
		tab.Selections = []*sel.Selection{s}
	}
}

// PrimaryClickAlt handles mouse double primary button click with Alt modifier.
// x and y are tab frame coordinates.
func (tab *Tab) PrimaryClickAlt(x, y int) {
	x -= tab.LineNumWidth() + 1
	if x < 0 {
		x = 0
	}

	line := tab.Lines.Get(tab.View.Line + y)
	if line == nil {
		return
	}

	rIdx, _, ok := line.RuneIdx(tab.View.Column + x)
	if !ok {
		return
	}

	r := line.Rune(rIdx)
	if unicode.IsSpace(r) {
		// Select whole whitespace sequence here
		return
	}

	c := cursor.New(line, line.Num(), rIdx)
	s := sel.FromCursorWord(c)

	if s == nil {
		return
	}

	if len(tab.Cursors) > 0 {
		tab.Cursors = nil
		tab.Selections = []*sel.Selection{s}
	} else {
		tab.Selections = append(tab.Selections, s)
		tab.Selections = sel.Prune(tab.Selections)
	}
}
