package cursor

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (c *Cursor) Render(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	frame frame.Frame,
	view view.View,
	primary bool,
) {
	x := c.Line.ColumnIdx(c.RuneIdx, config.TabWidth) - view.Column
	/*
		if x >= frame.Width {
			return
		}
	*/
	if primary {
		frame.ShowCursor(x, 0)
	} else {
		r := frame.GetContent(x, 0)
		frame.SetContent(x, 0, r, colors.Cursor)
	}
}
