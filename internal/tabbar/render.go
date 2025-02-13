package tabbar

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/tab"
)

func Render(currentTab *tab.Tab) {
	currentItem := getCurrentItem(currentTab)

	hls := []highlight.Highlight{
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: 0,
			EndRuneIdx:   currentItem.StartIdx,
			Style:        cfg.Style.TabBar,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currentItem.StartIdx,
			EndRuneIdx:   currentItem.EndIdx + 1,
			Style:        cfg.Style.CurrentTab,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currentItem.EndIdx + 1,
			EndRuneIdx:   line.RuneCount(),
			Style:        cfg.Style.TabBar,
		},
	}

	line.Render(1, iFrame, view, hls, nil)

	// Fill missing space
	for x := line.Columns(); x < iFrame.Width; x++ {
		iFrame.SetContent(x, 0, ' ', cfg.Style.TabBar)
	}

	renderLeftArrow()
	renderRightArrow()
}

func renderLeftArrow() {
	r := ' '
	if view.Column > 1 {
		r = '<'
	}
	lFrame.SetContent(0, 0, r, cfg.Style.TabBar)
	lFrame.SetContent(1, 0, ' ', cfg.Style.TabBar)
}

func renderRightArrow() {
	rFrame.SetContent(0, 0, ' ', cfg.Style.TabBar)
	r := ' '
	if view.LastColumn() < line.Columns() {
		r = '>'
	}
	rFrame.SetContent(1, 0, r, cfg.Style.TabBar)
}
