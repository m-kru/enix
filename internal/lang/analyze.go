package lang

import (
	"regexp"

	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

// The line argument  must be the first line of file.
// StartLineIdx is the index of the first visible line.
// EndLineIdx is the index of the last visible line.
func (hl Highlighter) Analyze(
	line *line.Line, // First tab line
	startLine *line.Line,
	startLineIdx int,
	endLineIdx int,
	cursor *cursor.Cursor,
) []highlight.Highlight {
	highlights := []highlight.Highlight{}

	if len(hl.Regions) == 0 {
		return nil
	}

	for _, r := range hl.Regions {
		r.CursorWord = nil
	}

	cursorWord := ""
	if cursor != nil {
		cursorWord = cursor.GetWord()
	}
	if len(cursorWord) == 1 && util.IsBracket(rune(cursorWord[0])) {
		// Unimplemented
	} else if cursorWord != "" {
		re, err := regexp.Compile(`\b` + cursorWord + `\b`)
		if err == nil {
			for _, r := range hl.Regions {
				r.CursorWord = re
			}
		}
	}

	var hls []highlight.Highlight
	lineIdx := startLineIdx
	sections := splitIntoSections(hl.Regions, line, startLine, startLineIdx, endLineIdx)
	line = startLine
	for _, sec := range sections {
		// Progress to the start line of the current section or view.
		for {
			if lineIdx == sec.StartLine || lineIdx == startLineIdx {
				break
			}
			line = line.Next
			lineIdx++
		}

		hls, line = sec.Analyze(line, lineIdx)
		highlights = append(highlights, hls...)
		lineIdx = sec.EndLine
	}

	return highlights
}
