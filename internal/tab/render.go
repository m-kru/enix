package tab

import (
	"fmt"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/search"
)

func (tab *Tab) UpdateView() {
	if len(tab.Cursors) > 0 {
		tab.UpdateViewCursors()
	} else {
		tab.UpdateViewSelections()
	}
}

func (tab *Tab) UpdateViewCursors() {
	v := tab.Cursors[len(tab.Cursors)-1].View()
	tab.View = tab.View.MinAdjust(v)
}

func (tab *Tab) UpdateViewSelections() {
	sel := tab.LastSel()

	// Handle all the possible cases
	if sel.Next == nil {
		// This is single line view, behave the same as in case of cursor.
		v := sel.GetCursor().View()
		tab.View = tab.View.MinAdjust(v)
		return
	}
	// This is multi line view.
	// Adjust so that line with the cursor is visible
	// The cursor is not necessarily visible.
	line := sel.LineNum
	if !sel.CursorOnLeft() {
		// Cursor in last line.
		line = sel.Last().LineNum
	}
	tab.View = tab.View.LineAdjust(line)
}

func (tab *Tab) HasCursorInLine(line *line.Line) bool {
	if len(tab.Cursors) > 0 {
		for _, c := range tab.Cursors {
			if c.Line == line {
				return true
			}
		}
	} else {
		for _, sel := range tab.Selections {
			s := sel
			for s != nil {
				if s.Line == line && s.Cursor != nil {
					return true
				}
				s = s.Next
			}
		}
	}

	return false
}

func (tab *Tab) RenderLineNums(line *line.Line, lineNum int, frame frame.Frame) {
	y := 0

	for {
		str := fmt.Sprintf("%*d ", frame.Width-1, lineNum)
		for i, r := range str {
			if tab.HasCursorInLine(line) && i < len(str)-1 {
				frame.SetContent(i, y, r, cfg.Style.Cursor)
			} else {
				frame.SetContent(i, y, r, cfg.Style.LineNum)
			}
		}

		line = line.Next
		lineNum++
		y++

		if y >= frame.Height || lineNum > tab.LineCount {
			break
		}
	}

	// Clear remaining line numbers.
	for ; y < frame.Height; y++ {
		for i := range frame.Width {
			frame.SetContent(i, y, ' ', cfg.Style.Default)
		}
	}
}

// Line is first visible line
func (tab *Tab) RenderLines(line *line.Line, lineNum int, frame frame.Frame) {
	renderedCount := 0

	endLineIdx := tab.View.LastLine()
	if endLineIdx > tab.LineCount {
		endLineIdx = tab.LineCount
	}
	hls := tab.Highlighter.Analyze(tab.Lines, lineNum, endLineIdx, tab.Cursors)

	tab.SearchCtx.FirstVisLineNum = lineNum
	search.Search(tab.Lines, &tab.SearchCtx)
	finds := tab.SearchCtx.FindsFromVisible()

	for line != nil && renderedCount < frame.Height {
		hls, finds = line.Render(
			lineNum,
			frame.Line(0, renderedCount),
			tab.View,
			hls,
			finds,
		)

		line = line.Next
		lineNum++
		renderedCount++
	}

	// Line clearing
	for renderedCount < frame.Height {
		for w := range frame.Width {
			frame.SetContent(w, renderedCount, ' ', cfg.Style.Default)
		}

		renderedCount++
	}
}

func (tab *Tab) RenderCursors(frame frame.Frame) {
	for _, c := range tab.Cursors {
		if !tab.View.IsVisible(c.View()) {
			continue
		}

		c.Render(frame.Line(0, c.LineNum-tab.View.Line), tab.View)
	}
}

func (tab *Tab) RenderSelections(frame frame.Frame) {
	for _, s := range tab.Selections {
		if !tab.View.IsVisible(s.View()) {
			continue
		}

		s.Render(frame, tab.View)
	}
}

func (tab *Tab) Render() {
	frame := *tab.Frame

	lineNumWidth := tab.LineNumWidth()

	// TODO: Should view Width and Height be set here?
	tab.View.Width = frame.Width - lineNumWidth - 1
	tab.View.Height = frame.Height

	lineNum := tab.View.Line
	line := tab.Lines.Get(lineNum)

	// Render line numbers
	lineNumFrame := frame.ColumnSubframe(0, lineNumWidth+1)
	if lineNumFrame.Screen != nil {
		tab.RenderLineNums(line, lineNum, lineNumFrame)
	}

	// Render lines and cursors
	linesFrame := frame.ColumnSubframe(lineNumWidth+1, frame.Width-lineNumWidth-1)
	if linesFrame.Screen != nil {
		tab.RenderLines(line, lineNum, linesFrame)

		if len(tab.Cursors) > 0 {
			tab.RenderCursors(linesFrame)
		} else {
			tab.RenderSelections(linesFrame)
		}
	}
}
