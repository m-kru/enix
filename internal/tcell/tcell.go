// Package tcell contains wrappers for tcell code.
package tcell

import "github.com/gdamore/tcell/v2"

func EventKeyName(ev *tcell.EventKey) string {
	switch ev.Key() {
	case tcell.KeyCtrlH:
		return "Ctrl+H"
	}

	return ev.Name()
}
