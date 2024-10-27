package enix

import (
	"strings"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/tab"
	"github.com/m-kru/enix/internal/view"
)

type TabBar struct {
	View view.View
}

func (tb TabBar) Render(
	tabs *tab.Tab,
	currentTab *tab.Tab,
	colors *cfg.Colorscheme,
	frame frame.Frame,
) {
	b := strings.Builder{}

	cTabStartIdx := 0
	cTabEndIdx := 0

	t := tabs

	for {
		if t == nil {
			break
		}

		if t == currentTab {
			cTabStartIdx = b.Len()
		}

		b.WriteRune(' ')
		if t.HasChanges {
			b.WriteRune('*')
		}
		b.WriteString(t.Path)
		b.WriteRune(' ')

		if t == currentTab {
			cTabEndIdx = b.Len()
		}

		t = t.Next
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
