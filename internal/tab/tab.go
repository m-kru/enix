package tab

import (
	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
)

type Tab struct {
	Colors *cfg.Colorscheme

	Screen tcell.Screen

	Name       string // Path of the file
	Newline    string // Newline encoding
	FileType   string
	HasChanges bool

	Cursors *cursor.Cursor // First cursor

	Lines *line.Line // First line pointer

	FirstVisLine *line.Line // First visible line
	LastVisLine  *line.Line // Last visible line
}

func (t *Tab) LineCount() int { return t.Lines.Count() }

func (t *Tab) Save() error {
	panic("unimplemented")
}

// IsLineVisible returns info whether line is visible on the screen.
// If visible, the first return is y coordinate of frame used while
// rendering lines.
func (t *Tab) IsLineVisible(l *line.Line) (int, bool) {
	y := 0

	if l == t.FirstVisLine {
		return y, true
	}

	for {
		l = l.Prev
		y++

		if l == t.LastVisLine || l == nil {
			return 0, false
		} else if l == t.FirstVisLine {
			return y, true
		}
	}
}
