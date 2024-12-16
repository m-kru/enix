package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Align() error {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions

	if len(tab.Cursors) > 0 {
		actions = tab.alignCursors()
	} else {
		return fmt.Errorf("unimplemented for selections")
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	return nil
}

func (tab *Tab) alignCursors() action.Actions {
	if len(tab.Cursors) == 1 {
		return nil
	}

	maxCol := 0
	cols := make([]int, len(tab.Cursors))
	for i, cur := range tab.Cursors {
		col := cur.Column()
		if col > maxCol {
			maxCol = col
		}
		cols[i] = col
	}

	maxDiff := 0
	for _, col := range cols {
		diff := maxCol - col
		if diff > maxDiff {
			maxDiff = diff
		}
	}

	// Cursor columns are already aligned.
	if maxDiff == 0 {
		return nil
	}

	b := strings.Builder{}
	for range maxDiff {
		b.WriteRune(' ')
	}
	maxInsert := b.String()

	actions := make(action.Actions, 0, len(tab.Cursors)-1)

	for i, cur := range tab.Cursors {
		col := cols[i]
		if col == maxCol {
			continue
		}

		diff := maxCol - col

		act := cur.InsertString(maxInsert[0:diff])
		// act can't be nil here
		actions = append(actions, act)
		tab.handleAction(act)

		for _, cur2 := range tab.Cursors {
			if cur2 != cur {
				cur2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	return actions
}
