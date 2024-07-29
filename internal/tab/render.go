package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/util"
)

func (t *Tab) RenderStatusLine(frame frame.Frame) {
	// Fill the background
	for i := 0; i < frame.Width; i++ {
		frame.SetContent(i, 0, ' ', t.Colors.StatusLine)
	}

	// TODO: Handle case when file path is wider than frame width.
	for i, r := range t.Name {
		frame.SetContent(i, 0, r, t.Colors.StatusLine)
	}

	b := strings.Builder{}
	if t.Cursors != nil {
		b.WriteString(
			fmt.Sprintf("%d:%d | ", t.Cursors.LineNum(), t.Cursors.Column()),
		)
	}
	b.WriteString(fmt.Sprintf("%s ", t.FileType))
	statusStr := b.String()

	if len(statusStr) > frame.Width {
		return
	}

	startIdx := frame.Width - len(statusStr)
	for i, r := range statusStr {
		frame.SetContent(startIdx+i, 0, r, t.Colors.StatusLine)
	}
}

func (t *Tab) RenderLineNums(frame frame.Frame) {
	n := t.View.Line
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
	lineNum := t.View.Line
	renderedCount := 0
	line := t.Lines.Get(lineNum)

	// TODO: Handle line clearing.
	for {
		if line == nil || renderedCount == frame.Height {
			break
		}

		// TODO: Fix view
		line.Render(t.Colors, frame.Line(0, renderedCount), t.View)

		line = line.Next
		lineNum++
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
		c.Render(t.Colors, frame.Line(0, c.LineNum()-t.View.Line), t.View)

		c = c.Next
	}
}

func (t *Tab) Render(frame frame.Frame) {
	// Render status line
	t.RenderStatusLine(frame.LastLine())

	if frame.Height > 1 {
		frame.Height -= 1
	}

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
