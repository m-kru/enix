package tab

import (
	"fmt"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/util"
)

func (t *Tab) RenderLineNums(frame frame.Frame) {
	n := t.View.LineNum
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
	lineIdx := t.View.LineNum
	renderedCount := 0
	line := t.Lines.Get(lineIdx)

	// TODO: Handle line clearing.
	for {
		if line == nil || renderedCount == frame.Height {
			break
		}

		// TODO: Fix view
		line.Render(t.Colors, frame.Line(0, renderedCount), t.View)

		line = line.Next
		lineIdx++
		renderedCount++
	}
}

func (t *Tab) RenderCursors(frame frame.Frame) {
	c := t.Cursors

	for {
		if c == nil {
			break
		}

		if !t.View.IsVisible(c) {
			c = c.Next
			continue
		}

		// TODO: Handle view
		c.Render(t.Colors, frame.Line(0, c.LineNum()-t.View.LineNum), t.View)

		c = c.Next
	}
}

func (t *Tab) Render(frame frame.Frame) {
	// Render line numbers
	lineCount := t.LineCount()
	lineNumWidth := util.IntWidth(lineCount)
	t.RenderLineNums(frame.Column(0, lineNumWidth))

	// Render lines
	t.View.Width = frame.Width - lineNumWidth
	t.View.Height = frame.Height
	linesFrame := frame.Column(lineNumWidth+1, frame.Width-lineNumWidth-1)
	t.RenderLines(linesFrame)

	// Render cursors
	t.RenderCursors(linesFrame)
}
