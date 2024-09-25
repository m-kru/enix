package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"
)

type Tab struct {
	Config *cfg.Config
	Colors *cfg.Colorscheme
	Keys   *cfg.Keybindings

	Path     string // File path
	Newline  string // Newline encoding
	FileType string

	HasFocus     bool
	HasChanges   bool
	InInsertMode bool

	Cursors *cursor.Cursor // First cursor

	Lines *line.Line // First line

	View view.View

	Prev *Tab
	Next *Tab
}

func (tab *Tab) LineCount() int { return tab.Lines.Count() }

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

func (tab *Tab) HasCursorInLine(n int) bool {
	c := tab.Cursors
	for {
		if c == nil {
			break
		}
		if c.Line.Num() == n {
			return true
		}
		c = c.Next
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
		if l.Trim() > 0 {
			trimmedLines = append(trimmedLines, l)
		}
		l = l.Next
	}

	c := tab.Cursors
	for {
		if c == nil {
			break
		}

		for _, tl := range trimmedLines {
			if c.Line == tl {
				if c.BufIdx > tl.Len() {
					c.BufIdx = tl.Len()
				}
			}
		}

		c = c.Next
	}

	tab.Cursors.Prune()
}

// AddCursor spawns a new cursor in given line and column.
func (tab *Tab) AddCursor(lineNum int, colIdx int) {
	line := tab.Lines.Get(lineNum)
	if line == nil {
		line = tab.Lines.Last()
	}

	runeIdx, _, ok := line.RuneIdx(colIdx, tab.Config.TabWidth)
	if !ok {
		runeIdx = line.Len() - 1
	}

	lastCur := tab.Cursors.Last()

	c := cursor.Cursor{
		Config: lastCur.Config,
		Line:   line,
		Idx:    runeIdx,
		BufIdx: runeIdx,
		Prev:   lastCur,
		Next:   nil,
	}

	lastCur.Next = &c

	tab.Cursors.Prune()
}
