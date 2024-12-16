package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/clip"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
)

func (tab *Tab) Paste() error {
	if len(tab.Cursors) > 0 {
		tab.pasteCursors()
	} else {
		return fmt.Errorf("unimplemented for selections")
	}

	return nil
}

func (tab *Tab) pasteCursors() {
	text := clip.Read()
	if len(text) == 0 {
		return
	}

	prevCurs := cursor.Clone(tab.Cursors)

	var actions action.Actions
	if strings.HasSuffix(text, "\n") {
		actions = tab.pasteCursorsLineBased(text)
	} else {
		panic("unimplemented")
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, nil)
	}
}

func (tab *Tab) pasteCursorsLineBased(text string) action.Actions {
	actions := make(action.Actions, 0, len(tab.Cursors))
	lines, lineCount := line.FromString(text[0 : len(text)-1])

	curs := cursor.Uniques(tab.Cursors, true)
	newCurs := make([]*cursor.Cursor, 0, len(curs))

	for _, cur := range curs {
		acts := make(action.Actions, 0, lineCount)

		line := lines
		for {
			if line == nil {
				break
			}

			act := cur.InsertLineBelow(line.String())
			acts = append(acts, act)

			cur.Down()

			for _, cur2 := range curs {
				if cur2 != cur {
					cur2.Inform(act)
				}
			}

			for _, m := range tab.Marks {
				m.Inform(act)
			}

			line = line.Next
		}

		actions = append(actions, acts)

		cur.LineEnd()
		newCurs = append(newCurs, cur)
	}

	tab.Cursors = newCurs

	return actions
}
