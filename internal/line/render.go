package line

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/find"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
	"github.com/mattn/go-runewidth"
)

func (l *Line) Render(
	lineNum int,
	frame frame.Frame,
	view view.View,
	hls []highlight.Highlight,
	finds []find.Find,
) ([]highlight.Highlight, []find.Find) {
	hlIdx := 0   // Current highlight index
	findIdx := 0 // Current find index
	x := 0       // Frame x coordinate
	rIdx, runeSubcol, ok := l.RuneIdx(view.Column)

	var r rune

	setTab := func(tabSubcol int, style tcell.Style) {
		var colIdx int
		if tabSubcol == 0 {
			frame.SetContent(x, 0, cfg.Cfg.TabRune, style)
			x++
			colIdx = l.ColumnIdx(rIdx) + 1
		} else {
			colIdx = l.ColumnIdx(rIdx) + tabSubcol
		}

		for {
			if colIdx%8 == 1 || x >= frame.Width {
				break
			}
			frame.SetContent(x, 0, cfg.Cfg.TabPadRune, style)
			x++
			colIdx++
		}
	}

	style := cfg.Style.Default
	setStyle := func() {
		style = cfg.Style.Default
		if hls != nil {
			for {
				// TODO: We shouldn't need this check, is there some bug?
				if hlIdx >= len(hls) {
					break
				}

				if hls[hlIdx].CoversCell(lineNum, rIdx) {
					style = hls[hlIdx].Style
					break
				}
				hlIdx++
			}
		}
		if len(finds) > 0 && findIdx < len(finds) {
			if finds[findIdx].CoversRune(lineNum, rIdx) {
				style = cfg.Style.Find
				if finds[findIdx].IsLastRune(rIdx) {
					findIdx++
				}
			}
		}
	}

	if !ok {
		goto clear
	}

	setStyle()

	// Handle first column in a little bit different way.
	// The column might start at the second column of a rune.
	r = l.Rune(rIdx)
	if r == '\t' {
		if style != cfg.Style.Find {
			style = cfg.Style.Whitespace
		}
		setTab(runeSubcol, style)
		rIdx++
	} else if runeSubcol > 0 {
		r = ' '
		frame.SetContent(x, 0, r, style)
		x += runewidth.RuneWidth(r)
		rIdx++
	}

	for {
		if rIdx == l.RuneCount() || x >= frame.Width {
			break
		}

		r = l.Rune(rIdx)

		setStyle()

		if r == '\t' {
			if style != cfg.Style.Find {
				style = cfg.Style.Whitespace
			}
			setTab(0, style)
		} else {
			frame.SetContent(x, 0, r, style)
			x += runewidth.RuneWidth(r)
		}
		rIdx++
	}

clear:
	frame.SetContent(x, 0, cfg.Cfg.LineEndRune, cfg.Style.Whitespace)
	x++

	for x < frame.Width {
		frame.SetContent(x, 0, ' ', cfg.Style.Default)
		x++
	}

	return hls[hlIdx:], finds[findIdx:]
}
