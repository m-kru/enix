package tabbar

import (
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/tab"
	"github.com/m-kru/enix/internal/view"
)

type TabBar struct {
	View  view.View
	items []item
}

func (tb TabBar) Render(
	tabs *tab.Tab,
	currentTab *tab.Tab,
	colors *cfg.Colorscheme,
	frame frame.Frame,
) {
	tb.items = createItems(tabs)

	b := strings.Builder{}

	cTabStartIdx := 0
	cTabEndIdx := 0

	for x := range tb.items {
		t := tb.items[x].Tab

		tb.items[x].StartIdx = b.Len()
		if t == currentTab {
			cTabStartIdx = b.Len()
		}

		b.WriteRune(' ')
		if t.HasChanges {
			b.WriteRune('*')
		}
		b.WriteString(tb.items[x].Name)
		b.WriteRune(' ')

		tb.items[x].EndIdx = b.Len()
		if t == currentTab {
			cTabEndIdx = b.Len()
		}
	}

	for i, r := range b.String() {
		style := colors.TabBar
		if cTabStartIdx <= i && i < cTabEndIdx {
			style = colors.CurrentTab
		}
		frame.SetContent(i, 0, r, style)
	}

	// Clear remaining cells
	for x := b.Len(); x < frame.Width; x++ {
		frame.SetContent(x, 0, ' ', colors.TabBar)
	}
}
