package highlight

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/xxx"
)

type Section struct {
	StartLine int
	StartIdx  int
	EndLine   int
	EndIdx    int

	Region *Region
}

func (sec Section) Analyze(line *line.Line, startLineIdx int, colors *cfg.Colorscheme) ([]xxx.Highlight, *line.Line) {
	hls := make([]xxx.Highlight, 0, 16*(sec.EndLine-sec.StartLine+1))

	for lineIdx := startLineIdx; lineIdx <= sec.EndLine; lineIdx++ {
		if line.Len() == 0 {
			line = line.Next
			continue
		}

		startIdx := 0
		if lineIdx == sec.StartLine {
			startIdx = sec.StartIdx
		}

		endIdx := line.Len() - 1
		if lineIdx == sec.EndLine {
			endIdx = sec.EndIdx
		}

		matches := sec.Region.Match(line, startIdx, endIdx)

		// First create region default highlight
		hl := xxx.Highlight{
			Line:     lineIdx,
			StartIdx: startIdx,
			EndIdx:   endIdx,
			Style:    colors.Style(sec.Region.Style),
		}
		hls = append(hls, hl)
		hlsStartIdx := len(hls) - 1

		for _, m := range matches.CursorWords {
			hl := xxx.Highlight{Line: lineIdx, StartIdx: m[0], EndIdx: m[1], Style: colors.CursorWord}
			insertHighlight(&hls, hlsStartIdx, hl)
		}

		for _, m := range matches.Types {
			hl := xxx.Highlight{Line: lineIdx, StartIdx: m[0], EndIdx: m[1], Style: colors.Type}
			insertHighlight(&hls, hlsStartIdx, hl)
		}

		line = line.Next
	}

	return hls, line
}

func insertHighlight(hls *[]xxx.Highlight, startIdx int, hl xxx.Highlight) {
	for i := startIdx; i < len(*hls); i++ {
		if !(*hls)[i].Contains(hl) {
			continue
		}

		newHls := (*hls)[i].Split(hl)
		for range len(newHls) - 1 {
			*hls = append(*hls, xxx.Highlight{})
		}

		if len(newHls) > 1 {
			for j := 0; j < len(*hls)-i-len(newHls); j++ {
				(*hls)[len(*hls)-1-j] = (*hls)[(len(*hls)-1-j)-len(newHls)]
			}
		}

		for j := range len(newHls) {
			(*hls)[i+j] = newHls[j]
		}

		break
	}
}
