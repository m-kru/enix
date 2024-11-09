package search

import (
	"github.com/m-kru/enix/internal/find"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

// Search returns a list of finds and an index of the first find after the start line.
// It doesn't mean that the first find after the start line is visible, as it may be
// placed after the last visible line.
//
// line points to the first line of tab.
// startLine is the index of the first visible line.
func Search(line *line.Line, startLine int, ctx *Context) {
	ctx.Finds = make([]find.Find, 0, 64)
	ctx.StartIdx = -1

	if ctx.Regexp == nil {
		return
	}

	lineNum := 1
	for {
		if line == nil {
			break
		}

		str := line.String()

		matches := ctx.Regexp.FindAllStringIndex(str, -1)
		if len(matches) == 0 {
			line = line.Next
			lineNum++
			continue
		}

		if lineNum >= startLine && ctx.StartIdx == -1 {
			ctx.StartIdx = len(ctx.Finds)
		}

		for _, m := range matches {
			f := find.Find{
				LineNum:      lineNum,
				StartRuneIdx: util.ByteIdxToRuneIdx(str, m[0]),
				EndRuneIdx:   util.ByteIdxToRuneIdx(str, m[1]),
			}
			ctx.Finds = append(ctx.Finds, f)
		}

		line = line.Next
		lineNum++
	}
}
