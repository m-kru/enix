package tab

import (
	"fmt"
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/clip"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Cut() string {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	info, actions := tab.cut()

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}

	return info
}

func (tab *Tab) cut() (string, action.Actions) {
	if len(tab.Cursors) > 0 {
		return tab.cutCursors()
	} else {
		return tab.cutSelections()
	}
}

func (tab *Tab) cutCursors() (string, action.Actions) {
	curs := cursor.LineUnique(tab.Cursors, true)

	// Copy lines to the clipboard.
	b := strings.Builder{}
	for _, c := range curs {
		b.WriteString(c.Line.String())
		b.WriteString(tab.Newline)
	}
	clip.Write(b.String())

	actions := make(action.Actions, 0, len(curs))
	delLines := make([]*line.Line, 0, len(curs))

	// Delete lines.
	for _, c := range curs {
		act := c.DeleteLine()
		actions = append(actions, act)

		tab.handleAction(act)

		if ld, ok := act.(*action.LineDelete); ok {
			delLines = append(delLines, ld.Line)
		}

		for _, c2 := range curs {
			if c2 != c {
				c2.Inform(act)
			}
		}

		for _, m := range tab.Marks {
			m.Inform(act)
		}
	}

	// Create new cursors.
	newCurs := make([]*cursor.Cursor, 0, len(curs))
	for _, c := range curs {
		add := true
		for _, dl := range delLines {
			if c.Line == dl {
				add = false
				break
			}
		}
		if add {
			newCurs = append(newCurs, c)
		}
	}

	tab.Cursors = newCurs

	info := "cut line"
	if len(curs) > 1 {
		info = fmt.Sprintf("cut %d lines", len(curs))
	}

	return info, actions
}

func (tab *Tab) cutSelections() (string, action.Actions) {
	selCount := len(tab.Selections)

	tab.yankSelections()
	acts := tab.deleteSelections()

	info := "cut selection"
	if selCount > 1 {
		info = fmt.Sprintf("cut %d selections", selCount)
	}

	return info, acts
}
