package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/lang"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mark"
	"github.com/m-kru/enix/internal/search"
	"github.com/m-kru/enix/internal/sel"
	"github.com/m-kru/enix/internal/undo"
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
	State      string // Valid states: "" - normal mode, "insert", "replace", "key-name".
	RepCount   int    // Command repetition count in normal mode

	Lines     *line.Line // First line
	LineCount int

	Cursors    []*cursor.Cursor
	Selections []*sel.Selection

	SearchCtx search.Context

	Marks map[string]mark.Mark

	View view.View

	Highlighter *lang.Highlighter

	UndoStack *undo.Stack
	RedoStack *undo.Stack

	Prev *Tab
	Next *Tab
}

func (tab *Tab) LineNumWidth() int {
	return util.IntWidth(tab.LineCount)
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

func (tab *Tab) Exists(path string) bool {
	t := tab.First()

	for {
		if t == nil {
			break
		}

		if t.Path == path {
			return true
		}

		t = t.Next
	}

	return false
}

func (tab *Tab) Append(newTab *Tab) {
	last := tab.Last()
	last.Next = newTab
	newTab.Prev = last
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
				if c.RuneIdx > tl.RuneCount() {
					c.RuneIdx = tl.RuneCount()
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

	runeIdx, _, ok := line.RuneIdx(colIdx)
	if !ok {
		runeIdx = line.RuneCount()
	}

	c := &cursor.Cursor{
		Line:    line,
		LineNum: lineNum,
		ColIdx:  runeIdx,
		RuneIdx: runeIdx,
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

		cols := l.Columns()
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
	case "key-name":
		tab.RxEventKeyKeyName(ev)
	case "replace":
		tab.RxEventKeyReplace(ev)
	}
}

// GetWord returns word under last cursor or selection.
func (tab *Tab) GetWord() string {
	if len(tab.Cursors) > 0 {
		return tab.Cursors[len(tab.Cursors)-1].GetWord()
	} else {
		return tab.Selections[len(tab.Selections)-1].GetCursor().GetWord()
	}
}
