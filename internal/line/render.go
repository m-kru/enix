package line

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"
)

func (l *Line) Render(colors *cfg.Colorscheme, frame frame.Frame, view view.View) {
	i := 0
	for _, r := range l.buf[view.Column-1 : len(l.buf)] {
		if i == frame.Width-1 {
			break
		}

		frame.SetContent(i, 0, r, colors.Default)
		i++
	}

	for i < frame.Width {
		frame.SetContent(i, 0, ' ', colors.Default)
		i++
	}
}
