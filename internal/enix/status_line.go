package enix

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/tab"
)

// tab must be pointer to the current tab.
func renderStatusLine(frame frame.Frame, tab *tab.Tab) {
	// Fill the background
	for i := range frame.Width {
		frame.SetContent(i, 0, ' ', cfg.Style.StatusLine)
	}

	// Render file path
	for i, r := range tab.Path {
		if i >= frame.Width {
			break
		}
		frame.SetContent(i, 0, r, cfg.Style.StatusLine)
	}

	// Render extra status information
	b := strings.Builder{}

	repCountStartIdx := 0
	if tab.RepCount != 0 {
		b.WriteString(fmt.Sprintf("%d ", tab.RepCount))
	}
	repCountEndIdx := b.Len() - 1

	stateStartIdx := b.Len()
	if tab.State != "" {
		b.WriteString(fmt.Sprintf("%s ", tab.State))
	}
	stateEndIdx := b.Len() - 1

	findStartIdx := b.Len()
	if tab.SearchCtx.Regexp != nil {
		if len(tab.SearchCtx.Finds) == 1 {
			b.WriteString("1 find ")
		} else {
			b.WriteString(fmt.Sprintf("%d finds ", len(tab.SearchCtx.Finds)))
		}
	}
	findEndIdx := b.Len() - 1

	var c *cursor.Cursor
	if len(tab.Cursors) > 0 {
		c = tab.Cursors[len(tab.Cursors)-1]
	} else {
		c = tab.LastSel().GetCursor()
		if len(tab.Selections) == 1 {
			b.WriteString("1 sel ")
		} else {
			b.WriteString(fmt.Sprintf("%d sels ", len(tab.Selections)))
		}
	}
	b.WriteString(fmt.Sprintf("%d:%d | ", c.LineNum, c.RuneIdx+1))

	b.WriteString(fmt.Sprintf("%s ", tab.Filetype))
	statusStr := b.String()

	if len(statusStr) > frame.Width {
		return
	}

	startIdx := frame.Width - len(statusStr)
	for i, r := range statusStr {
		style := cfg.Style.StatusLine
		if tab.RepCount > 0 && repCountStartIdx <= i && i < repCountEndIdx {
			style = cfg.Style.RepCount
		} else if tab.State != "" && stateStartIdx <= i && i < stateEndIdx {
			style = cfg.Style.StateMark
		} else if len(tab.SearchCtx.Finds) > 0 && findStartIdx <= i && i < findEndIdx {
			style = cfg.Style.FindMark
		}

		frame.SetContent(startIdx+i, 0, r, style)
	}
}
