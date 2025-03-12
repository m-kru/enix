package lang

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

type Section struct {
	StartLine int
	StartIdx  int
	EndLine   int
	EndIdx    int

	Region *Region
}

func (sec Section) Analyze(line *line.Line, startLineIdx int) ([]highlight.Highlight, *line.Line) {
	hls := make([]highlight.Highlight, 0, 16*(sec.EndLine-sec.StartLine+1))

	for lineIdx := startLineIdx; lineIdx <= sec.EndLine; lineIdx++ {
		if line.RuneCount() == 0 {
			line = line.Next
			continue
		}

		startIdx := 0
		if lineIdx == sec.StartLine {
			startIdx = sec.StartIdx
		}

		endIdx := len(line.Buf)
		if lineIdx == sec.EndLine {
			endIdx = sec.EndIdx
		}

		matches := sec.Region.match(line, startIdx, endIdx)

		// First create region default highlight
		hl := highlight.Highlight{
			LineNum:      lineIdx,
			StartRuneIdx: util.ByteIdxToRuneIdx(line.Buf, startIdx),
			EndRuneIdx:   util.ByteIdxToRuneIdx(line.Buf, endIdx),
			Style:        cfg.Style.Get(sec.Region.Style),
		}
		hls = append(hls, hl)

		for _, m := range matches.CursorWords {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.CursorWord,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Attributes {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Attribute,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Bolds {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Bold,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Comments {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Comment,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Headings {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Heading,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Italics {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Italic,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Keywords {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Keyword,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Metas {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Meta,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Monos {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Mono,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Numbers {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Number,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Operators {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Operator,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Strings {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.String,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Types {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Type,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Values {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Value,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Variables {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Style.Variable,
			}
			insertHighlight(&hls, hl)
		}

		if lineIdx < sec.EndLine {
			line = line.Next
		}
	}

	return hls, line
}

func insertHighlight(hls *[]highlight.Highlight, hl highlight.Highlight) {
	for i := range len(*hls) {
		if !(*hls)[i].Contains(hl) {
			continue
		}

		newHls := (*hls)[i].Split(hl)
		for i := range len(newHls) - 1 {
			hl := newHls[i]
			*hls = append(*hls, hl)
		}

		if len(newHls) > 1 {
			for j := range len(*hls) - i - len(newHls) {
				(*hls)[len(*hls)-1-j] = (*hls)[(len(*hls)-1-j)-len(newHls)+1]
			}
		}

		for j := range newHls {
			(*hls)[i+j] = newHls[j]
		}

		break
	}
}
