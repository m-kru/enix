package search

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/find"
	"github.com/m-kru/enix/internal/line"
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
	for line != nil {
		matches := ctx.Regexp.FindAllIndex(line.Buf, -1)
		if len(matches) == 0 {
			line = line.Next
			lineNum++
			continue
		}

		if ctx.FirstVisFindIdx == -1 && lineNum >= ctx.FirstVisLineNum {
			ctx.FirstVisFindIdx = len(ctx.Finds)
		}

		// Matches are ordered and non-overlapping, so decode each span once.
		byteIdx, runeIdx := 0, 0
		for _, m := range matches {
			runeIdx += utf8.RuneCount(line.Buf[byteIdx:m[0]])
			startRuneIdx := runeIdx
			runeIdx += utf8.RuneCount(line.Buf[m[0]:m[1]])
			byteIdx = m[1]

			f := find.Find{
				LineNum:      lineNum,
				StartRuneIdx: startRuneIdx,
				EndRuneIdx:   runeIdx,
			}
			ctx.Finds = append(ctx.Finds, f)
		}

		line = line.Next
		lineNum++
	}

	ctx.Modified = false

	return ctx.Finds
}
