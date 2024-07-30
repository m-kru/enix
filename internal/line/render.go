package line

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/view"

	"github.com/mattn/go-runewidth"
)

func (l *Line) Render(colors *cfg.Colorscheme, frame frame.Frame, view view.View) {
	frameIdx := 0
	runeIdx, runeFirstCol, ok := l.RuneIdx(view.Column)
	var r rune

	if !ok {
		goto clear
	}

	// Handle first column in a little bit different way.
	// The column might start at the second column of a rune.
	r = l.buf[runeIdx]
	if !runeFirstCol {
		r = ' '
	}
	frame.SetContent(frameIdx, 0, r, colors.Default)
	runeIdx++
	frameIdx += runewidth.RuneWidth(r)

	for {
		if runeIdx == l.Len() || frameIdx >= frame.Width-1 {
			break
		}

		r = l.buf[runeIdx]
		frame.SetContent(frameIdx, 0, r, colors.Default)

		frameIdx += runewidth.RuneWidth(r)
		runeIdx++
	}

clear:
	for frameIdx < frame.Width {
		frame.SetContent(frameIdx, 0, ' ', colors.Default)
		frameIdx++
	}
}
