package cursor

import (
	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (c *Cursor) Render(frame frame.Frame, view view.View, style tcell.Style) {
	x := c.Line.ColumnIdx(c.RuneIdx) - view.Column

	r := frame.GetContent(x, 0)
	frame.SetContent(x, 0, r, style)
}
