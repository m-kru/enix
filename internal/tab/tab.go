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

	Name       string
	Path       string // File path
	Newline    string // Newline encoding
	FileType   string
	HasChanges bool

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
		if c.LineNum() == n {
			return true
		}
		c = c.Next
	}
	return false
}

func (t *Tab) Save() error {
	panic("unimplemented")
}
