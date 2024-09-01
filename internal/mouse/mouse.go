package mouse

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type State int

const (
	Idle State = iota
	PrimaryClick
	DoublePrimaryClick
	TriplePrimaryClick
)

type Mouse struct {
	mtx sync.Mutex

	state     State
	PrevEvent *tcell.EventMouse
}

func (m *Mouse) TimerFunc() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	switch m.state {
	case PrimaryClick:
		panic("unimplemented")
	default:
		panic("unimplemented")
	}
}

func (m *Mouse) RxTcellEventMouse(ev *tcell.EventMouse) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	switch m.state {
	case Idle:
		m.rxEventIdle(ev)
	case PrimaryClick:
		m.rxEventPrimaryClick(ev)
	default:
		panic("unimplemented")
	}
}

func (m *Mouse) rxEventIdle(ev *tcell.EventMouse) {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		m.state = PrimaryClick
		time.AfterFunc(500*time.Millisecond, m.TimerFunc)
	default:
		// Do nothing, other mouse event
	}
}

func (m *Mouse) rxEventPrimaryClick(ev *tcell.EventMouse) {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		panic("unimplemented")
	default:
		// Do nothing, other mouse event
	}
}
