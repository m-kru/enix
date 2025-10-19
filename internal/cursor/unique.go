package cursor

import "sort"

// LineUnique returns line unique cursors.
// Line unique cursors are cursors pointing to different lines.
// The ascending parameter controls the sort order.
func LineUnique(curs []*Cursor, ascending bool) []*Cursor {
	uniques := make([]*Cursor, 0, len(curs))

	for _, c := range curs {
		found := false
		for _, c2 := range uniques {
			if c.Line == c2.Line {
				found = true
				break
			}
		}

		if found {
			continue
		}

		uniques = append(uniques, c)
	}

	softFunc := func(i, j int) bool {
		if ascending {
			return uniques[i].LineNum < uniques[j].LineNum
		} else {
			return uniques[i].LineNum > uniques[j].LineNum
		}
	}

	sort.Slice(uniques, softFunc)

	return uniques
}
