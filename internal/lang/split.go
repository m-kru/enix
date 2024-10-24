package lang

import (
	"sort"

	"github.com/m-kru/enix/internal/line"
)

// SplitIntoSections splits file into sections based on defined regions.
func splitIntoSections(
	regions []*Region,
	line *line.Line,
	startLineIdx int,
	endLineIdx int,
) ([]Section, *line.Line) {
	secs := make([]Section, 0, 128)
	reg := regions[0] // Current region
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
			// TODO: Iterate in downward direction to improve performance.
			startIdx := 0
			for i, _ := range secs {
				if secs[i].EndLine < startLineIdx {
					startIdx = i + 1
				} else {
					break
				}
			}
			secs = secs[startIdx:]
		}

		lineToks := tokenizeLine(regions, line.String())

		for i, tok := range lineToks {
			if reg.Name == "Default" {
				if !tok.Start || (i > 0 && tok.Overlaps(lineToks[i-1])) {
					continue
				}

				// Append previous default section
				if tok.StartIdx == 0 {
					if line.Prev != nil {
						sec.EndLine = lineIdx - 1
						sec.EndIdx = line.Prev.Len()
					}
				} else {
					sec.EndLine = lineIdx
					sec.EndIdx = tok.StartIdx
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

				// Check if this is the end of text
				if tok.EndIdx == line.Len() && line.Next == nil {
					continue
				}

				sec.EndLine = lineIdx
				sec.EndIdx = tok.EndIdx
				secs = append(secs, sec)

				// Start new default section
				reg = regions[0]
				sec.Region = regions[0]
				if tok.EndIdx >= line.Len() {
					if line.Next != nil {
						sec.StartLine = lineIdx + 1
						sec.StartIdx = 0
					}
				} else {
					sec.StartLine = lineIdx
					sec.StartIdx = tok.EndIdx
				}
			}
		}

		if lineIdx == endLineIdx {
			break
		}
		line = line.Next
		lineIdx++
	}

	// Terminate unterminated section.
	if sec.EndLine < sec.StartLine || sec.StartIdx >= sec.EndIdx {
		sec.EndLine = lineIdx
		sec.EndIdx = 0
		if line.Len() > 0 {
			sec.EndIdx = line.Len()
		}
		secs = append(secs, sec)
	}

	return secs, startLine
}

func tokenizeLine(regions []*Region, line string) []RegionToken {
	toks := []RegionToken{}
	var tok RegionToken

	// Skip the default region
	for _, r := range regions[1:] {
		tok.Region = r

		tok.Start = true
		locs := r.StartRegexp.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.StartIdx = l[0]
			tok.EndIdx = l[1]
			toks = append(toks, tok)
		}

		tok.Start = false
		locs = r.EndRegexp.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.StartIdx = l[0]
			tok.EndIdx = l[1]
			toks = append(toks, tok)
		}
	}

	sortFunc := func(i, j int) bool {
		return toks[i].StartIdx < toks[j].StartIdx
	}

	sort.Slice(toks, sortFunc)

	return toks
}
