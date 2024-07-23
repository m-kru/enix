package cfg

import "github.com/gdamore/tcell/v2"

// Keybindings struct represents configured keybindings.
type Keybindings map[string]string

// KeybindingsDefault returns default keybindings.
func KeybindingsDefault() Keybindings {
	return map[string]string{
		// Cmd
		"Ctrl+E": "cmd",
		// Cursor
		"Down":  "cursor-down",
		"Left":  "cursor-left",
		"Right": "cursor-right",
		"Up":    "cursor-up",
		// File
		"Ctrl+O": "file-open",
		"Ctrl+S": "file-save",
		// Tab
		"Ctrl+T": "tab-open",
		// Miscellaneous
		"Esc":    "escape",
		"Ctrl+Q": "quit",
	}
}
func (keys Keybindings) ToCmd(ev *tcell.EventKey) string {
	name := ev.Name()

	if cmd, ok := keys[name]; ok {
		return cmd
	}

	return ""
}
