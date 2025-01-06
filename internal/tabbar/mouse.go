package tabbar

import (
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"
)

func RxMouseEvent(ev mouse.Event) *tab.Tab {
	switch ev.(type) {
	case mouse.PrimaryClick, mouse.DoublePrimaryClick, mouse.TriplePrimaryClick:
		if lFrame.Within(ev.X(), ev.Y()) {
			viewLeft()
		} else if rFrame.Within(ev.X(), ev.Y()) {
			viewRight()
		} else {
			return clickItemsFrame(ev.X() - iFrame.X)
		}
	case mouse.WheelDown:
		viewRight()
	case mouse.WheelUp:
		viewLeft()
	}

	return nil
}

func clickItemsFrame(x int) *tab.Tab {
	rIdx, _, ok := line.RuneIdx(view.Column + x)
	if !ok {
		return nil
	}

	for _, item := range items {
		if item.StartIdx <= rIdx && rIdx < item.EndIdx {
			return item.Tab
		}
	}

	return nil
}
