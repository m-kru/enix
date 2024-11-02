package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/frame"
)

// Currently view updating works in such a way, that the last cursor is always visible.
func (tab *Tab) UpdateView() {
	c := tab.Cursors[len(tab.Cursors)-1]
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

	repCountStartIdx := 0
	if tab.RepCount != 0 {
		b.WriteString(fmt.Sprintf("%d ", tab.RepCount))
	}
	repCountEndIdx := b.Len()

	stateStartIdx := b.Len()
	if tab.State != "" {
		b.WriteString(fmt.Sprintf("%s ", tab.State))
	}
	stateEndIdx := b.Len()

	if len(tab.Cursors) > 0 {
		b.WriteString(
			fmt.Sprintf("%d:%d | ", tab.Cursors[0].Line.Num(), tab.Cursors[0].BufIdx+1),
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
		if tab.RepCount > 0 && repCountStartIdx <= i && i < repCountEndIdx {
			style = tab.Colors.RepCount
		} else if tab.State != "" && stateStartIdx <= i && i < stateEndIdx {
			style = tab.Colors.StateMark
		}

		frame.SetContent(startIdx+i, 0, r, style)
	}
}

func (tab *Tab) RenderLineNums(frame frame.Frame) {
	n := tab.View.Line
	y := 0

	for {
		str := fmt.Sprintf("%*d ", frame.Width-1, n)
		for i, r := range str {
			if tab.HasCursorInLine(n) && i < len(str)-1 {
				frame.SetContent(i, y, r, tab.Colors.Cursor)
			} else {
				frame.SetContent(i, y, r, tab.Colors.LineNum)
			}
		}

		n++
		y++

		if y >= frame.Height || n > tab.LineCount {
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

	endLineIdx := tab.View.LastLine()
	if endLineIdx > tab.Lines.Last().Num() {
		endLineIdx = tab.Lines.Last().Num()
	}
	hls := tab.Highlighter.Analyze(
		tab.Lines, lineNum, endLineIdx, tab.Cursors[len(tab.Cursors)-1], tab.Colors,
	)

	for {
		if line == nil || renderedCount == frame.Height {
			break
		}

		hls = line.Render(tab.Config, tab.Colors, frame.Line(0, renderedCount), tab.View, hls)

		line = line.Next
		lineNum++
		renderedCount++
	}

	// Line clearing
	for {
		if renderedCount == frame.Height {
			break
		}

		for w := range frame.Width {
			frame.SetContent(w, renderedCount, ' ', tab.Colors.Default)
		}

		renderedCount++
	}
}

func (tab *Tab) RenderCursors(frame frame.Frame) {
	// This is required for view commands, as the primary cursors is rendered
	// by the tcell all the time.
	if tab.HasFocus {
		frame.HideCursor()
	}

	for i, c := range tab.Cursors {
		if !tab.View.IsVisible(c.View()) {
			continue
		}

		primary := false
		if tab.HasFocus && i == len(tab.Cursors)-1 {
			primary = true
		}
		c.Render(tab.Config, tab.Colors, frame.Line(0, c.Line.Num()-tab.View.Line), tab.View, primary)
	}
}

func (tab *Tab) Render(frame frame.Frame) {
	// Render status line
	tab.RenderStatusLine(frame.LastLine())

	if frame.Height > 1 {
		frame.Height -= 1
	}

	lineNumWidth := tab.LineNumWidth()

	// TODO: Should view Width and Height be set here?
	tab.View.Width = frame.Width - lineNumWidth - 1
	tab.View.Height = frame.Height

	// Render line numbers
	lineNumFrame := frame.Column(0, lineNumWidth+1)
	if lineNumFrame.Screen != nil {
		tab.RenderLineNums(lineNumFrame)
	}

	// Render lines and cursors
	linesFrame := frame.Column(lineNumWidth+1, frame.Width-lineNumWidth-1)
	if linesFrame.Screen != nil {
		tab.RenderLines(linesFrame)
		tab.RenderCursors(linesFrame)
	}
}
