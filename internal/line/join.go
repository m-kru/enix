package line

import (
	"strings"
)

// Join joins line l with the next line.
// Trim indicates whether leading whiespaces from the next line shal be removed.
// Returns true if lines were joined.
func (l *Line) Join(trim bool) bool {
	if l.Next == nil {
		return false
	}

	delLine := l.Next
	str := string(delLine.Buf)
	if trim {
		str = " " + strings.TrimLeft(str, " \t")
	}

	l.Append([]byte(str))

	l.Next = delLine.Next
	if delLine.Next != nil {
		delLine.Next.Prev = l
	}

	return true
}
