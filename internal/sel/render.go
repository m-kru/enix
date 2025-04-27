package sel

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (s *Selection) Render(
	frame frame.Frame, // Tab frame
	view view.View,
) {
	for s != nil {
		sv := s.LineView()
		if !view.IsVisible(sv) {
			s = s.Next
			continue
		}

		iv := view.Intersection(sv)

		lastColNum := s.LastColumnNumber()

		for c := iv.Column; c <= iv.LastColumn(); c++ {
			// + 1 because of the newline character rendering.
			if c > lastColNum+1 {
				break
			}

			x := c - view.Column
			y := iv.Line - view.Line
			r := frame.GetContent(x, y)

			if c == sv.Column && s.CursorOnLeft() || c == sv.LastColumn() && s.CursorOnRight() {
				frame.SetContent(x, y, r, cfg.Style.Cursor)
			} else {
				frame.SetContent(x, y, r, cfg.Style.Selection)
			}
		}

		s = s.Next
	}
}
