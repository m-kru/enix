package enix

import "github.com/gdamore/tcell/v2"

type TcellEventReceiver interface {
	RxTcellEvent(ev tcell.Event) TcellEventReceiver
}
