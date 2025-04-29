package cursor

import (
	"unicode/utf8"
)

// lowerBound and upperBound and lower and upper line number bounds.
// If bounds should not be applied, provide 0 value.
//
// The bounds are useful for working with large files, when highlighting matching
// brackets out of the screen view makes no sense.
// If some bracket didn't have a matching bracket, then the match functions would
// scan the whole tab content, what would create a noticeable lag.
//
// The same documentation applies to the MatchCurly and MatchParen methods.
func (cur *Cursor) MatchBracket(lowerBound int, upperBound int) *Cursor {
	r := cur.Line.Rune(cur.RuneIdx)
	if r == ']' {
		return cur.matchLeft('[', ']', lowerBound)
	}
	return cur.matchRight(']', '[', upperBound)
}

func (cur *Cursor) MatchCurly(lowerBound int, upperBound int) *Cursor {
	r := cur.Line.Rune(cur.RuneIdx)
	if r == '}' {
		return cur.matchLeft('{', '}', lowerBound)
	}
	return cur.matchRight('}', '{', upperBound)
}

func (cur *Cursor) MatchParen(lowerBound int, upperBound int) *Cursor {
	r := cur.Line.Rune(cur.RuneIdx)
	if r == ')' {
		return cur.matchLeft('(', ')', lowerBound)
	}
	return cur.matchRight(')', '(', upperBound)
}

// Match rune r going left, cr is the counter rune.
//
// matchLeft performance is worse than matchRight, because line buffers are
// traversed in the backward direction.
func (cur *Cursor) matchLeft(r rune, cr rune, lowerBound int) *Cursor {
	cur = cur.Clone()

	// Counter of encountered counter runes.
	cnt := 0

	// Previous line and rune index
	line := cur.Line
	rIdx := cur.RuneIdx

	for {
		cur.Left()
		if cur.LineNum < lowerBound || (cur.Line == line && cur.RuneIdx == rIdx) {
			return nil
		}

		line = cur.Line
		rIdx = cur.RuneIdx

		curR := cur.Line.Rune(cur.RuneIdx)
		if curR == r {
			if cnt == 0 {
				break
			}
			cnt--
		} else if curR == cr {
			cnt++
		}
	}

	return cur
}

// Match rune r going right, cr is the counter rune.
//
// matchRight is optimized compared to the matchLeft because when going right
// the rune length of the current rune is required.
func (cur *Cursor) matchRight(r rune, cr rune, upperBound int) *Cursor {
	// Counter of encountered counter runes.
	cnt := 0

	line := cur.Line
	lineNum := cur.LineNum
	rIdx := cur.RuneIdx
	bIdx := line.BufIdx(rIdx)

	var curR rune
	rLen := 0

	for {
		if upperBound != 0 && lineNum > upperBound {
			return nil
		}

		if bIdx >= len(line.Buf) {
			if line.Next == nil {
				return nil
			}
			line = line.Next
			lineNum++
			if len(line.Buf) == 0 {
				continue
			}
			bIdx = 0
			rIdx = 0
		} else {
			// Check the start condition.
			// We must move to the right.
			if rLen == 0 {
				_, rLen = utf8.DecodeRune(line.Buf[bIdx:])
			}
			bIdx += rLen
			rIdx++
		}

		curR, rLen = utf8.DecodeRune(line.Buf[bIdx:])
		if curR == r {
			if cnt == 0 {
				break
			}
			cnt--
		} else if curR == cr {
			cnt++
		}
	}

	return New(line, lineNum, rIdx)
}
