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
	var firstLine *Line = nil
	var line *Line
	var nextLine *Line

	for i, r := range str {
		if r == '\n' || i == len(str)-1 {
			if firstLine == nil {
				if r == '\n' {
					firstLine = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i]))}
					firstLine.Insert(str[startIdx:i], 0)
				} else {
					firstLine = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i+1]))}
					firstLine.Insert(str[startIdx:i+1], 0)
				}

				line = firstLine
				startIdx = i + 1
			} else {
				if r == '\n' {
					nextLine = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i])), Prev: line}
					nextLine.Insert(str[startIdx:i], 0)
				} else {
					nextLine = &Line{buf: make([]rune, 0, bufCap(str[startIdx:i+1])), Prev: line}
					nextLine.Insert(str[startIdx:i+1], 0)
				}
				line.Next = nextLine
				line = nextLine
				startIdx = i + 1
			}
		}
	}

	// Add one extra line if string ends with newline
	if str[len(str)-1] == '\n' {
		nextLine = &Line{buf: make([]rune, 0, bufCap("")), Prev: line}
		line.Next = nextLine
	}

	return firstLine
}
