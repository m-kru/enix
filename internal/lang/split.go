package lang

import (
	"sort"

	"github.com/m-kru/enix/internal/line"
)

// SplitIntoSections splits file into sections based on defined regions.
func splitIntoSections(
	regions []*Region,
	line *line.Line, // First tab line
	startLine *line.Line,
	startLineIdx int,
	endLineIdx int,
) []Section {
	secs := make([]Section, 0, 128)
	reg := regions[0] // Current region
	sec := Section{
		StartLine: 1,
		StartIdx:  0,
		EndLine:   1,
		EndIdx:    0,
		Region:    reg,
	}

	lineIdx := 1
	if len(regions) == 1 {
		sec.StartLine = startLineIdx
		sec.EndLine = endLineIdx

		line = startLine
		for i := 0; i < endLineIdx-startLineIdx; i++ {
			line = line.Next
		}
		sec.EndIdx = line.RuneCount()

		secs = append(secs, sec)
		return secs
	}

	for {
		if lineIdx == startLineIdx {
			// Drop all irrelevant sections before first visible line.
			// TODO: Iterate in downward direction to improve performance.
			startIdx := 0
			for i := range secs {
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
				sec.EndLine = lineIdx
				sec.EndIdx = tok.StartIdx
				if sec.StartLine < sec.EndLine || sec.StartIdx != sec.EndIdx {
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
				if tok.EndIdx == line.RuneCount() && line.Next == nil {
					continue
				}

				sec.EndLine = lineIdx
				sec.EndIdx = tok.EndIdx
				secs = append(secs, sec)

				// Start new default section
				reg = regions[0]
				sec.Region = regions[0]
				if tok.EndIdx >= line.RuneCount() {
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
		if line.RuneCount() > 0 {
			sec.EndIdx = line.RuneCount()
		}
		secs = append(secs, sec)
	}

	return secs
}

func tokenizeLine(regions []*Region, line string) []RegionToken {
	toks := []RegionToken{}
	var tok RegionToken

	// Skip the default region
	for _, r := range regions[1:] {
		tok.Region = r

		tok.Start = true
		locs := r.StartRegex.Regex.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.StartIdx = l[0]
			tok.EndIdx = l[1]
			toks = append(toks, tok)
		}

		tok.Start = false
		locs = r.EndRegex.Regex.FindAllStringIndex(line, -1)
		for _, l := range locs {
			tok.StartIdx = l[0]
			tok.EndIdx = l[1]
			toks = append(toks, tok)
		}
	}

	sortFunc := func(i, j int) bool {
		ti := toks[i]
		tj := toks[j]

		if ti.StartIdx == tj.StartIdx {
			return ti.Start
		}

		return ti.StartIdx < tj.StartIdx
	}

	sort.Slice(toks, sortFunc)

	return toks
}
