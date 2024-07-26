package tab

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

type Tab struct {
	Colors *cfg.Colorscheme

	Screen tcell.Screen
	StartX int
	EndX   int
	StartY int
	EndY   int

	Name       string // Path of the file
	Newline    string // Newline encoding
	FileType   string
	HasChanges bool

	Cursor *cursor.Cursor // First cursor

	Lines *line.Line // First line pointer

	FirstVisLineIdx int // First visible line index
}

func (t *Tab) LineCount() int { return t.Lines.Count() }

func (t *Tab) Save() error {
	panic("unimplemented")
}

func (t *Tab) RenderLineNums(frame frame.Frame) {
	n := t.FirstVisLineIdx
	y := t.StartY
	lineCount := t.LineCount()

	for {
		str := fmt.Sprintf("%*d", frame.Width, n)
		for i, r := range str {
			t.Screen.SetContent(i+t.StartX, y, r, nil, t.Colors.LineNum)
		}

		n++
		y++

		if y > t.EndY-1 || n > lineCount {
			break
		}
	}

	// Clear remaining line numbers.
	for ; y < t.EndY-1; y++ {
		for i := 0; i < frame.Width; i++ {
			t.Screen.SetContent(i+t.StartX, y, ' ', nil, t.Colors.Default)
		}
	}
}

func (t *Tab) Render(frame frame.Frame) {
	lineCount := t.LineCount()
	lineNumWidth := util.IntWidth(lineCount)
	t.RenderLineNums(frame.Column(0, lineNumWidth))
}
