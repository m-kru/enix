package tabbar

import (
	"github.com/m-kru/enix/internal/mouse"
)

func RxMouseEvent(ev mouse.Event) {
	switch ev.(type) {
	case mouse.WheelDown:
		for range 2 {
			viewRight()
		}
	case mouse.WheelUp:
		for range 2 {
			viewLeft()
		}
	}
}
