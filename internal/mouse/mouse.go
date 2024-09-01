package mouse

import (
	"sync"
	"time"

	"github.com/gdamore/tcell/v2"
)

type State int

const (
	idle State = iota
	primaryClick
	doublePrimaryClick
	triplePrimaryClick
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
	case primaryClick:
		panic("unimplemented")
	default:
		panic("unimplemented")
	}
}

func (m *Mouse) RxTcellEventMouse(ev *tcell.EventMouse) {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	switch m.state {
	case idle:
		m.rxEventIdle(ev)
	case primaryClick:
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
		m.state = primaryClick
		m.PrevEvent = ev
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
