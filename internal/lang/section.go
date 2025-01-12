package lang

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
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

		endIdx := line.RuneCount()
		if lineIdx == sec.EndLine {
			endIdx = sec.EndIdx
		}

		matches := sec.Region.match(line, startIdx, endIdx)

		// First create region default highlight
		hl := highlight.Highlight{
			LineNum:      lineIdx,
			StartRuneIdx: startIdx,
			EndRuneIdx:   endIdx,
			Style:        cfg.Colors.Style(sec.Region.Style),
		}
		hls = append(hls, hl)

		for _, m := range matches.CursorWords {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.CursorWord,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Attributes {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Attribute,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Bolds {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Bold,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Comments {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Comment,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Functions {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Function,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Headings {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Heading,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Italics {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Italic,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Keywords {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Keyword,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Links {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Link,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Metas {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Meta,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Monos {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Mono,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Numbers {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Number,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Operators {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Operator,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Strings {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.String,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Types {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Type,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Values {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Value,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Variables {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m.start, EndRuneIdx: m.end, Style: cfg.Colors.Variable,
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
