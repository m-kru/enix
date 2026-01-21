package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Align(r rune) error {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	var actions action.Actions

	if len(tab.Cursors) > 0 {
		actions = tab.alignCursors(r)
	} else {
		return fmt.Errorf("unimplemented for selections")
	}

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	return nil
}

func (tab *Tab) alignCursors(r rune) action.Actions {
	if len(tab.Cursors) == 1 {
		return nil
	}

	tmpCurs := cursor.LineUnique(tab.Cursors, true)
	tmpCurs = cursor.Clone(tmpCurs)
	alignCurs := make([]*cursor.Cursor, 0, len(tmpCurs))

	for _, cur := range tmpCurs {
		if r == 0 {
			alignCurs = append(alignCurs, cur)
			continue
		}

		for {
			if cur.Rune() == r {
				alignCurs = append(alignCurs, cur)
				break
			}
			cur.Right()
			if cur.RuneIdx == cur.Line.RuneCount() {
				break
			}
		}
	}

	maxCol := 0
	cols := make([]int, len(alignCurs))
	for i, cur := range alignCurs {
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

	for i, cur := range alignCurs {
		col := cols[i]
		if col == maxCol {
			continue
		}

		diff := maxCol - col

		act := cur.InsertString(maxInsert[0:diff])
		// act can't be nil here
		actions = append(actions, act)
		tab.handleAction(act)

		for _, cur2 := range alignCurs {
			if cur2 != cur {
				cur2.Inform(act)
			}
		}

		for _, tabCur := range tab.Cursors {
			tabCur.Inform(act)
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	tab.Cursors = alignCurs

	return actions
}
