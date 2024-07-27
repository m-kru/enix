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

	Name       string // Path of the file
	Newline    string // Newline encoding
	FileType   string
	HasChanges bool

	Cursors *cursor.Cursor // First cursor

	Lines *line.Line // First line pointer

	FirstVisLineIdx int // First visible line index
}

func (t *Tab) LineCount() int { return t.Lines.Count() }

func (t *Tab) Save() error {
	panic("unimplemented")
}

func (t *Tab) RenderLineNums(frame frame.Frame) {
	n := t.FirstVisLineIdx
	y := frame.Y
	lineCount := t.LineCount()

	for {
		str := fmt.Sprintf("%*d", frame.Width, n)
		for i, r := range str {
			frame.SetContent(i, y, r, t.Colors.LineNum)
		}

		n++
		y++

		if y > frame.Height || n > lineCount {
			break
		}
	}

	// Clear remaining line numbers.
	for ; y < frame.Height; y++ {
		for i := 0; i < frame.Width; i++ {
			frame.SetContent(i, y, ' ', t.Colors.Default)
		}
	}
}

func (t *Tab) Render(frame frame.Frame) {
	lineCount := t.LineCount()
	lineNumWidth := util.IntWidth(lineCount)
	t.RenderLineNums(frame.Column(0, lineNumWidth))

	linesFrame := frame.Column(lineNumWidth, frame.Width-lineNumWidth)

	lineIdx := t.FirstVisLineIdx
	renderedCount := 0
	line := t.Lines.Get(lineIdx)
	// TODO: Handle line clearing.
	for {
		if line == nil || lineCount == frame.Height {
			break
		}

		lineFrame := linesFrame.Line(lineNumWidth, renderedCount)
		line.Render(t.Colors, lineFrame, 0)

		line = line.Next
		lineIdx++
		renderedCount++
	}
}
