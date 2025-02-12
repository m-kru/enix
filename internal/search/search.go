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
func Search(line *line.Line, ctx *Context) []find.Find {
	if ctx.Regexp == nil {
		return nil
	}

	ctx.FirstVisFindIdx = -1

	// There were no changes, just update start index.
	if !ctx.Modified {
		for i, f := range ctx.Finds {
			if f.LineNum >= ctx.FirstVisLineNum {
				ctx.FirstVisFindIdx = i
				break
			}
		}
		return ctx.Finds
	}

	ctx.Finds = make([]find.Find, 0, 64)

	lineNum := 1
	for {
		if line == nil {
			break
		}

		matches := ctx.Regexp.FindAllIndex(line.Buf, -1)
		if len(matches) == 0 {
			line = line.Next
			lineNum++
			continue
		}

		if ctx.FirstVisFindIdx == -1 && lineNum >= ctx.FirstVisLineNum {
			ctx.FirstVisFindIdx = len(ctx.Finds)
		}

		for _, m := range matches {
			f := find.Find{
				LineNum:      lineNum,
				StartRuneIdx: util.ByteIdxToRuneIdx(line.Buf, m[0]),
				EndRuneIdx:   util.ByteIdxToRuneIdx(line.Buf, m[1]),
			}
			ctx.Finds = append(ctx.Finds, f)
		}

		line = line.Next
		lineNum++
	}

	ctx.Modified = false

	return ctx.Finds
}
