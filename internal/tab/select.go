package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) SelAll() {
	tab.Cursors = nil
	tab.Selections = []*sel.Selection{sel.SelToTabEnd(tab.Lines, 1, 0)}
}

func (tab *Tab) SelDown() {
	if len(tab.Cursors) > 0 {
		tab.selDownCursors()
	} else {
		tab.selDownSelections()
	}
}

func (tab *Tab) selDownCursors() {
	tab.Selections = sel.FromCursorsDown(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selDownSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Down()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelLeft() {
	if len(tab.Cursors) > 0 {
		tab.selLeftCursors()
	} else {
		tab.selLeftSelections()
	}
}

func (tab *Tab) selLeftCursors() {
	tab.Selections = sel.FromCursorsLeft(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selLeftSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Left()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelLine() {
	if len(tab.Cursors) > 0 {
		tab.selLineCursors()
	} else {
		tab.selLineSelections()
	}
}

func (tab *Tab) selLineCursors() {
	tab.Selections = sel.FromCursorsLine(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selLineSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.NextLine()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelLineEnd() {
	if len(tab.Cursors) > 0 {
		tab.selLineEndCursors()
	} else {
		tab.selLineEndSelections()
	}
}

func (tab *Tab) selLineEndCursors() {
	tab.Selections = sel.FromCursorsLineEnd(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selLineEndSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.LineEnd()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelLineStart() {
	if len(tab.Cursors) > 0 {
		tab.selLineStartCursors()
	} else {
		tab.selLineStartSelections()
	}
}

func (tab *Tab) selLineStartCursors() {
	tab.Selections = sel.FromCursorsLineStart(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selLineStartSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.LineStart()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelPrevWordStart() {
	if len(tab.Cursors) > 0 {
		tab.selPrevWordStartCursors()
	} else {
		tab.selPrevWordStartSelections()
	}
}

func (tab *Tab) selPrevWordStartCursors() {
	tab.Selections = sel.FromCursorsPrevWordStart(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selPrevWordStartSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.PrevWordStart()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelRight() {
	if len(tab.Cursors) > 0 {
		tab.selRightCursors()
	} else {
		tab.selRightSelections()
	}
}

func (tab *Tab) selRightCursors() {
	tab.Selections = sel.FromCursorsRight(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selRightSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Right()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelUp() {
	if len(tab.Cursors) > 0 {
		tab.selUpCursors()
	} else {
		tab.selUpSelections()
	}
}

func (tab *Tab) selUpCursors() {
	tab.Selections = sel.FromCursorsUp(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selUpSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.Up()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelTabEnd() {
	if len(tab.Cursors) > 0 {
		tab.selTabEndCursors()
	} else {
		tab.selTabEndSelections()
	}
}

func (tab *Tab) selTabEndCursors() {
	tab.Selections = sel.FromCursorsTabEnd(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selTabEndSelections() {
	// Leave only the most upper selection
	s := tab.Selections[0]
	for _, s2 := range tab.Selections[1:] {
		if s2.LineNum < s.LineNum || (s2.LineNum == s.LineNum && s2.StartRuneIdx < s.StartRuneIdx) {
			s = s2
		}
	}

	s.TabEnd()

	tab.Selections = []*sel.Selection{s}
}

func (tab *Tab) SelWord() {
	if len(tab.Cursors) > 0 {
		tab.selWordCursors()
	} else {
		tab.selWordSelections()
	}
}

func (tab *Tab) selWordCursors() {
	sels := sel.FromCursorsWord(tab.Cursors)
	if len(sels) > 0 {
		tab.Selections = sels
		tab.Cursors = nil
	}
}

func (tab *Tab) selWordSelections() {
	for i, s1 := range tab.Selections {
		c := s1.GetCursor()
		s2 := sel.FromCursorWord(c)
		if s2 == nil {
			continue
		}
		if s1.Line == s2.Line && s1.StartRuneIdx == s2.StartRuneIdx && s1.EndRuneIdx == s2.EndRuneIdx {
			continue
		}
		tab.Selections[i] = s2
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelWordEnd() {
	if len(tab.Cursors) > 0 {
		tab.selWordEndCursors()
	} else {
		tab.selWordEndSelections()
	}
}

func (tab *Tab) selWordEndCursors() {
	tab.Selections = sel.FromCursorsWordEnd(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selWordEndSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.WordEnd()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelWordStart() {
	if len(tab.Cursors) > 0 {
		tab.selWordStartCursors()
	} else {
		tab.selWordStartSelections()
	}
}

func (tab *Tab) selWordStartCursors() {
	tab.Selections = sel.FromCursorsWordStart(tab.Cursors)
	tab.Cursors = nil
}

func (tab *Tab) selWordStartSelections() {
	for i, s := range tab.Selections {
		tab.Selections[i] = s.WordStart()
	}
	tab.Selections = sel.Prune(tab.Selections)
}

func (tab *Tab) SelToTab(path string) *Tab {
	if len(tab.Selections) == 0 {
		return nil
	}

	str := sel.ToString(tab.Selections)

	return FromString(tab.Frame, str, path)
}

func (tab *Tab) SelSwitchCursor() {
	if len(tab.Cursors) > 0 {
		return
	}

	for _, s := range tab.Selections {
		s.SwitchCursor()
	}
}

func (tab *Tab) SelBracket() {
	var curs []*cursor.Cursor
	if len(tab.Cursors) > 0 {
		curs = tab.Cursors
	} else {
		curs = make([]*cursor.Cursor, 0, len(tab.Selections))
		for _, s := range tab.Selections {
			c := s.GetCursor()
			curs = append(curs, c)
		}
	}

	sels := make([]*sel.Selection, 0, len(curs))

	for _, cur := range curs {
		c1 := cur.MatchBracket(0, 0)
		if c1 == nil {
			continue
		}

		c2 := c1.MatchBracket(0, 0)
		if c2 == nil {
			continue
		}

		if c2.LineNum < c1.LineNum ||
			(c2.LineNum == c1.LineNum && c2.RuneIdx < c1.RuneIdx) {
			tmp := c1
			c1 = c2
			c2 = tmp
		}

		c1.Right()
		if c1.LineNum == c2.LineNum && c1.RuneIdx == c2.RuneIdx {
			continue
		}
		c2.Left()

		s := sel.FromTo(c1, c2)
		sels = append(sels, s)

	}

	if len(sels) > 0 {
		tab.Cursors = nil
		tab.Selections = sel.Prune(sels)
	}
}

func (tab *Tab) SelCurly() {
	var curs []*cursor.Cursor
	if len(tab.Cursors) > 0 {
		curs = tab.Cursors
	} else {
		curs = make([]*cursor.Cursor, 0, len(tab.Selections))
		for _, s := range tab.Selections {
			c := s.GetCursor()
			curs = append(curs, c)
		}
	}

	sels := make([]*sel.Selection, 0, len(curs))

	for _, cur := range curs {
		c1 := cur.MatchCurly(0, 0)
		if c1 == nil {
			continue
		}

		c2 := c1.MatchCurly(0, 0)
		if c2 == nil {
			continue
		}

		if c2.LineNum < c1.LineNum ||
			(c2.LineNum == c1.LineNum && c2.RuneIdx < c1.RuneIdx) {
			tmp := c1
			c1 = c2
			c2 = tmp
		}

		c1.Right()
		if c1.LineNum == c2.LineNum && c1.RuneIdx == c2.RuneIdx {
			continue
		}
		c2.Left()

		s := sel.FromTo(c1, c2)
		sels = append(sels, s)

	}

	if len(sels) > 0 {
		tab.Cursors = nil
		tab.Selections = sel.Prune(sels)
	}
}

func (tab *Tab) SelParen() {
	var curs []*cursor.Cursor
	if len(tab.Cursors) > 0 {
		curs = tab.Cursors
	} else {
		curs = make([]*cursor.Cursor, 0, len(tab.Selections))
		for _, s := range tab.Selections {
			c := s.GetCursor()
			curs = append(curs, c)
		}
	}

	sels := make([]*sel.Selection, 0, len(curs))

	for _, cur := range curs {
		c1 := cur.MatchParen(0, 0)
		if c1 == nil {
			continue
		}

		c2 := c1.MatchParen(0, 0)
		if c2 == nil {
			continue
		}

		if c2.LineNum < c1.LineNum ||
			(c2.LineNum == c1.LineNum && c2.RuneIdx < c1.RuneIdx) {
			tmp := c1
			c1 = c2
			c2 = tmp
		}

		c1.Right()
		if c1.LineNum == c2.LineNum && c1.RuneIdx == c2.RuneIdx {
			continue
		}
		c2.Left()

		s := sel.FromTo(c1, c2)
		sels = append(sels, s)

	}

	if len(sels) > 0 {
		tab.Cursors = nil
		tab.Selections = sel.Prune(sels)
	}
}
