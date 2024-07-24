package tab

import (
	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
)

type Tab struct {
	Screen tcell.Screen
	StartX int
	EndX   int
	StartY int
	EndY   int

	Colors *cfg.Colorscheme

	Name       string // Path of the file
	Newline    string // Newline encoding
	FileType   string
	HasChanges bool

	Cursor *cursor.Cursor // First cursor

	Line *line.Line // First line
}

func (t *Tab) LineCount() int { return t.Line.Count() }

func (t *Tab) Save() error {
	panic("unimplemented")
}
