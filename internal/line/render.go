package line

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/find"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/view"

	"github.com/mattn/go-runewidth"
)

func (l *Line) Render(
	cfg *cfg.Config,
	colors *cfg.Colorscheme,
	lineNum int,
	frame frame.Frame,
	view view.View,
	hls []highlight.Highlight,
	finds []find.Find,
) ([]highlight.Highlight, []find.Find) {
	currentHl := 0
	currentFind := 0
	frameX := 0
	runeIdx, runeSubcol, ok := l.RuneIdx(view.Column)

	var r rune

	setTab := func(tabSubcol int) {
		var colIdx int
		if tabSubcol == 0 {
			frame.SetContent(frameX, 0, cfg.TabRune, colors.Whitespace)
			frameX++
			colIdx = l.ColumnIdx(runeIdx) + 1
		} else {
			colIdx = l.ColumnIdx(runeIdx) + tabSubcol
		}

		for {
			if colIdx%8 == 1 || frameX >= frame.Width {
				break
			}
			frame.SetContent(frameX, 0, cfg.TabPadRune, colors.Whitespace)
			frameX++
			colIdx++
		}
	}

	if !ok {
		goto clear
	}

	// Handle first column in a little bit different way.
	// The column might start at the second column of a rune.
	r = l.Rune(runeIdx)
	if r == '\t' {
		setTab(runeSubcol)
	} else if runeSubcol > 0 {
		r = ' '
		frame.SetContent(frameX, 0, r, colors.Default)
		frameX += runewidth.RuneWidth(r)
	} else {
		color := colors.Default
		if len(hls) > 0 {
			for {
				if hls[currentHl].CoversCell(lineNum, runeIdx) {
					color = hls[currentHl].Style
					break
				}
				currentHl++
			}
		}
		if len(finds) > 0 && currentFind < len(finds) {
			if finds[currentFind].CoversCell(lineNum, runeIdx) {
				color = colors.Find
			}

			if finds[currentFind].IsLastCell(lineNum, runeIdx) {
				currentFind++
			}
		}

		frame.SetContent(frameX, 0, r, color)
		frameX += runewidth.RuneWidth(r)
	}
	runeIdx++

	for {
		if frameX >= frame.Width {
			break
		} else if runeIdx == l.RuneCount() {
			frame.SetContent(frameX, 0, cfg.LineEndRune, colors.Whitespace)
			frameX++
			break
		}

		r = l.Rune(runeIdx)

		if r == '\t' {
			setTab(0)
		} else {
			color := colors.Default
			if hls != nil {
				for {
					// TODO: We shouldn't need this check, ss there some bug?
					if currentHl >= len(hls) {
						break
					}

					if hls[currentHl].CoversCell(lineNum, runeIdx) {
						color = hls[currentHl].Style
						break
					}
					currentHl++
				}
			}

			if len(finds) > 0 && currentFind < len(finds) {
				if finds[currentFind].CoversCell(lineNum, runeIdx) {
					color = colors.Find
				}

				if finds[currentFind].IsLastCell(lineNum, runeIdx) {
					currentFind++
				}
			}

			frame.SetContent(frameX, 0, r, color)
			frameX += runewidth.RuneWidth(r)
		}
		runeIdx++
	}

clear:
	if l.RuneCount() == 0 && l.Next != nil {
		frame.SetContent(0, 0, cfg.LineEndRune, colors.Whitespace)
		frameX = 1
	}

	for frameX < frame.Width {
		frame.SetContent(frameX, 0, ' ', colors.Default)
		frameX++
	}

	return hls[currentHl:], finds[currentFind:]
}
