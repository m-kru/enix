package mouse

import (
	"github.com/gdamore/tcell/v2"
)

type State int

const (
	idle State = iota
	primaryClick
	primaryClickAlt
	primaryClickCtrl
	doublePrimaryClick
	triplePrimaryClick
	middleClick
)

var state State
var prevEv *tcell.EventMouse

func RxTcellEventMouse(ev *tcell.EventMouse) Event {
	if prevEv != nil {
		if ev.When().UnixMilli()-prevEv.When().UnixMilli() > 500 {
			state = idle
		}
	}

	switch state {
	case idle:
		return rxEventIdle(ev)
	case primaryClick:
		return rxEventPrimaryClick(ev)
	case primaryClickAlt:
		return rxEventPrimaryClickAlt(ev)
	case primaryClickCtrl:
		return rxEventPrimaryClickCtrl(ev)
	case doublePrimaryClick:
		return rxEventDoublePrimaryClick(ev)
	case middleClick:
		return rxEventMiddleClick(ev)
	}

	return nil
}

func rxEventIdle(ev *tcell.EventMouse) Event {
	prevEv = ev
	x, y := ev.Position()

	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.ButtonPrimary:
		if ev.Modifiers()&tcell.ModCtrl != 0 {
			state = primaryClickCtrl
			return PrimaryClickCtrl{x, y}
		} else if ev.Modifiers()&tcell.ModAlt != 0 {
			state = primaryClickAlt
			return PrimaryClickAlt{x, y}
		} else if ev.Modifiers() == 0 {
			state = primaryClick
			return PrimaryClick{x, y}
		}
	case tcell.ButtonMiddle:
		if ev.Modifiers()&tcell.ModCtrl != 0 {
			return MiddleClickCtrl{x, y}
		} else if ev.Modifiers() == 0 {
			state = middleClick
			return MiddleClick{x, y}
		}
	case tcell.WheelDown, tcell.WheelUp:
		return rxScrollEvent(ev)
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func rxEventPrimaryClick(ev *tcell.EventMouse) Event {
	prevX, prevY := prevEv.Position()

	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.ButtonPrimary:
		x, y := ev.Position()
		prevEv = ev
		if x == prevX && y == prevY {
			state = doublePrimaryClick
			if ev.Modifiers() == 0 {
				return DoublePrimaryClick{x, y}
			}
		} else {
			return PrimaryClick{x, y}
		}
	case tcell.WheelDown, tcell.WheelUp:
		return rxScrollEvent(ev)
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func rxEventDoublePrimaryClick(ev *tcell.EventMouse) Event {
	prevX, prevY := prevEv.Position()

	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.ButtonPrimary:
		x, y := ev.Position()
		prevEv = ev
		if x == prevX && y == prevY {
			// Go back to idle state.
			// Triple primary click is an ultimate event.
			state = idle
			return TriplePrimaryClick{x, y}
		} else {
			return PrimaryClick{x, y}
		}
	case tcell.WheelDown, tcell.WheelUp:
		return rxScrollEvent(ev)
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func rxEventPrimaryClickAlt(ev *tcell.EventMouse) Event {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.ButtonPrimary:
		prevEv = ev
		x, y := ev.Position()

		if ev.Modifiers()&tcell.ModCtrl != 0 {
			state = primaryClickCtrl
			return PrimaryClickCtrl{x, y}
		} else if ev.Modifiers()&tcell.ModAlt != 0 {
			state = primaryClickAlt
			return PrimaryClickAlt{x, y}
		} else {
			state = primaryClick
			return PrimaryClick{x, y}
		}
	default:
		// Do nothing, other mouse event
	}

	return nil
}

func rxEventPrimaryClickCtrl(ev *tcell.EventMouse) Event {
	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.ButtonPrimary:
		prevEv = ev
		x, y := ev.Position()

		if ev.Modifiers()&tcell.ModCtrl != 0 {
			state = primaryClickCtrl
			return PrimaryClickCtrl{x, y}
		} else if ev.Modifiers()&tcell.ModAlt != 0 {
			state = primaryClickAlt
			return PrimaryClickAlt{x, y}
		} else {
			state = primaryClick
			return PrimaryClick{x, y}
		}
	default:
		// Do nothing, other mouse event
	}

	return nil
}

// Scrolls events are handled in the same way in all states.
func rxScrollEvent(ev *tcell.EventMouse) Event {
	x, y := ev.Position()

	switch ev.Buttons() {
	case tcell.WheelDown:
		state = idle
		if ev.Modifiers()&tcell.ModShift != 0 {
			return WheelRight{x, y}
		} else {
			return WheelDown{x, y}
		}
	case tcell.WheelUp:
		state = idle
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

func rxEventMiddleClick(ev *tcell.EventMouse) Event {
	x, y := ev.Position()

	switch ev.Buttons() {
	case tcell.ButtonNone:
		// Do nothing, just mouse movement.
	case tcell.ButtonPrimary:
		if ev.Modifiers()&tcell.ModCtrl != 0 {
			state = primaryClickCtrl
			return PrimaryClickCtrl{x, y}
		} else if ev.Modifiers()&tcell.ModAlt != 0 {
			state = primaryClickAlt
			return PrimaryClickAlt{x, y}
		} else if ev.Modifiers() == 0 {
			state = primaryClick
			return PrimaryClick{x, y}
		}
	case tcell.ButtonMiddle:
		prevX, prevY := prevEv.Position()
		x, y := ev.Position()
		prevEv = ev
		if x == prevX && y == prevY {
			state = idle
			if ev.Modifiers() == 0 {
				return DoubleMiddleClick{x, y}
			}
		} else {
			return MiddleClick{x, y}
		}
	case tcell.WheelDown, tcell.WheelUp:
		return rxScrollEvent(ev)
	default:
		// Do nothing, other mouse event
	}

	return nil
}
