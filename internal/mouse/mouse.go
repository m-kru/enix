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
	prevEvent *tcell.EventMouse

	EventChan chan Event
}

func (m *Mouse) TimerFunc() {
	m.mtx.Lock()
	defer m.mtx.Unlock()

	switch m.state {
	case primaryClick:
		m.state = idle
	case doublePrimaryClick:
		m.state = idle
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
	case doublePrimaryClick:
		m.rxEventDoublePrimaryClick(ev)
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
		m.prevEvent = ev
		x, y := ev.Position()
		m.EventChan <- PrimaryClick{x: x, y: y}
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
		x, y := m.prevEvent.Position()
		x2, y2 := ev.Position()
		m.prevEvent = ev
		if x == x2 && y == y2 {
			m.state = doublePrimaryClick
			m.EventChan <- DoublePrimaryClick{x: x, y: y}
		} else {
			m.EventChan <- PrimaryClick{x: x2, y: y2}
		}
	default:
		// Do nothing, other mouse event
	}
}

func (m *Mouse) rxEventDoublePrimaryClick(ev *tcell.EventMouse) {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		// Implement TriplePrimaryClick event handling here.
	default:
		// Do nothing, other mouse event
	}
}
