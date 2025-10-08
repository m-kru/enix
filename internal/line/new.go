package line

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/util"
)

// bufCap returns buffer capacity recommended for string of given byte length.
func bufCap(startIdx, endIdx int) int {
	diff := endIdx - startIdx
	if diff < 32 {
		return 32
	}
	return util.NextPowerOfTwo(diff)
}

func Empty() *Line {
	return &Line{Buf: make([]byte, 0, 32), Prev: nil, Next: nil}
}

func FromString(str string) (*Line, int) {
	if len(str) == 0 {
		return Empty(), 1
	}

	lineCount := 1
	startIdx := 0
	var first *Line = nil
	var prev *Line
	var next *Line
	var prevR rune

	for bIdx, r := range str {
		// Ignore carriage return after newline.
		if r == '\r' && prevR == '\n' {
			startIdx++
			prevR = r
			continue
		}

		if r == '\n' {
			endIdx := bIdx
			// Ignore carriage return before newline.
			if prevR == '\r' {
				endIdx--
			}

			if first == nil {
				first = &Line{
					Buf:  make([]byte, 0, bufCap(startIdx, endIdx)),
					Prev: nil,
					Next: nil,
				}
				first.InsertString(str[startIdx:endIdx], 0)
				prev = first
			} else {
				next = &Line{
					Buf:  make([]byte, 0, bufCap(startIdx, endIdx)),
					Prev: prev,
					Next: nil,
				}
				next.InsertString(str[startIdx:endIdx], 0)
				prev.Next = next
				prev = next
			}
			startIdx = bIdx + 1
			lineCount++
		} else if bIdx == len(str)-utf8.RuneLen(r) {
			runeLen := utf8.RuneLen(r)
			if first == nil {
				first = &Line{
					Buf:  make([]byte, 0, bufCap(startIdx, bIdx+runeLen)),
					Prev: nil,
					Next: nil,
				}
				first.InsertString(str[startIdx:bIdx+runeLen], 0)
			} else {
				next = &Line{
					Buf:  make([]byte, 0, bufCap(startIdx, bIdx+runeLen)),
					Prev: prev,
					Next: nil,
				}
				next.InsertString(str[startIdx:bIdx+runeLen], 0)
				prev.Next = next
			}
		}

		prevR = r
	}

	// Add one extra line if string ends with newline
	if str[len(str)-1] == '\n' {
		next = &Line{
			Buf:  make([]byte, 0, bufCap(0, 0)),
			Prev: prev,
			Next: nil,
		}
		prev.Next = next
	}

	return first, lineCount
}
