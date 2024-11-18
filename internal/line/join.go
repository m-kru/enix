package line

import (
	"strings"
)

// Join joins line l with the next line.
// Trim indicates whether leading whiespaces from the next line shal be removed.
// Returns pointer to the new line, or nil if there was nothing to join.
func (l *Line) Join(trim bool) *Line {
	if l.Next == nil {
		return nil
	}

	l2 := l.Next
	str := string(l2.Buf)
	if trim {
		prefix := " "
		if len(l.Buf) == 0 {
			prefix = ""
		}
		str = prefix + strings.TrimLeft(str, " \t")
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

	return newLine
}
