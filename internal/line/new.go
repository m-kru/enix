package line

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/util"
)

// bufCap returns buffer capacity recommended for string given string.
func bufCap(str string) int {
	runeCount := utf8.RuneCountInString(str)
	if runeCount < 64 {
		return 64
	}
	return util.NextPowerOfTwo(runeCount)
}

func Empty() *Line {
	return &Line{buf: make([]rune, 0, 64), Prev: nil, Next: nil}
}

func FromString(str string) *Line {
	if len(str) == 0 {
		return Empty()
	}

	startIdx := 0
	var first *Line = nil
	var prev *Line
	var next *Line

	for i, r := range str {
		if r == '\n' {
			if first == nil {
				first = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i]))}
				first.Insert(str[startIdx:i], 0)
				prev = first
			} else {
				next = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i])), Prev: prev}
				next.Insert(str[startIdx:i], 0)
				prev.Next = next
				prev = next
			}
			startIdx = i + 1
		} else if i == len(str)-utf8.RuneLen(r) {
			runeLen := utf8.RuneLen(r)
			if first == nil {
				first = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i+runeLen]))}
				first.Insert(str[startIdx:i+runeLen], 0)
			} else {
				next = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i+runeLen])), Prev: prev}
				next.Insert(str[startIdx:i+runeLen], 0)
			}
			startIdx = i + runeLen
		}
	}

	// Add one extra line if string ends with newline
	if str[len(str)-1] == '\n' {
		next = &Line{buf: make([]rune, 0, bufCap("")), Prev: prev}
		prev.Next = next
	}

	return first
}
