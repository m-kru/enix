package lang

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
)

// The line argument  must be the first line of file.
// StartLineIdx is the index of the first visible line.
// EndLineIdx is the index of the last visible line.
func (hl *Highlighter) Analyze(
	line *line.Line, // First tab line
	firstVisLineNum int,
	lastVisLineNum int,
	cursors []*cursor.Cursor,
) []highlight.Highlight {
	hls := make([]highlight.Highlight, 0, 1024)

	if len(hl.Regions) == 0 {
		return nil
	}

	hl.reset(cursors, firstVisLineNum, lastVisLineNum)

	for !hl.done() {
		hl.analyzeLine(line, &hls)
		line = line.Next
	}

	return hls
}
