package highlight

import (
	"regexp"
	"sort"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/xxx"
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
) []xxx.Highlight {
	highlights := []xxx.Highlight{}

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

	var hls []xxx.Highlight
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
				if tok.Idx == 0 {
					if line.Prev != nil {
						sec.EndLine = lineIdx - 1
						sec.EndIdx = line.Prev.Len() - 1
					}
				} else {
					sec.EndLine = lineIdx
					sec.EndIdx = tok.Idx - 1
				}
				if sec.StartLine != sec.EndLine || sec.StartIdx != sec.EndIdx {
					secs = append(secs, sec)
				}

				reg = tok.Region
				sec.Region = reg
				sec.StartLine = lineIdx
				sec.StartIdx = tok.Idx
			} else {
				if tok.Start || tok.Region != reg || tok.Idx == sec.StartIdx {
					continue
				}
				sec.EndLine = lineIdx
				sec.EndIdx = tok.Idx
				secs = append(secs, sec)

				// Start new default section
				sec.Region = hl.Regions[0]
				if tok.Idx >= line.Len() {
					if line.Next != nil {
						sec.StartLine = lineIdx + 1
						sec.StartIdx = 0
					}
				} else {
					sec.StartLine = lineIdx
					sec.StartIdx = tok.Idx + 1
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
			tok.Idx = l[0]
			toks = append(toks, tok)
		}

		tok.Start = false
		locs = r.EndRegexp.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.Idx = l[1]
			toks = append(toks, tok)
		}
	}

	sortFunc := func(i, j int) bool {
		return toks[i].Idx < toks[j].Idx
	}

	sort.Slice(toks, sortFunc)

	return toks
}
