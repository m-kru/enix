package cfg

import (
	"github.com/m-kru/enix/internal/cmd"
	enixTcell "github.com/m-kru/enix/internal/tcell"

	"github.com/gdamore/tcell/v2"
)

var Keys Keybindings
var KeysInsert Keybindings
var KeysPrompt Keybindings

// Keybindings struct represents configured keybindings.
type Keybindings map[string]string

func (keys Keybindings) ToCmd(ev *tcell.EventKey) (cmd.Command, error) {
	name := enixTcell.EventKeyName(ev)

	str, ok := keys[name]
	if !ok {
		return cmd.Command{RepCount: 0, Name: "", Args: nil}, nil
	}

	return cmd.Parse(str)
}

// DefaultKeybindings returns default keybindings.
func DefaultKeybindings() Keybindings {
	return map[string]string{
		// Alternation
		"Alt+Rune[j]": "join",
		"Alt+Rune[d]": "line-down",
		"Alt+Rune[u]": "line-up",
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
		"Rune[f]": "line-end",
		"Rune[t]": "sh -i tmpl -enix $ENIX_FILETYPE",
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
		"Rune[s]": "save",
		// Find
		"Rune[n]": "find-next",
		"Rune[N]": "find-sel-next",
		"Rune[g]": "find-prev",
		"Rune[G]": "find-sel-prev",
		"Ctrl+N":  "find-desel-next",
		"Ctrl+G":  "find-desel-prev",
		// Tab
		"Ctrl+T": "tab-open",
		// Miscellaneous
		"Enter":   "newline",
		"Esc":     "esc",
		"Rune[,]": "align",
		"Rune[-]": "go -1",
		"Home":    "go 1",
		"End":     "go -1",
		"Rune[c]": "change",
		"Rune[/]": "search",
		"Rune[i]": "insert",
		"Rune[o]": "insert-line-below",
		"Rune[O]": "insert-line-above",
		"Rune[m]": "mark tmp",
		"Rune[M]": "go tmp",
		"Ctrl+Z":  "suspend",
		"Rune[q]": "quit",
		"Rune[p]": "paste",
		"Rune[P]": "paste-before",
		"Rune[r]": "replace",
		"Rune[u]": "undo",
		"Rune[U]": "redo",
		"Rune[x]": "cut",
		"Rune[y]": "yank",
		"Rune[ ]": "space",
		"Rune[>]": "indent",
		"Tab":     "insert-tab",
		"Ctrl+P":  "path",
		"Rune[]]": "match-bracket",
		"Rune[}]": "match-curly",
		"Rune[)]": "match-paren",
		// Selection
		"Rune[%]": "sel-all",
		"Rune[v]": "sel-line",
		"Rune[A]": "sel-line-start",
		"Rune[F]": "sel-line-end",
		"Rune[V]": "sel-prev-line",
		"Rune[B]": "sel-prev-word-start",
		"Rune[J]": "sel-down",
		"Rune[H]": "sel-left",
		"Rune[L]": "sel-right",
		"Rune[K]": "sel-up",
		"Rune[W]": "sel-word-start",
		"Rune[E]": "sel-word-end",
		"Rune[z]": "sel-word",
		"Rune[;]": "sel-switch-cursor",
		"Rune[C]": "sel-tab-end",
		"Rune[[]": "sel-bracket",
		"Rune[{]": "sel-curly",
		"Rune[(]": "sel-paren",
		// View
		"Ctrl+W":     "view-center",
		"Down":       "view-down",
		"Ctrl+D":     "view-down-half",
		"Ctrl+Down":  "5 view-down",
		"Right":      "view-right",
		"Ctrl+Right": "5 view-right",
		"Up":         "view-up",
		"Ctrl+U":     "view-up-half",
		"Ctrl+Up":    "5 view-up",
		"Left":       "view-left",
		"Ctrl+Left":  "5 view-left",
	}
}

// DefaultPromptKeybindings returns default keybindings for command prompt.
func DefaultPromptKeybindings() Keybindings {
	return map[string]string{
		"Down":       "down",
		"Up":         "up",
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

// DefaultInsertKeybindings returns default keybindings for tab insert mode.
func DefaultInsertKeybindings() Keybindings {
	return map[string]string{
		"Ctrl+A":     "line-start",
		"Ctrl+T":     "sh -i tmpl -enix $ENIX_FILETYPE",
		"Ctrl+P":     "path",
		"Backspace":  "backspace",
		"Backspace2": "backspace",
		"Delete":     "del",
		"Enter":      "enter",
		"Esc":        "esc",
		// View
		"Ctrl+W":     "view-center",
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
