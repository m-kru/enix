package tab

import (
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Esc() {
	if tab.RepCount != 0 {
		tab.RepCount = 0
		return
	}

	if len(tab.Selections) > 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
		return
	}

	if len(tab.Cursors) > 1 {
		tab.Cursors = tab.Cursors[len(tab.Cursors)-1:]
		return
	}

	if tab.SearchCtx.Regexp != nil {
		tab.SearchCtx.PrevRegexp = tab.SearchCtx.Regexp
		tab.SearchCtx.Regexp = nil
		return
	}
}
