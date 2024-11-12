package sel

import (
	"github.com/m-kru/enix/internal/cursor"
)

func ToCursors(sels []*Selection) []*cursor.Cursor {
	curs := make([]*cursor.Cursor, 0, len(sels))

	for _, s := range sels {
		curs = append(curs, s.GetCursor())
	}

	return curs
}
