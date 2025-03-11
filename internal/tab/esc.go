package tab

import (
	"github.com/m-kru/enix/internal/sel"
)

// Esc returns true if tab view should be updated after escape execution.
func (tab *Tab) Esc() bool {
	if tab.RepCount != 0 {
		tab.RepCount = 0
		return false
	}

	if len(tab.Selections) > 0 {
		tab.Cursors = sel.ToCursors(tab.Selections)
		tab.Selections = nil
		return false
	}

	if len(tab.Cursors) > 1 {
		tab.Cursors = tab.Cursors[len(tab.Cursors)-1:]
		return false
	}

	if tab.SearchCtx.Regexp != nil {
		tab.SearchCtx.PrevRegexp = tab.SearchCtx.Regexp
		tab.SearchCtx.Regexp = nil
		return false
	}

	return true
}
