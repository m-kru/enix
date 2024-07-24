package enix

import "github.com/gdamore/tcell/v2"

type EventReceiver interface {
	RxEvent(ev tcell.Event) EventReceiver
}
