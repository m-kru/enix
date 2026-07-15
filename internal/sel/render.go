package sel

import (
	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (s *Selection) Render(
	frame frame.Frame, // Tab frame
	view view.View,
	selStyle tcell.Style,
	curStyle tcell.Style,
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
				frame.SetContent(x, y, r, curStyle)
			} else {
				frame.SetContent(x, y, r, selStyle)
			}
		}

		s = s.Next
	}
}
