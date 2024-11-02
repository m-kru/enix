package line

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/view"

	"github.com/mattn/go-runewidth"
)

func (l *Line) Render(
	cfg *cfg.Config,
	colors *cfg.Colorscheme,
	frame frame.Frame,
	view view.View,
	hls []highlight.Highlight,
) []highlight.Highlight {
	consumedHls := 0
	frameIdx := 0
	runeIdx, runeSubcol, ok := l.RuneIdx(view.Column, cfg.TabWidth)
	var r rune

	setTab := func(tabSubcol int) {
		var colIdx int
		if tabSubcol == 0 {
			frame.SetContent(frameIdx, 0, cfg.TabRune, colors.Whitespace)
			frameIdx++
			colIdx = l.ColumnIdx(runeIdx, cfg.TabWidth) + 1
		} else {
			colIdx = l.ColumnIdx(runeIdx, cfg.TabWidth) + tabSubcol
		}

		if cfg.TabWidth == 1 {
			return
		}

		for {
			if colIdx%cfg.TabWidth == 1 || frameIdx >= frame.Width {
				break
			}
			frame.SetContent(frameIdx, 0, cfg.TabPadRune, colors.Whitespace)
			frameIdx++
			colIdx++
		}
	}

	if !ok {
		goto clear
	}

	// Handle first column in a little bit different way.
	// The column might start at the second column of a rune.
	r = l.Buf[runeIdx]
	if r == '\t' {
		setTab(runeSubcol)
	} else if runeSubcol > 0 {
		r = ' '
		frame.SetContent(frameIdx, 0, r, colors.Default)
		frameIdx += runewidth.RuneWidth(r)
	} else {
		color := colors.Default
		if len(hls) > 0 {
			for {
				if hls[consumedHls].CoversCell(l.Num(), runeIdx) {
					color = hls[consumedHls].Style
					break
				}
				consumedHls++
			}
		}

		frame.SetContent(frameIdx, 0, r, color)
		frameIdx += runewidth.RuneWidth(r)
	}
	runeIdx++

	for {
		if frameIdx >= frame.Width {
			break
		} else if runeIdx == l.RuneCount() {
			frame.SetContent(frameIdx, 0, cfg.NewlineRune, colors.Whitespace)
			frameIdx++
			break
		}

		r = l.Buf[runeIdx]

		if r == '\t' {
			setTab(0)
		} else {
			color := colors.Default
			if hls != nil {
				for {
					// TODO: We shouldn't need this check, there is more bug.
					if consumedHls >= len(hls) {
						break
					}

					if hls[consumedHls].CoversCell(l.Num(), runeIdx) {
						color = hls[consumedHls].Style
						break
					}
					consumedHls++
				}
			}

			frame.SetContent(frameIdx, 0, r, color)
			frameIdx += runewidth.RuneWidth(r)
		}
		runeIdx++
	}

clear:
	if l.RuneCount() == 0 && l.Next != nil {
		frame.SetContent(0, 0, cfg.NewlineRune, colors.Whitespace)
		frameIdx = 1
	}

	for frameIdx < frame.Width {
		frame.SetContent(frameIdx, 0, ' ', colors.Default)
		frameIdx++
	}

	return hls[consumedHls:]
}
