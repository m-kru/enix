package line

import (
	"strings"
)

// Join joins line l with the next line.
// Trim indicates whether leading whiespaces from the next line shal be removed.
// Returned line points to the deleted line.
func (l *Line) Join(trim bool) *Line {
	if l.Next == nil {
		return nil
	}

	delLine := l.Next
	str := string(delLine.Buf)
	if trim {
		str = strings.TrimLeft(str, " \t")
	}

	l.Append([]byte(str))

	l.Next = delLine.Next
	if delLine.Next != nil {
		delLine.Next.Prev = l
	}

	return delLine
}
