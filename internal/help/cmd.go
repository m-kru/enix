package help

import (
	"slices"
	"strings"
)

var Commands = map[string]string{

	// Command

	"cmd": `cmd # Starts command prompt.
The command does not work in the prompt, as the prompt is already active.`,

	"cmd-error": `cmd-error # Displays error message in the command prompt.
The command is useful for enix debugging and development.`,

	"cmd-info": `cmd-info # Displays message in the command prompt.
The command is useful for enix debugging and development.`,

	"cmd-prev": `cmd-prev # Executes previous command.`,

	// Cursor

	"add-cursor": `add-cursor L [C=1] # Creates a new cursor at line L and column C.
If L exceeds than the number of lines in the tab, the cursor is placed in the last line.
If C esceeds than the number of columns in a given line, the cursor is placed in the last column.`,

	"cursor-count": `cursor-count # Prints the number of cursors in the current tab.`,

	"down": `down # Moves cursor down.
If cursor is in the last line, nothing happens.`,

	"dump-cursor": `dump-cursor N # Dumps Nth cursor data.
The command is useful for enix debugging and development.`,

	"g": `g # An alias for the 'go' command.`,

	"go": `go position|mark-name # Goes to the position or restores a mark.
Valid syntaxes for the position are:
  - go 10    # Goes to line 10 column 1,
  - go 10 5  # Goes to line 10 column 5,
  - go 10:5  # Goes to line 10 column 5,
  - go -1    # Goes to the last line column 1,
  - go -1 -1 # Goes to the last line last column,
  - go -1:-1 # Goes to the last line last column,
  - go tmp   # Goes to the mark named "tmp".`,

	"left": `left # Moves cursor left.
If the cursor is in the first column of a line, it is moved to the last column of the previous line.
Unless this is the first line.
In such a case, nothing happens.`,

	"line-down": `line-down # Moves line down by one line.`,

	"line-end": `line-end # Moves cursor to the line end.`,

	"line-start": `line-start # Moves cursor to the line start.
The first non whitespace rune is considered the start of the line.
However, if a cursor already is at the first non whitespace rune in the line,
then line-start moves cursor to the first line rune, even if it is a whitespace.`,

	"line-up": `line-up # Moves line up by one line.`,

	"prev-word-start": `word-start # Moves cursor to the previous word start.`,

	"right": `right # Moves cursor right.
If cursor is in the last column of a line, it is moved into the first column of the next line.
Unless this is the lastt line.
In such a case, nothing happens.`,

	"spawn-down": `spawn-down # Spawns a new cursor in the below line.`,

	"spawn-up": `spawn-up # Spawns a new cursor in the above line.`,

	"up": `up # Moves cursor up.
If cursor is in the fisrt line, nothing happens.`,

	"word-end": `word-end # Moves cursor to the word end.`,

	"word-start": `word-start # Moves cursor to the next word start.`,

	// Deletion

	"backspace": `del # Deletes text before cursor or selected text.`,

	"del": `del # Deletes text under cursor or selection.`,

	"del-line": `del-line # Deletes lines with cursor or selection.`,

	"del-word": `del-word # Deletes words under cursor or selection.`,

	"trim": `trim # Trims trailing whitespaces from all lines.`,

	// File

	"e": `e # An alias for the 'edit' command.`,

	"edit": `edit path/to/file ... # Opens a file in a new tab.
If path to the file is not provided, it opens a new empty tab.

The command is named "edit" instead of "open" for the following reason.
If you have an alias for the enix in a shell, then the alias is probably 'e'.
The edit command has an alias 'e' inside the enix.
This means that opening a file in a shell and the enix has the same syntax when using aliases.`,

	"s": `s # An alias for the 'save' command.`,

	"save": `save [path/to/file] # Saves tab to the file using the provided path.
If the path to a file is not provided, then it uses the current file path.`,

	"type": `type type # Enforces file type.`,

	// Find

	"find-next": "find-next # Selects the next find.",

	"find-sel-next": "find-sel-next # Preserves selections and selects the next find.",

	"find-prev": "find-prev # Selects the previous find.",

	"find-sel-prev": "find-sel-prev # Preserves selections and selects the previous find.",

	// Indent

	"deindent": `deindent # Decreases the indent of lines with cursor.
In the case of selections, it deindents all selected lines.
Even if the selection starts or ends in the middle of a line.`,

	"indent": `indent # Increases indent of lines with cursor.
In the case of selections, it increases all selected lines.
Even if the selection starts or ends in the middle of a line.`,

	// Insert

	"insert": `insert # Enters tab insert mode.`,

	"insert-line-above": `insert-line-above # Adds an empty line above and enters the insert mode.`,

	"insert-line-below": `insert-line-below # Adds an empty line below and enters the insert mode.`,

	"insert-tab": `tab # Inserts tab rune.`,

	"newline": `newline # Inserts a newline.`,

	"rune": `rune r # Inserts rune r under the cursor or selection position.`,

	"space": `space # Inserts space rune.`,

	// Match

	"mb": `mb # An alias for the 'match-bracket' command.`,

	"match-bracket": `match-bracket # Moves cursor to the matching square bracket.
If cursor is placed on ']', the command looks for matching '['.
Otherwise, the command looks for ']'.`,

	"mc": `mc # An alias for the 'match-curly' command.`,

	"match-curly": `match-curly # Moves cursor to the matching curly brace.
If cursor is placed on '}', the command looks for matching '{'.
Otherwise, the command looks for '}'.`,

	"mp": `mp # An alias for the 'match-paren' command.`,

	"match-paren": `match-paren # Moves cursor to the matching parenthesis.
If cursor is placed on ')', the command looks for matching '('.
Otherwise, the command looks for ')'.`,

	// Miscellaneous

	"a": `a # An alias for the 'align' command.`,

	"align": `align # Aligns columns of cursors.`,

	"autosave": `autosave n # Changes value for autosave period to n seconds.
n must be natural.

The value applies only to the current session.
The AutoSave value in the config file is not changed.

The tab is autosaved every n seconds, not n seconds after the last rune insert.
Setting low n value on a constrained system may lead to performance drop.

The tab is not autosaved if the corresponding file in the file system doesn't exist.`,

	"cut": `cut # Cuts selected text.
For cursors, the cut command cuts all lines containing cursors.
This is because the user rarely wants to cut a single rune.
On the other hand, cutting a single line is a very common action.
Cutting a single rune is still possible by creating a selection of width 1.
For example, by executing commands 'sel-right' and 'sel-left.`,

	"esc": `esc # Escapes the current context.
The actual action depends on the context.
For example, if the focus is on the command prompt, then the command escapes the command prompt and allows user to continue tab editing.
If the focus in on a tab editing and there is any selection, then the command clears all selections.`,

	"ft": `ft ft # An alias for the 'filetype' command.`,

	"filetype": `filetype ft # Changees filetype to ft.`,

	"h": `h # An alias for the 'help' command.`,

	"help": `help topic|command-name # Displays help message for a given topic or command in a new tab.
If neither topic nor command name is provided, the command displays help message for the 'help' command.
The same as 'help help' command would do.
Valid topics are:
  - commands - list of available commands with their synopsys,
  - config - general description of enix configuration,
  - cursors - what they are and how they work,
  - highlighting - synax highlighting,
  - enix - general overview of the enix editor,
  - keybindings - how to set, and what to watch out for,
  - selections - what they are and how they work.`,

	"join": `join # Joins line with cursor with the below line.`,

	"key-name": `key-name # Returns name of the pressed key combo.
The command opens a new tab named "key-name".
On each keystroke a new line is added to the tab with the key combo name.
Once user leaves the tab key-name state by hitting the escape key, it can only be reentered by executing the command again.`,

	"m": `m # An alias for the 'mark' command.`,

	"mark": `mark name # Creates a new named mark.
Marks allow to record current cursor or selection position.
Mark name must not start with a digit or '-' rune.
To restore marks one has to use the 'go' command providing as an argument the name of a mark.`,

	"line-count": `line-count # Prints the number of lines in the current tab.`,

	"path": `path file/system/path # Inserts file system path.
The command is useful because user can automatically expand the path using the tab key.
Enix doesn't verify if the provided path exists.`,

	"pwd": `pwd # Prints working directory.`,

	"q": `q # An alias for the 'quit' command.`,

	"quit": `quit # Quits the current tab.
If the tab is the last tab, then the command also quits the enix editor.
The command returns an error if current tab has unsaved changes.`,

	"q!": `q! # An alias for the 'quit!' command.`,

	"quit!": `quit! # Forces quit.,
Forced version of the 'quit' command.
The command quits the tab even if there are unsaved changes.`,

	"replace": `replace # Replaces rune under cursor or runes under selection.
The command first deletes runes under the cursor or selection, and then enters the insert moode for a single insertion.`,

	"search": `search regex # Searches for patterns matching given regex.
The regex can include space ' ' runes.
However, the first space after the 'search' command name is not included in the regex.

The command does not automatically jump to the first find.
Such a behavior helps to avoid unexpected tab view changes.
The user has to explicitly execute 'find-' commands to navigate between finds.`,

	"sh": `sh [-i] cmd [arg] ... # Executes command cmd in the shell.
The command obtains a shell via the $SHELL environment variable.
If this variable is not set, the command returns an error.

The text directed by the command to the stdout gets pasted into the tab.
If the text ends with the newline character, then the line-based paste is used.
Otherwise, the regular paste is used.

In the case of cursors, nothing is fed into the stdin.
The command is executed only once and the text from stdout is pasted for every cursor.
However, the paste behaves slightly different than the paste command.
The paste starts at the cursors position, instead of one position to the right.

In the case of selections, the command is executed once per each selection.
For each selection, the selection text is used as the stdin for the command.

Stderr can be used to execute arbitrary enix commands after stdout paste.
In such a case, a line within the stderr stream must start with the 'enix:' label.
A care must be taken, as those commands are executed as in the script mode.
In the case of multiple selections, only commands for the last call are executed.

The -i flag controls whether current indent shall be added to the text generated by the command to the stdout.
For cursors, current indent is the indent of a line with the cursors.
For selections, current indent is the indent of the first line of the selection.

The command can be used for various purposes and is the only official plug-in system for enix.
For example, the command can be used with https://github.com/m-kru/tmpl for template inserting.

Before executing the command, enix sets the following environment variables.
  - ENIX_FILEPATH - absolute path to the current tab file,
  - ENIX_FILETYPE - file type of the current tab.
These variables can be used not only within the executed program, but also within the enix 'sh' command.
As these are environment variables, they will be expanded by the shell.`,

	"suspend": `suspend # Stops enix and gives control to the shell.
The command has been so far tested only on Linux.`,

	"trim-on-save": `trim-on-save [value] # Sets the TrimOnSave config.
Possible values are 0 for false, and ony other integer value for true.
If value is not present, the command returns current value.`,

	"undo": `undo # Undos last action modifying the tab content.`,

	"yank": `yank # Copies line with cursor or selection.
For cursors, the command copies all lines containing cursors.
This is because the user rarely wants to copy a single rune.
On the other hand, copying a single line is a very common action.
Copying a single rune is still possible by creating a selection of width 1.
For example, by executing commands 'sel-right' and 'sel-left.`,

	// Selection

	"sel": `sel regex # Selects all finds of regex pattern within selections.
If there are no selections within the tab, the whole tab is searched.`,

	"sel-all": `sel-all # Selets all tab content.`,

	"sb": `sel-bracket # An alias for the 'sel-bracket' command.`,

	"sel-bracket": `sel-bracket # Selects content within matching brackets.`,

	"sc": `sel-curly # An alias for the 'sel-curly' command.`,

	"sel-curly": `sel-curly # Selects content within matching curly braces.`,

	"sel-count": `sel-count # Prints the number of selections in the current tab.`,

	"sel-left": `sel-left # Extends or shrinks selection one rune left.`,

	"sel-line": `sel-line # Selects whole line.
If whole current line is already selected, it extends the selection with the next line.`,

	"sel-right": `sel-right # Extends or shrinks selection one rune right.`,

	"sel-white": `sel-white # Selects sequence of whitespaces.
If cursors is not placed on a whitespace character, nothing is selected.`,

	"sp": `sel-paren # An alias for the 'sel-paren' command.`,

	"sel-paren": `sel-paren # Selects content within matching parentheses.`,

	"sel-prev-word-start": `sel-prev-word-start # Selects, extends or shrinks selection to the start of the previous word.`,

	"sel-to-tab": `sel-to-tab [tab-path] # Opens a new tab with the content of current selections.`,

	"sel-word": `sel-word # Selects word under cursor.`,

	"sel-word-end": `sel-word-end # Selects, extends or shrinks selection to the end of next word.`,

	"sel-switch-cursor": `sel-switch-cursor # Changes position of the cursor within selection.`,

	"sel-tab-end": `sel-tab-end # Selects or extends selection to the tab end.`,

	// Tab

	"t": `t  # An alias for the 'tab' command.`,

	"tab": `tab regex # Switch to the tab which path matches the regex.
If more than one tab is found, the error is reported.`,

	"tab-count": `tab-count # Prints the number of tabs.`,

	"tn": `tn # An alias for the 'tab-next' command.`,

	"tab-next": `tab-next # Cycles to the next tab.
If the current tab is the last tab, then it wraps to the first tab.`,

	"tab-move": `tab-move N [pattern] # Moves tab which name matches pattern to position N.
If pattern is not provided, then current tab is moved.
The pattern must be unambiguous.
If N is not provided, it is assumed to equal 1.`,

	"tab-open": `tab-open # Opens a new empty tab.`,

	"tp": `tp # An alias for the 'tab-prev' command.`,

	"tab-prev": `tab-prev # Cycles to the previous tab.
If the current tab is the first tab, it wraps to the last tab.`,

	"tab-switch": `tab-switch pattern1 [pattern2] # Switch tabs which names match patterns.
The patterns must be unambiguous.
If the second is absent, it is assumed to be the current tab.`,

	// View

	"view": "view line-number column # Sets the view start at provided line number and column.",

	"vc": "vc # An alias for the 'view-center' command.",

	"view-center": "view-center # Centers the view.",

	"view-down": "view-down # Scrolls the view down by one line.",

	"view-down-half": "view-down-half # Scrolls the view down by half of the screen height.",

	"ve": "ve # An alias for the 'view-end' command.",

	"view-end": "view-end # Scrolls the view to the tab end.",

	"view-left": "view-left # Scrolls the view left by one column.",

	"view-right": "view-right # Scrolls the view right by one column.",

	"vs": "vs # An alias for the 'view-start' command.",

	"view-start": "view-start # Scrolls the view to the tab start.",

	"view-up": "view-up # Scrolls the view up by one line.",

	"view-up-half": "view-up-half # Scrolls the view up by half of the screen height.",
}

func GetCommandNames(prefix string) []string {
	cmds := make([]string, 0, len(Commands))

	for name := range Commands {
		if prefix != "" && !strings.HasPrefix(name, prefix) {
			continue
		}
		cmds = append(cmds, name)
	}

	slices.Sort(cmds)

	return cmds
}

// IsPathCmd returns true if provided command accepts a file system path as the last argument.
func IsPathCmd(cmd string) bool {
	pathCmds := map[string]bool{
		"e": true, "edit": true, "path": true, "s": true, "save": true,
	}
	if _, ok := pathCmds[cmd]; ok {
		return true
	}
	return false
}
