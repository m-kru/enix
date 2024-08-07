package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/util"
)

// Currently view updating works in such a way, that the last cursor is always visible.
func (t *Tab) UpdateView() {
	c := t.Cursors.Last()
	t.View = t.View.MinAdjust(c.View())
}

func (t *Tab) RenderStatusLine(frame frame.Frame) {
	// Fill the background
	for i := 0; i < frame.Width; i++ {
		frame.SetContent(i, 0, ' ', t.Colors.StatusLine)
	}

	// Render file path or name
	path := t.Path
	if len(path) == 0 {
		path = t.Name
	}
	for i, r := range path {
		if i >= frame.Width {
			break
		}
		frame.SetContent(i, 0, r, t.Colors.StatusLine)
	}

	// Render extra status information
	b := strings.Builder{}
	if t.Cursors != nil {
		b.WriteString(
			fmt.Sprintf("%d:%d | ", t.Cursors.Line.Num(), t.Cursors.BufIdx+1),
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
			if t.HasCursorInLine(n) {
				frame.SetContent(i, y, r, t.Colors.Cursor)
			} else {
				frame.SetContent(i, y, r, t.Colors.LineNum)
			}
		}

		n++
		y++

		if y >= frame.Height || n > lineCount {
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

		line.Render(t.Config, t.Colors, frame.Line(0, renderedCount), t.View)

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

		if !t.View.IsVisible(c.View()) {
			c = c.Next
			continue
		}

		c.Render(t.Config, t.Colors, frame.Line(0, c.Line.Num()-t.View.Line), t.View)

		c = c.Next
	}
}

func (t *Tab) Render(frame frame.Frame) {
	// Render status line
	t.RenderStatusLine(frame.LastLine())

	if frame.Height > 1 {
		frame.Height -= 1
	}

	lineCount := t.LineCount()
	lineNumWidth := util.IntWidth(lineCount)

	// TODO: Should view Width and Height be set here?
	t.View.Width = frame.Width - lineNumWidth - 1
	t.View.Height = frame.Height
	t.UpdateView()

	// Render line numbers
	t.RenderLineNums(frame.Column(0, lineNumWidth))

	// Render lines
	linesFrame := frame.Column(lineNumWidth+1, frame.Width-lineNumWidth-1)
	t.RenderLines(linesFrame)

	// Render cursors
	t.RenderCursors(linesFrame)
}
