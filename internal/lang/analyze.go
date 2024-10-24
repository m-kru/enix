package lang

import (
	"regexp"
	"sort"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

// The line argument  must be the first line of file.
// StartLineIdx is the index of the first visible line.
// EndLineIdx is the index of the last visible line.
func (hl Highlighter) Analyze(
	line *line.Line,
	startLineIdx int,
	endLineIdx int,
	cursor *cursor.Cursor,
	colors *cfg.Colorscheme,
) []highlight.Highlight {
	highlights := []highlight.Highlight{}

	if len(hl.Regions) == 0 {
		return nil
	}

	cursorWord := cursor.GetWord()
	if len(cursorWord) == 1 && util.IsBracket(rune(cursorWord[0])) {
		// Unimplemented
	} else if cursorWord != "" {
		re, err := regexp.Compile(`\b` + cursorWord + `\b`)
		if err == nil {
			for _, r := range hl.Regions {
				r.CursorWord = re
			}
		}
	} else {
		for _, r := range hl.Regions {
			r.CursorWord = nil
		}
	}

	var hls []highlight.Highlight
	lineIdx := startLineIdx
	sections, line := hl.splitIntoSections(line, startLineIdx, endLineIdx)
	for _, sec := range sections {
		// Progress to the start line of the current section or view.
		for {
			if lineIdx == sec.StartLine || lineIdx == startLineIdx {
				break
			}
			line = line.Next
			lineIdx++
		}

		hls, line = sec.Analyze(line, lineIdx, colors)
		highlights = append(highlights, hls...)
		lineIdx = sec.EndLine
	}

	return highlights
}

// SplitIntoSections splits file into sections based on defined regions.
func (hl Highlighter) splitIntoSections(
	line *line.Line,
	startLineIdx int,
	endLineIdx int,
) ([]Section, *line.Line) {
	secs := []Section{}
	reg := hl.Regions[0] // Current region
	sec := Section{
		StartLine: 1,
		StartIdx:  0,
		EndLine:   1,
		EndIdx:    0,
		Region:    reg,
	}

	startLine := line
	lineIdx := 1

	for {
		if lineIdx == startLineIdx {
			startLine = line

			// Drop all irrelevant sections before first visible line.
			if len(secs) > 1 {
				firstSecIdx := len(secs) - 1
				for i := len(secs) - 2; i >= 0; i-- {
					if secs[i].EndLine == startLineIdx {
						firstSecIdx = i
					} else {
						break
					}
				}
				secs = secs[firstSecIdx:]
			}
		}

		lineToks := hl.tokenizeLine(line.String())

		for _, tok := range lineToks {
			if reg.Name == "Default" {
				if !tok.Start {
					continue
				}

				// Append previous default section
				if tok.StartIdx == 0 {
					if line.Prev != nil {
						sec.EndLine = lineIdx - 1
						sec.EndIdx = line.Prev.Len() - 1
					}
				} else {
					sec.EndLine = lineIdx
					sec.EndIdx = tok.StartIdx - 1
				}
				if sec.StartLine != sec.EndLine || sec.StartIdx != sec.EndIdx {
					secs = append(secs, sec)
				}

				reg = tok.Region
				sec.Region = reg
				sec.StartLine = lineIdx
				sec.StartIdx = tok.StartIdx
			} else {
				if tok.Start || tok.Region != reg || tok.StartIdx == sec.StartIdx {
					continue
				}
				sec.EndLine = lineIdx
				sec.EndIdx = tok.EndIdx
				secs = append(secs, sec)

				// Start new default section
				reg = hl.Regions[0]
				sec.Region = hl.Regions[0]
				if tok.EndIdx >= line.Len() {
					if line.Next != nil {
						sec.StartLine = lineIdx + 1
						sec.StartIdx = 0
					}
				} else {
					sec.StartLine = lineIdx
					sec.StartIdx = tok.EndIdx + 1
				}
			}
		}

		if lineIdx == endLineIdx {
			break
		}
		line = line.Next
		lineIdx++
	}

	if reg.Name == "Default" {
		sec.EndLine = lineIdx
		sec.EndIdx = 0
		if line.Len() > 0 {
			sec.EndIdx = line.Len() - 1
		}
		secs = append(secs, sec)
	}

	return secs, startLine
}

func (hl Highlighter) tokenizeLine(line string) []RegionToken {
	toks := []RegionToken{}
	var tok RegionToken

	// Skip the default region
	for _, r := range hl.Regions[1:] {
		tok.Region = r

		tok.Start = true
		locs := r.StartRegexp.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.StartIdx = l[0]
			tok.EndIdx = l[1] - 1
			toks = append(toks, tok)
		}

		tok.Start = false
		locs = r.EndRegexp.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.StartIdx = l[0]
			tok.EndIdx = l[1] - 1
			toks = append(toks, tok)
		}
	}

	sortFunc := func(i, j int) bool {
		return toks[i].StartIdx < toks[j].StartIdx
	}

	sort.Slice(toks, sortFunc)

	return toks
}
