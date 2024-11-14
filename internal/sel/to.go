package sel

import (
	"strings"

	"github.com/m-kru/enix/internal/cursor"
)

func ToCursors(sels []*Selection) []*cursor.Cursor {
	curs := make([]*cursor.Cursor, 0, len(sels))

	for _, s := range sels {
		curs = append(curs, s.GetCursor())
	}

	return curs
}

func (s *Selection) ToString() string {
	b := strings.Builder{}

	for {
		sIdx := s.Line.BufIdx(s.StartRuneIdx)
		eIdx := s.Line.BufIdx(s.EndRuneIdx + 1)
		b.Write(s.Line.Buf[sIdx:eIdx])
		// Below conditions are reduntant. However, the first check makes it
		// faster for selections consisting of multiple selections as the first
		// check is O(1), and the second one is O(n).
		if s.Next != nil || s.EndRuneIdx == s.Line.RuneCount() {
			b.WriteRune('\n')
		}

		s = s.Next
		if s == nil {
			break
		}
	}

	return b.String()
}

func ToString(sels []*Selection) string {
	b := strings.Builder{}

	for i, s := range sels {
		b.WriteString(s.ToString())
		if i < len(sels)-1 {
			b.WriteRune('\n')
		}
	}

	return b.String()
}
