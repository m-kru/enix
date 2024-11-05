package cfg

import (
	"github.com/m-kru/enix/internal/cmd"
	enixTcell "github.com/m-kru/enix/internal/tcell"

	"github.com/gdamore/tcell/v2"
)

// Keybindings struct represents configured keybindings.
type Keybindings map[string]string

func (keys Keybindings) ToCmd(ev *tcell.EventKey) (cmd.Command, error) {
	name := enixTcell.EventKeyName(ev)

	str, ok := keys[name]
	if !ok {
		return cmd.Command{}, nil
	}

	return cmd.Parse(str)
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
		"Rune[;]": "line-end",
		// Deletion
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Rune[d]":    "del",
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
		"Enter":       "newline",
		"Esc":         "esc",
		"Rune[c]":     "copy",
		"Rune[f]":     "find",
		"Rune[i]":     "insert",
		"Rune[m]":     "mark tmp",
		"Rune[M]":     "go tmp",
		"Alt+Rune[j]": "join",
		"Ctrl+Z":      "suspend",
		//"Rune[h]": "help",
		"Rune[q]": "quit",
		// "Ctrl+U":     "undo", Alt+U
		"Rune[v]": "paste",
		"Rune[W]": "sel-word",
		"Rune[ ]": "space",
		"Tab":     "tab",
		// Selection
		"Rune[J]": "sel-down",
		"Rune[H]": "sel-left",
		"Rune[L]": "sel-right",
		// View
		"Down":       "view-down",
		"Ctrl+Down":  "5 view-down",
		"Right":      "view-right",
		"Ctrl+Right": "5 view-right",
		"Up":         "view-up",
		"Ctrl+Up":    "5 view-up",
		"Left":       "view-left",
		"Ctrl+Left":  "5 view-left",
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
