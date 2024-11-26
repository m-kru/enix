package cursor

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (c *Cursor) Render(
	colors *cfg.Colorscheme,
	frame frame.Frame,
	view view.View,
) {
	x := c.Line.ColumnIdx(c.RuneIdx) - view.Column

	r := frame.GetContent(x, 0)
	frame.SetContent(x, 0, r, colors.Cursor)
}
