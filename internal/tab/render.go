package tab

import (
	"fmt"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/util"
)

func (t *Tab) RenderLineNums(frame frame.Frame) {
	n := t.FirstVisLine.Num()
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

func (t *Tab) RenderLines(frame frame.Frame) {
	lineIdx := t.FirstVisLine.Num()
	renderedCount := 0
	line := t.Lines.Get(lineIdx)
	// TODO: Handle line clearing.
	for {
		if line == nil || renderedCount == frame.Height {
			break
		}

		line.Render(t.Colors, frame.Line(0, renderedCount), 0)

		line = line.Next
		lineIdx++
		renderedCount++
	}

	t.LastVisLine = t.FirstVisLine.Get(renderedCount)
}

func (t *Tab) Render(frame frame.Frame) {
	// Render line numbers
	lineCount := t.LineCount()
	lineNumWidth := util.IntWidth(lineCount)
	t.RenderLineNums(frame.Column(0, lineNumWidth))

	// Render lines
	linesFrame := frame.Column(lineNumWidth + 1, frame.Width-lineNumWidth-1)
	t.RenderLines(linesFrame)

	// Render cursors
}
