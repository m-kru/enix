package help

var Commands = map[string]string{

	// Config

	"config-dir": `config-dir  # Shows enix config home directory path.`,

	// Command

	"cmd": `cmd # Starts command prompt.`,

	"cmd-error": `cmd-error # Displays error message in the command prompt.
This command is useful for enix debugging and development.`,

	"cmd-info": `cmd-info # Displays message in the command prompt.
This command is useful for enix debugging and development.`,

	"cmd-list": `cmd-list # Lists all available commands in alphabetical order with one sentence summary.`,

	"cmd-prev": `cmd-prev # Executes previous command.`,

	// Cursor

	"add-cursor": `add-cursor L [C=1] # Creates new cursor at line L and column C.
If L is greater than the number of lines in the file, then the cursor is placed
in the last line. If C is greater than the number of columns in a given line,
then the cursor is placed in the last column.`,

	"cursor-count": `cursor-count # Prints the number of cursors in the current tab.`,

	"down": `down # Moves cursor down.
If cursor is in the last line, nothing happens.`,

	"dump-cursor": `dump-cursor N # Dumps Nth cursor data.
This command is useful for enix debugging and development.`,

	"g": `g # An alias to the go command.`,

	"go": `go position|mark-name # Goes to the position or restores a mark.
Valid syntaxes for the position are:
  - go 10    # Goes to line 10 column 1,
  - go 10 5  # Goes to line 10 column 5,
  - go 10:5  # Goes to line 10 column 5,
  - go -1    # Goes to the last line column 1,
  - go -1 -1 # Goes to the last line last column,
  - go -1:-1 # Goes to the last line last column,
  - go tmp   # Goes to mark named tmp.`,

	"left": `left # Moves cursor left.
If cursor is in the first column of a line, then it is moved into the last
column of the previous line. Unless this is the first line. In such a case,
nothing happens.`,

	"line-down": `line-down # Moves line down by one line.`,

	"line-end": `line-end # Moves cursor to the line end.`,

	"line-start": `line-start # Moves cursor to the line start.
The first non whitespace rune is considered the start of the line.
However, if a cursor already is at the first non whitespace rune in the line,
then line-start moves cursor to the first line rune, even if it is a whitespace.`,

	"line-up": `line-up # Moves line up by one line.`,

	"prev-word-start": `word-start # Moves cursor to the previous word start.`,

	"right": `right # Moves cursor right.
If cursor is in the last column of a line, then it is moved into the first
column of the next line. Unless this is the lastt line. In such a case,
nothing happens.`,

	"spawn-down": `spawn-down # Spawns a new cursor in the below line.`,

	"spawn-up": `spawn-up # Spawns a new cursor in the above line.`,

	"up": `up # Moves cursor up.
If cursor is in the fisrt line, nothing happens.`,

	"word-end": `word-end # Moves cursor to the word end.`,

	"word-start": `word-start # Moves cursor to the next word start.`,

	// Deletion

	"backspace": `del # Deletes text before cursor or selected text.`,

	"del": `del # Deletes text under cursor/selection.`,

	"del-line": `del-line # Deletes lines with cursor/selection.`,

	"del-word": `del-word # Deletes words under cursor/selection.`,

	"trim": `trim # Trims trailing whitespaces from all lines.`,

	// File

	"open": `open path/to/file ... # Opens file in a new tab.
If path to file is not provided, then it opens a new empty tab.`,

	"o": `o # An alias to the open command.`,

	"save": `save [path/to/file] # Saves file using the provided path to file.
If path to file is not provided, then it uses the current file path.`,

	"type": `type type # Enforces file type.`,

	// Find

	"find-next": "find-next # Selects next find.",

	"find-sel-next": "find-sel-next # Preserves selections and selects next find.",

	"find-prev": "find-prev # Selects previous find.",

	"find-sel-prev": "find-sel-prev # Preserves selections and selects previous find.",

	// Indent

	"deindent": `deindent # Decreases indent of lines with cursor.
In the case of selections, it deindents all selected lines. Even if the
selection starts/ends in the middle of a line.`,

	"indent": `indent # Increases indent of lines with cursor.
In the case of selections, it increases all selected lines. Even if the
selection starts/ends in the middle of a line.`,

	// Insert

	"insert": `insert # Enters tab insert mode.`,

	"insert-line-above": `insert-line-above # Adds an empty line above and enters the insert mode.`,

	"insert-line-below": `insert-line-below # Adds an empty line below and enters the insert mode.`,

	"newline": `newline # Inserts a newline.`,

	"rune": `rune r # Inserts rune r under the cursor or selection position.`,

	"space": `space # Inserts space rune.`,

	"tab": `tab # Inserts tab rune.`,

	// Match

	/*
	   "match-brace":   struct{}{},
	   "match-bracket": struct{}{},
	   "match-paren":   struct{}{},
	*/

	// Miscellaneous

	"align": `align # Aligns columns of cursors.`,

	"cut": `cut # Cuts selected text.
For cursors, the cut command cuts all lines containing cursors.
This is because the user rarely wants to cut a single rune.
On the other hand, cutting a single line is a very common action.
Cutting a single rune is still possible by creating a selection of width 1.
For example, by executing commands 'sel-right' and 'sel-left.`,

	"dump-colors": `dump-colors # Dumps colorscheme configuration to JSON format.`,

	"esc": `esc # Escapes the current context.
The actual action depends on the context. For example, if the focues is on the
command prompt, then escape command escapes the command prompt and allows user
to continue file editing. If the focus in on a file editing and there is any
selection, then escape command clears all selections.`,

	"h": `h # An alias to the help command.`,

	"help": `help topic|command-name # Displays help message for a given topic or command in a new tab.
If neither topic nor command name is provided displays help message
for the help command. The same as 'help help'. Valid topics are:
  - commands - list of available commands with their synopsys,
  - config - general description of enix configuration,
  - cursors - what they are and how they work,
  - enix - general overview of the enix editor,
  - keybindings - how to set, and what to watch out for,
  - selections - what they are and how they work.`,

	"join": `join # Joins line with cursor with below line.`,

	"key-name": `key-name # Returns name of the pressed key combo.
The command opens a new tab named "key-name". On each keystroke a new line
is added to the tab with the key combo name. Once user leaves the tab key-name
state by hitting the escape key, it can only be reentered by executing the
key-name command again.`,

	"mark": `mark name # Creates new named mark.
Marks allow to record current cursors or selections positions.
Mark name must not start with a digit or '-' rune. To restore marks one
has to use the go command providing as an argument the name of a mark.`,
	"m": `m # An alias to the mark command.`,

	"pwd": `pwd # Prints working directory.`,

	"q": `q # An alias to the quit command.`,

	"quit": `quit # Quits tab.
If the tab is the last tab opened tab, then quit also quits the enix editor.
Quit returns an error if current tab has unsaved changes.`,

	"q!": `q! # An alias to the quit! command.`,

	"quit!": `quit! # Force quit.,
Forced version of the quit command. It quits the tab even if there are unsaved changes.`,

	"replace": `replace # Replaces rune under cursor or runes under selection.
The command first deletes rune under the cursor, or runes under selection, and then
enters the insert moode for a single insertion.`,

	"undo": `undo # Undos last action modifying the tab content.`,

	"yank": `yank # Copies selected text.
For cursors, the yank command copies all lines containing cursors.
This is because the user rarely wants to copy a single rune.
On the other hand, copying a single line is a very common action.
Copying a single rune is still possible by creating a selection of width 1.
For example, by executing commands 'sel-right' and 'sel-left.`,

	// Selection

	"sel-count": `sel-count # Prints the number of selections in the current tab.`,

	"sel-left": `sel-left # Extends/shrinks selection one rune left.`,

	"sel-line": `sel-line # Selects whole line.
If whole current line is already selected, then it extends the selection with the next line.`,

	"sel-right": `sel-right # Extends/shrinks selection one rune right.`,

	"sel-white": `sel-white # Selects sequence of whitespaces.
If cursors is not placed on a whitespace character, nothing is selected.`,

	"sel-prev-word-start": `sel-prev-word-start # Selects or extends selection to the start of previous  word.`,

	"sel-to-tab": `sel-to-tab [tab-path] # Opens a new tab with the content of current selections.`,

	"sel-word": `sel-word # Selects word under cursor.`,

	"sel-word-end": `sel-word-end # Selects or extends selection to the end of next word.`,

	"suspend": `suspend # Stops process and gives control to shell.
This command has been so far tested only on Linux.`,

	// Tab

	"tab-count": `tab-count # Prints the number of opened tabs.`,

	"tab-next": `tab-next # Cycles to the next tab.
If the current tab is the last tab, then it wraps to the first tab.`,

	"tn": `tn # An alias to the tab-next command.`,

	"tab-move": `tab-move N [pattern] # Moves tab which name matches pattern to position N.
If pattern is not provided, then current tab is moved. The pattern must be
unambiguous. If N is not provided, then it is assumed to be equal 1.`,

	"tab-open": `tab-open # Opens a new empty tab.`,

	"tab-prev": `tab-prev # Cycles to the previous tab.
If the current tab is the first tab, then it wraps to the last tab.`,

	"tp": `tp # An alias to the tab-prev command.`,

	"tab-switch": `tab-switch pattern1 [pattern2] # Switch tabs which names match patterns.
The patterns must be unambiguous. If the second is absent, then it is assumed
to be the current tab.`,

	// View

	"view": "view line-number column # Sets view start at provided line number and column.",

	"view-center": "view-center # Centers view.",

	"view-down": "view-down # Scrolls view down by one line.",

	"view-left": "view-left # Scrolls view left by one column.",

	"view-right": "view-right # Scrolls view right by one column.",

	"view-up": "view-up # Scrolls view up by one line.",
}
