package sel

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (s *Selection) Render(
	colors *cfg.Colorscheme,
	frame frame.Frame, // Tab frame
	view view.View,
) {
	for {
		if s == nil {
			break
		}

		sv := s.View()
		if !view.IsVisible(sv) {
			s = s.Next
			continue
		}

		iv := view.Intersection(sv)

		for c := iv.Column; c <= iv.LastColumn(); c++ {
			x := c - view.Column
			y := iv.Line - view.Line
			r := frame.GetContent(x, y)

			if c == sv.Column && s.CursorOnLeft() || c == sv.LastColumn() && s.CursorOnRight() {
				frame.SetContent(x, y, r, colors.Cursor)
			} else {
				frame.SetContent(x, y, r, colors.Selection)
			}
		}

		s = s.Next
	}
}
