package mouse

import (
	"github.com/gdamore/tcell/v2"
)

type State int

const (
	idle State = iota
	primaryClick
	doublePrimaryClick
	triplePrimaryClick
	primaryClickCtrl
)

type Mouse struct {
	state  State
	prevEv *tcell.EventMouse
}

func (m *Mouse) RxTcellEventMouse(ev *tcell.EventMouse) Event {
	if m.prevEv != nil {
		if ev.When().UnixMilli()-m.prevEv.When().UnixMilli() > 500 {
			m.state = idle
		}
	}

	switch m.state {
	case idle:
		return m.rxEventIdle(ev)
	case primaryClick:
		return m.rxEventPrimaryClick(ev)
	case doublePrimaryClick:
		return m.rxEventDoublePrimaryClick(ev)
	case primaryClickCtrl:
		return m.rxEventPrimaryClickCtrl(ev)
	}

	return nil
}

func (m *Mouse) rxEventIdle(ev *tcell.EventMouse) Event {
	m.prevEv = ev
	x, y := ev.Position()

	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		if ev.Modifiers()&tcell.ModCtrl != 0 {
			m.state = primaryClickCtrl
			return PrimaryClickCtrl{x, y}
		} else {
			m.state = primaryClick
			return PrimaryClick{x, y}
		}
	case tcell.WheelDown:
		m.state = idle
		if ev.Modifiers()&tcell.ModShift != 0 {
			return WheelRight{x, y}
		} else {
			return WheelDown{x, y}
		}
	case tcell.WheelUp:
		m.state = idle
		if ev.Modifiers()&tcell.ModShift != 0 {
			return WheelLeft{x, y}
		} else {
			return WheelUp{x, y}
		}
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func (m *Mouse) rxEventPrimaryClick(ev *tcell.EventMouse) Event {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		x, y := m.prevEv.Position()
		x2, y2 := ev.Position()
		m.prevEv = ev
		if x == x2 && y == y2 {
			m.state = doublePrimaryClick
			return DoublePrimaryClick{x, y}
		} else {
			return PrimaryClick{x2, y2}
		}
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func (m *Mouse) rxEventDoublePrimaryClick(ev *tcell.EventMouse) Event {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		// Implement TriplePrimaryClick event handling here.
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func (m *Mouse) rxEventPrimaryClickCtrl(ev *tcell.EventMouse) Event {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.Button1:
		m.prevEv = ev
		x, y := ev.Position()

		if ev.Modifiers()&tcell.ModCtrl != 0 {
			m.state = primaryClickCtrl
			return PrimaryClickCtrl{x, y}
		} else {
			m.state = primaryClick
			return PrimaryClick{x, y}
		}
	default:
		// Do nothing, other mouse event
	}

	return nil
}
