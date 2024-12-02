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

func (sec Section) Analyze(line *line.Line, startLineIdx int, colors *cfg.Colorscheme) ([]highlight.Highlight, *line.Line) {
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

		matches := sec.Region.Match(line, startIdx, endIdx)

		// First create region default highlight
		hl := highlight.Highlight{
			LineNum:      lineIdx,
			StartRuneIdx: startIdx,
			EndRuneIdx:   endIdx,
			Style:        colors.Style(sec.Region.Style),
		}
		hls = append(hls, hl)

		for _, m := range matches.CursorWords {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.CursorWord,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Attributes {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Attribute,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Builtins {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Builtin,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.EscapeSequences {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.EscapeSequence,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.FormatSpecifiers {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.FormatSpecifier,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Functions {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Function,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Keywords {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Keyword,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Metas {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Meta,
			}
			insertHighlight(&hls, hl)
		}

		for _, n := range matches.Numbers {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: n[0], EndRuneIdx: n[1], Style: colors.Number,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Operators {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Operator,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Strings {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.String,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.ToDos {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.ToDo,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Types {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Type,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Values {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Value,
			}
			insertHighlight(&hls, hl)
		}

		for _, m := range matches.Variables {
			hl := highlight.Highlight{
				LineNum: lineIdx, StartRuneIdx: m[0], EndRuneIdx: m[1], Style: colors.Variable,
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

		for j := range len(newHls) {
			(*hls)[i+j] = newHls[j]
		}

		break
	}
}
