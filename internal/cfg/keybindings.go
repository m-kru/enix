package cfg

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	enix_tcell "github.com/m-kru/enix/internal/tcell"
)

// Keybindings struct represents configured keybindings.
type Keybindings map[string]string

// KeybindingsDefault returns default keybindings.
func KeybindingsDefault() Keybindings {
	return map[string]string{
		// Alternation
		"Alt+Up":    "line-up",
		"Alt+Down":  "line-down",
		"Alt+Left":  "deindent",
		"Alt+Right": "indent",
		"Alt+R":     "replace",
		"Alt+U":     "undo",
		"Alt+J":     "join-below",
		"Alt+K":     "join-above",
		// Cmd
		"Ctrl+E": "cmd",
		// Cursor
		"Down":       "down",
		"Left":       "left",
		"Right":      "right",
		"Up":         "up",
		"Ctrl+Down":  "spawn-down",
		"Ctrl+Left":  "word-start",
		"Ctrl+Right": "word-end",
		"Ctrl+A":     "line-start",
		// Deletion
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Alt+W":      "del-word",
		"Alt+{":      "del-within-brace",
		"Alt+[":      "del-within-bracket",
		"Alt+(":      "del-within-paren",
		// File
		"Ctrl+O": "open",
		"Ctrl+S": "save",
		// Tab
		"Ctrl+T": "tab-open",
		// Miscellaneous
		"Enter":  "enter",
		"Esc":    "esc",
		"Ctrl+C": "copy",
		"Ctrl+F": "find",
		"Ctrl+H": "help",
		"Ctrl+Q": "quit",
		// "Ctrl+U":     "undo", Alt+U
		"Ctrl+V": "paste",
		"Ctrl+W": "sel-word",
		// Selection
		"Ctrl+K": "sel-line-end",
		"Ctrl+L": "sel-line",
	}
}

func (keys Keybindings) ToCmd(ev *tcell.EventKey) (string, string) {
	name := enix_tcell.EventKeyName(ev)

	str, ok := keys[name]
	if !ok {
		return "", ""
	}

	cmd, args, _ := strings.Cut(strings.TrimSpace(str), " ")

	return cmd, args
}
