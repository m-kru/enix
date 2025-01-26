package tab

import (
	"github.com/m-kru/enix/internal/cursor"
)

func (tab *Tab) ViewCenter() {
	var cur *cursor.Cursor
	if len(tab.Cursors) > 0 {
		cur = tab.Cursors[len(tab.Cursors)-1]
	} else {
		cur = tab.LastSel().GetCursor()
	}

	lineNum := cur.LineNum - tab.Frame.Height/2
	if lineNum < 1 {
		lineNum = 1
	}
	tab.View.Line = lineNum

	// TODO: Should column be adjusted here as well?
}

func (tab *Tab) ViewDown() {
	if tab.View.LastLine() >= tab.Lines.Count()+tab.Frame.Height/2 {
		return
	}
	tab.View = tab.View.Down()
}

func (tab *Tab) ViewDownHalf() {
	for range tab.Frame.Height / 2 {
		tab.ViewDown()
	}
}

func (tab *Tab) ViewUp() {
	tab.View = tab.View.Up()
}

func (tab *Tab) ViewUpHalf() {
	for range tab.Frame.Height / 2 {
		tab.ViewUp()
	}
}

func (tab *Tab) ViewRight() {
	lastCol := tab.View.LastColumn() + tab.LineNumWidth() - tab.Frame.Width/2
	if lastCol >= tab.LastColumnIdx() {
		return
	}

	tab.View = tab.View.Right()
}

func (tab *Tab) ViewLeft() {
	tab.View = tab.View.Left()
}
