package tab

import (
	"strings"

	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/clip"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/sel"
)

func (tab *Tab) Cut() {
	prevCurs := cursor.Clone(tab.Cursors)
	prevSels := sel.Clone(tab.Selections)

	actions := tab.cut()

	if len(actions) > 0 {
		tab.undoPush(actions.Reverse(), prevCurs, prevSels)
	}
}

func (tab *Tab) cut() action.Actions {
	if len(tab.Cursors) > 0 {
		return tab.cutCursors()
	} else {
		return tab.cutSelections()
	}
}

func (tab *Tab) cutCursors() action.Actions {
	curs := cursor.Uniques(tab.Cursors, true)

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

	return actions
}

func (tab *Tab) cutSelections() action.Actions {
	return nil
}
