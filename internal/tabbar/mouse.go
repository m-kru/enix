package tabbar

import (
	"github.com/m-kru/enix/internal/mouse"
)

func RxMouseEvent(ev mouse.Event) {
	switch ev.(type) {
	case mouse.PrimaryClick, mouse.DoublePrimaryClick, mouse.TriplePrimaryClick:
		if lFrame.Within(ev.X(), ev.Y()) {
			viewLeft()
		} else if rFrame.Within(ev.X(), ev.Y()) {
			viewRight()
		}
	case mouse.WheelDown:
		viewRight()
	case mouse.WheelUp:
		viewLeft()
	}
}
