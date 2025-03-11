package tabbar

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/tab"
)

func (tb *TabBar) Render(currentTab *tab.Tab) {
	currentItem := tb.getCurrentItem(currentTab)

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
			EndRuneIdx:   currentItem.EndIdx,
			Style:        cfg.Style.CurrentTab,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currentItem.EndIdx,
			EndRuneIdx:   tb.line.RuneCount(),
			Style:        cfg.Style.TabBar,
		},
	}

	tb.line.Render(1, tb.iFrame, tb.view, hls, nil)

	// Fill missing space
	for x := tb.line.Columns(); x < tb.iFrame.Width; x++ {
		tb.iFrame.SetContent(x, 0, ' ', cfg.Style.TabBar)
	}

	tb.renderLeftArrow()
	tb.renderRightArrow()
}

func (tb *TabBar) renderLeftArrow() {
	r := ' '
	if tb.view.Column > 1 {
		r = '<'
	}
	tb.lFrame.SetContent(0, 0, r, cfg.Style.TabBar)
	tb.lFrame.SetContent(1, 0, ' ', cfg.Style.TabBar)
}

func (tb *TabBar) renderRightArrow() {
	tb.rFrame.SetContent(0, 0, ' ', cfg.Style.TabBar)
	r := ' '
	if tb.view.LastColumn() < tb.line.Columns() {
		r = '>'
	}
	tb.rFrame.SetContent(1, 0, r, cfg.Style.TabBar)
}
