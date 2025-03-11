package tabbar

import (
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"
)

func (tb *TabBar) RxMouseEvent(ev mouse.Event) *tab.Tab {
	switch ev.(type) {
	case mouse.PrimaryClick, mouse.DoublePrimaryClick, mouse.TriplePrimaryClick:
		if tb.lFrame.Within(ev.X(), ev.Y()) {
			tb.viewLeft()
		} else if tb.rFrame.Within(ev.X(), ev.Y()) {
			tb.viewRight()
		} else {
			return tb.clickItemsFrame(ev.X() - tb.iFrame.X)
		}
	case mouse.WheelDown:
		tb.viewRight()
	case mouse.WheelUp:
		tb.viewLeft()
	}

	return nil
}

func (tb *TabBar) clickItemsFrame(x int) *tab.Tab {
	rIdx, _, ok := tb.line.RuneIdx(tb.view.Column + x)
	if !ok {
		return nil
	}

	for _, item := range tb.items {
		if item.StartIdx <= rIdx && rIdx < item.EndIdx {
			return item.Tab
		}
	}

	return nil
}
