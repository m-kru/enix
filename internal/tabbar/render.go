package tabbar

import (
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/tab"
)

func Render(currentTab *tab.Tab) {
	b := strings.Builder{}

	cTabStartIdx := 0
	cTabEndIdx := 0

	for x := range items {
		t := items[x].Tab

		items[x].StartIdx = b.Len()
		if t == currentTab {
			cTabStartIdx = b.Len()
		}

		b.WriteRune(' ')
		if t.HasChanges() {
			b.WriteRune('*')
		}
		b.WriteString(items[x].Name)
		b.WriteRune(' ')

		items[x].EndIdx = b.Len()
		if t == currentTab {
			cTabEndIdx = b.Len()
		}
	}

	for i, r := range b.String() {
		style := cfg.Colors.TabBar
		if cTabStartIdx <= i && i < cTabEndIdx {
			style = cfg.Colors.CurrentTab
		}
		iFrame.SetContent(i, 0, r, style)
	}

	// Clear remaining cells
	for x := b.Len(); x < iFrame.Width; x++ {
		iFrame.SetContent(x, 0, ' ', cfg.Colors.TabBar)
	}

	renderLeftArrow()
	renderRightArrow()
}

func renderLeftArrow() {
	lFrame.SetContent(0, 0, ' ', cfg.Colors.TabBar)
	lFrame.SetContent(1, 0, ' ', cfg.Colors.TabBar)
}

func renderRightArrow() {
	rFrame.SetContent(0, 0, ' ', cfg.Colors.TabBar)
	rFrame.SetContent(1, 0, ' ', cfg.Colors.TabBar)
}
