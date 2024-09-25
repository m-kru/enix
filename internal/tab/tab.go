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

func (t *Tab) LineCount() int { return t.Lines.Count() }

func (t *Tab) Count() int {
	cnt := 1
	for {
		if t.Next == nil {
			break
		}
		t = t.Next
		cnt++
	}
	return cnt
}

func (t *Tab) HasCursorInLine(n int) bool {
	c := t.Cursors
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

func (t *Tab) Trim() {
	var trimmedLines []*line.Line

	l := t.Lines
	for {
		if l == nil {
			break
		}
		if l.Trim() > 0 {
			trimmedLines = append(trimmedLines, l)
		}
		l = l.Next
	}

	c := t.Cursors
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

	t.Cursors.Prune()
}

func (t *Tab) AddCursor(line int, col int) {
	l := t.Lines.Get(line)
	if l == nil {
		l = t.Lines.Last()
	}

	idx, _, ok := l.RuneIdx(col, t.Config.TabWidth)
	if !ok {
		idx = l.Len() - 1
	}

	lastCur := t.Cursors.Last()

	c := cursor.Cursor{
		Config: lastCur.Config,
		Line:   l,
		Idx:    idx,
		BufIdx: idx,
		Prev:   lastCur,
		Next:   nil,
	}

	lastCur.Next = &c

	t.Cursors.Prune()
}
