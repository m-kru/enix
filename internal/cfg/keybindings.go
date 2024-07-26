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
		// Alternation
		"Alt+Up":    "move-up",
		"Alt+Down":  "move-down",
		"Alt+Left":  "indent-decrease",
		"Alt+Right": "indent-increase",
		"Alt+R":     "replace",
		"Alt+U":     "undo",
		"Alt+J":     "line-join-below",
		"Alt+K":     "line-join-above",
		// Cmd
		"Ctrl+E": "cmd",
		// Cursor
		"Down":       "cursor-down",
		"Left":       "cursor-left",
		"Right":      "cursor-right",
		"Up":         "cursor-up",
		"Ctrl+Left":  "cursor-word-start",
		"Ctrl+Right": "cursor-word-end",
		"Ctrl+A":     "cursor-line-start",
		// Deletion
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Alt+W":      "del-word",
		"Alt+{":      "del-within-brace",
		"Alt+[":      "del-within-bracket",
		"Alt+(":      "del-within-paren",
		// File
		"Ctrl+O": "file-open",
		"Ctrl+S": "file-save",
		// Tab
		"Ctrl+T": "tab-open",
		// Miscellaneous
		"Enter":      "enter",
		"Esc":        "escape",
		"Ctrl+C":     "copy",
		"Ctrl+F":     "find",
		"Ctrl+H":     "help",
		"Ctrl+Q":     "quit",
		// "Ctrl+U":     "undo", Alt+U
		"Ctrl+V": "paste",
		"Ctrl+W": "sel-word",
		// Selection
		"Ctrl+K": "sel-line-end",
		"Ctrl+L": "sel-line",
	}
}

func (keys Keybindings) ToCmd(ev *tcell.EventKey) string {
	name := enix_tcell.EventKeyName(ev)

	if cmd, ok := keys[name]; ok {
		return cmd
	}

	return ""
}
