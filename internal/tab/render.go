package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/util"
)

// Currently view updating works in such a way, that the last cursor is always visible.
func (tab *Tab) UpdateView() {
	c := tab.Cursors.Last()
	tab.View = tab.View.MinAdjust(c.View())
}

func (tab *Tab) RenderStatusLine(frame frame.Frame) {
	// Fill the background
	for i := 0; i < frame.Width; i++ {
		frame.SetContent(i, 0, ' ', tab.Colors.StatusLine)
	}

	// Render file path
	for i, r := range tab.Path {
		if i >= frame.Width {
			break
		}
		frame.SetContent(i, 0, r, tab.Colors.StatusLine)
	}

	// Render extra status information
	b := strings.Builder{}

	if tab.InInsertMode {
		b.WriteString("insert ")
	}

	if tab.Cursors != nil {
		b.WriteString(
			fmt.Sprintf("%d:%d | ", tab.Cursors.Line.Num(), tab.Cursors.BufIdx+1),
		)
	}
	b.WriteString(fmt.Sprintf("%s ", tab.FileType))
	statusStr := b.String()

	if len(statusStr) > frame.Width {
		return
	}

	startIdx := frame.Width - len(statusStr)
	for i, r := range statusStr {
		style := tab.Colors.StatusLine
		if tab.InInsertMode && i < 6 {
			style = tab.Colors.InsertMark
		}
		frame.SetContent(startIdx+i, 0, r, style)
	}
}

func (tab *Tab) RenderLineNums(frame frame.Frame) {
	n := tab.View.Line
	y := frame.Y
	lineCount := tab.LineCount()

	for {
		str := fmt.Sprintf("%*d", frame.Width, n)
		for i, r := range str {
			if tab.HasCursorInLine(n) {
				frame.SetContent(i, y, r, tab.Colors.Cursor)
			} else {
				frame.SetContent(i, y, r, tab.Colors.LineNum)
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
			frame.SetContent(i, y, ' ', tab.Colors.Default)
		}
	}
}

func (tab *Tab) RenderLines(frame frame.Frame) {
	lineNum := tab.View.Line
	renderedCount := 0
	line := tab.Lines.Get(lineNum)

	// TODO: Handle line clearing.
	for {
		if line == nil || renderedCount == frame.Height {
			break
		}

		line.Render(tab.Config, tab.Colors, frame.Line(0, renderedCount), tab.View)

		line = line.Next
		lineNum++
		renderedCount++
	}
}

func (tab *Tab) RenderCursors(frame frame.Frame) {
	c := tab.Cursors

	for {
		if c == nil {
			break
		}

		if !tab.View.IsVisible(c.View()) {
			c = c.Next
			continue
		}

		primary := false
		if tab.HasFocus && c.Next == nil {
			primary = true
		}
		c.Render(tab.Config, tab.Colors, frame.Line(0, c.Line.Num()-tab.View.Line), tab.View, primary)

		c = c.Next
	}
}

func (tab *Tab) Render(frame frame.Frame) {
	// Render status line
	tab.RenderStatusLine(frame.LastLine())

	if frame.Height > 1 {
		frame.Height -= 1
	}

	lineCount := tab.LineCount()
	lineNumWidth := util.IntWidth(lineCount)

	// TODO: Should view Width and Height be set here?
	tab.View.Width = frame.Width - lineNumWidth - 1
	tab.View.Height = frame.Height

	// Render line numbers
	tab.RenderLineNums(frame.Column(0, lineNumWidth))

	// Render lines
	linesFrame := frame.Column(lineNumWidth+1, frame.Width-lineNumWidth-1)
	tab.RenderLines(linesFrame)

	// Render cursors
	tab.RenderCursors(linesFrame)
}
