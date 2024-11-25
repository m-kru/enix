package line

import (
	"strings"
	"unicode/utf8"
)

// Join joins line l with the next line.
// Trim indicates whether leading whiespaces from the next line shal be removed.
// Returns pointer to the new line, or nil if there was nothing to join.
// The second return is the number of trimmed runes.
func (l *Line) Join(trim bool) (*Line, int) {
	if l.Next == nil {
		return nil, 0
	}

	l2 := l.Next
	str := string(l2.Buf)
	trimmedCount := utf8.RuneCountInString(str) - 1 // - 1 becase we potentailly add one extra space ' '
	if trim {
		prefix := " "
		if len(l.Buf) == 0 {
			prefix = ""
			trimmedCount++ // Increment, there is not extra ' '
		}
		str = strings.TrimLeft(str, " \t")
		trimmedCount -= utf8.RuneCountInString(str)
		str = prefix + str
	}

	newLine, _ := FromString(l.String() + str)

	newLine.Prev = l.Prev
	newLine.Next = l2.Next
	if l.Prev != nil {
		l.Prev.Next = newLine
	}
	if l2.Next != nil {
		l2.Next.Prev = newLine
	}

	return newLine, trimmedCount
}
