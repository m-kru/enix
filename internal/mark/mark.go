package mark

import (
	"github.com/m-kru/enix/internal/action"
	"github.com/m-kru/enix/internal/cursor"
)

type Mark interface {
	Inform(action.Action)
}

type CursorMark struct {
	Cursors []*cursor.Cursor
}

func (cm *CursorMark) Inform(act action.Action) {
	for _, c := range cm.Cursors {
		c.Inform(act)
	}
}

func NewCursorMark(cs []*cursor.Cursor) *CursorMark {
	return &CursorMark{cursor.Clone(cs)}
}
