package cfg

import (
	"github.com/gdamore/tcell/v2"
	enix_tcell "github.com/m-kru/enix/internal/tcell"
)

// Keybindings struct represents configured keybindings.
type Keybindings map[string]string

// KeybindingsDefault returns default keybindings.
func KeybindingsDefault() Keybindings {
	return map[string]string{
		// Cmd
		"Ctrl+E": "cmd",
		// Cursor
		"Down":       "cursor-down",
		"Left":       "cursor-left",
		"Right":      "cursor-right",
		"Up":         "cursor-up",
		"Ctrl+Left":  "cursor-word-start",
		"Ctrl+Right": "cursor-word-end",
		// File
		"Ctrl+O": "file-open",
		"Ctrl+S": "file-save",
		// Tab
		"Ctrl+T": "tab-open",
		// Miscellaneous
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Esc":        "escape",
		"Ctrl+H":     "help",
		"Ctrl+Q":     "quit",
	}
}

func (keys Keybindings) ToCmd(ev *tcell.EventKey) string {
	name := enix_tcell.EventKeyName(ev)

	if cmd, ok := keys[name]; ok {
		return cmd
	}

	return ""
}
