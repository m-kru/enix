package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"
)

type Tab struct {
	Colors *cfg.Colorscheme

	Name       string
	Path       string // File path
	Newline    string // Newline encoding
	FileType   string
	HasChanges bool

	Cursors *cursor.Cursor // First cursor

	Lines *line.Line // First line pointer

	View view.View
}

func (t *Tab) LineCount() int { return t.Lines.Count() }

func (t *Tab) Save() error {
	panic("unimplemented")
}
