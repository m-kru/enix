package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/lang"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mark"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

type Tab struct {
	Config *cfg.Config
	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Path     string // File path
	Newline  string // Newline encoding
	FileType string

	HasFocus   bool
	HasChanges bool
	State      string // Valid states: "" - normal mode, "insert", "replace".
	RepCount   int    // Command repetition count in normal mode

	Cursors []*cursor.Cursor

	Lines *line.Line // First line

	Marks map[string]mark.Mark

	View view.View

	Highlighter *lang.Highlighter

	Prev *Tab
	Next *Tab
}

func (tab *Tab) LineCount() int { return tab.Lines.Count() }

func (tab *Tab) LineNumWidth() int {
	return util.IntWidth(tab.LineCount())
}

func (tab *Tab) Count() int {
	cnt := 1
	for {
		if tab.Next == nil {
			break
		}
		tab = tab.Next
		cnt++
	}
	return cnt
}

func (tab *Tab) First() *Tab {
	for {
		if tab.Prev == nil {
			break
		}
		tab = tab.Prev
	}
	return tab
}

func (tab *Tab) Last() *Tab {
	for {
		if tab.Next == nil {
			break
		}
		tab = tab.Next
	}
	return tab
}

func (tab *Tab) Append(newTab *Tab) {
	last := tab.Last()
	last.Next = newTab
	newTab.Prev = last
}

func (tab *Tab) HasCursorInLine(n int) bool {
	for _, c := range tab.Cursors {
		if c.Line.Num() == n {
			return true
		}
	}

	return false
}

func (tab *Tab) Trim() {
	var trimmedLines []*line.Line

	l := tab.Lines
	for {
		if l == nil {
			break
		}
		if l.TrimRight() > 0 {
			trimmedLines = append(trimmedLines, l)
		}
		l = l.Next
	}

	for _, c := range tab.Cursors {
		for _, tl := range trimmedLines {
			if c.Line == tl {
				if c.BufIdx > tl.Len() {
					c.BufIdx = tl.Len()
				}
			}
		}
	}

	tab.Cursors = cursor.Prune(tab.Cursors)
}

// AddCursor spawns a new cursor in given line and column.
func (tab *Tab) AddCursor(lineNum int, colIdx int) {
	line := tab.Lines.Get(lineNum)
	if line == nil {
		line = tab.Lines.Last()
	}

	runeIdx, _, ok := line.RuneIdx(colIdx, tab.Config.TabWidth)
	if !ok {
		runeIdx = line.Len()
	}

	c := &cursor.Cursor{
		Config: tab.Cursors[0].Config,
		Line:   line,
		Idx:    runeIdx,
		BufIdx: runeIdx,
	}

	tab.Cursors = append(tab.Cursors, c)
	tab.Cursors = cursor.Prune(tab.Cursors)
}

func (tab *Tab) LastColumnIdx() int {
	idx := 1

	l := tab.Lines
	for {
		if l == nil {
			break
		}

		cols := l.Columns(tab.Config.TabWidth)
		if cols > idx {
			idx = cols
		}

		l = l.Next
	}

	return idx
}

func (tab *Tab) RxEventKey(ev *tcell.EventKey) {
	switch tab.State {
	case "insert":
		tab.RxEventKeyInsert(ev)
	case "replace":
		tab.RxEventKeyReplace(ev)
	}
}
