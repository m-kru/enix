package cfg

import (
	"strings"

	"github.com/gdamore/tcell/v2"
	enix_tcell "github.com/m-kru/enix/internal/tcell"
)

// Keybindings struct represents configured keybindings.
type Keybindings map[string]string

func (keys Keybindings) ToCmd(ev *tcell.EventKey) (string, string) {
	name := enix_tcell.EventKeyName(ev)

	str, ok := keys[name]
	if !ok {
		return "", ""
	}

	cmd, args, _ := strings.Cut(strings.TrimSpace(str), " ")

	return cmd, args
}

// KeybindingsDefault returns default keybindings.
func KeybindingsDefault() Keybindings {
	return map[string]string{
		// Alternation
		"Alt+Up":   "line-up",
		"Alt+Down": "line-down",
		"Rune[<]":  "deindent",
		"Rune[>]":  "indent",
		"Rune[r]":  "replace",
		"Rune[u]":  "undo",
		"Alt+J":    "join-below",
		"Alt+K":    "join-above",
		// Cmd
		"Rune[:]": "cmd",
		// Cursor
		"Rune[b]": "prev-word-start",
		"Rune[e]": "word-end",
		"Rune[w]": "word-start",
		"Rune[j]": "down",
		"Rune[k]": "up",
		"Rune[h]": "left",
		"Rune[l]": "right",
		"Ctrl+J":  "spawn-down",
		"Ctrl+K":  "spawn-up",
		"Rune[a]": "line-start",
		"Rune[$]": "line-end",
		// Deletion
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Alt+W":      "del-word",
		"Alt+{":      "del-within-brace",
		"Alt+[":      "del-within-bracket",
		"Alt+(":      "del-within-paren",
		// File
		"Rune[o]": "open",
		"Rune[s]": "save",
		// Tab
		"Ctrl+T": "tab-open",
		// Miscellaneous
		"Enter":   "newline",
		"Esc":     "esc",
		"Rune[c]": "copy",
		"Rune[f]": "find",
		"Rune[i]": "insert",
		//"Rune[h]": "help",
		"Rune[q]": "quit",
		// "Ctrl+U":     "undo", Alt+U
		"Rune[v]": "paste",
		"Rune[W]": "sel-word",
		"Rune[ ]": "space",
		"Tab":     "tab",
		// Selection
		"Ctrl+L": "sel-line",
		// View
		"Down": "view-down",
		"Up":   "view-up",
	}
}

// PromptKeybindingsDefault returns default keybindings for command prompt.
func PromptKeybindingsDefault() Keybindings {
	return map[string]string{
		"Left":       "left",
		"Right":      "right",
		"Ctrl+Left":  "prev-word-start",
		"Ctrl+Right": "word-end",
		"Ctrl+A":     "line-start",
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Enter":      "enter",
		"Esc":        "esc",
	}
}

// InsertKeybindingsDefault returns default keybindings for tab insert mode.
func InsertKeybindingsDefault() Keybindings {
	return map[string]string{
		"Ctrl+Left":  "prev-word-start",
		"Ctrl+Right": "word-end",
		"Ctrl+A":     "line-start",
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Enter":      "enter",
		"Esc":        "esc",
	}
}
