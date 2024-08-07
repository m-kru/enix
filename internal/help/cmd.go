package help

var cmds = map[string]string{
	// Cursor

	"cursor-count": `cursor-count # Prints the number of cursors in current tab.`,

	"down": `down [N=1] # Moves cursor down.
The N parameter specifies how many lines cursor should be moved.
If cursor is in the last line, nothing happens.`,

	"left": `left # Moves cursor left.
If cursor is in the first column of a line, then it is moved into the last
column of the previous line. Unless this is the first line. In such a case,
nothing happens.`,

	"right": `right # Moves cursor right.
If cursor is in the last column of a line, then it is moved into the first
column of the next line. Unless this is the lastt line. In such a case,
nothing happens.`,

	"up": `up [N=1] # Moves cursor up.
The N parameter specifies how many lines cursor should be moved.
If cursor is in the fisrt line, nothing happens.`,

	"end": `end # Moves cursor to the last line.
If there are multiple cursors, they are first reduced to a single cursor.`,

	"word-start": `word-start # Moves cursor to the word start.`,
	"word-end":   `word-end # Moves cursor to the word end.`,
	"line-start": `line-start # Moves cursor to the line start.`,
	"line-end":   `line-end # Moves cursor to the line end.`,

	"spawn-down": `spawn-down # Spawns a new cursor in the below line.`,

	// Command

	"cmd": `cmd # Starts command prompt.`,
	"cmd-error": `cmd-error # Displays error message in the command prompt.
This command is useful for enix debugging and development.`,
	"cmd-info": `cmd-info # Displays message in the command prompt.
This command is useful for enix debugging and development.`,
	"cmd-list": `cmd-list # Lists all available commands in alphabetical order with one sentence summary.`,
	"cmd-prev": `cmd-prev # Executes previous command.`,

	// Config

	"tab-width": `tab-width N # Sets tab width to N.`,

	// Deletion

	"del":      `del # Deletes text under cursor/selection.`,
	"del-word": `del-word # Deletes words under cursor/selection.`,
	"del-line": `del-line # Deletes lines with cursor/selection.`,

	// File

	"open": `open [path/to/file] # Opens file in a new tab.
If path to file is not provided, then it opens a new empty tab.`,

	"save": `save [path/to/file] # Saves file using the provided path to file.
If path to file is not provided, then it uses the current file path.`,

	"type": `type type # Enforces file type.`,

	// Indent

	"indent": `indent # Increases indent of lines with cursor.
In the case of selections, it increases all selected lines. Even if the
selection starts/ends in the middle of a line.`,

	"deindent": `deindent # Decreases indent of lines with cursor.
In the case of selections, it deindents all selected lines. Even if the
selection starts/ends in the middle of a line.`,

	// Selection

	"sel-white": `sel-white # Selects sequence of whitespaces.
If cursors is not placed on a whitespace character, nothing is selected.`,

	"sel-word": `sel-word # Selects word under cursor.`,

	// Tab

	"tab-count": `tab-count # Prints the number of opened tabs.`,

	"tab-open": `tab-open # Opens a new empty tab.`,

	"tab-next": `tab-next # Cycles to the next tab.
If the current tab is the last tab, then it wraps to the first tab.`,

	"tab-prev": `tab-prev # Cycles to the previous tab.
If the current tab is the first tab, then it wraps to the last tab.`,

	"tab-move": `tab-move [N] [pattern] # Moves tab which name matches pattern to position N.
If pattern is not provided, then current tab is moved. The pattern must be
unambiguous. If N is not provided, then it is assumed to be equal 1.`,

	"tab-switch": `tab-switch pattern1 [pattern2] # Switch tabs which names match patterns.
The patterns must be unambiguous. If the second is absent, then it is assumed
to be the current tab.`,

	// Miscellaneous

	"dump-colors": `dump-colors # Dumps colorscheme configuration to JSON format.`,

	"esc": `esc # Escapes the current context.
The actual action depends on the context. For example, if the focues is on the
command prompt, then escape command escapes the command prompt and allows user
to continue file editing. If the focus in on a file editing and there is any
selection, then escape command clears all selections.`,

	"help": `help [topic|command-name] # Displays help message for a given topic or command in a newly open tab.
If neither topic nor command name is provided displays help message
for the help command. The same as 'help help'. Valid topics are:
  - commands - explanation of commands concept,
  - cursors - what they are and how they work,
  - keybindings - how to set, and what to watch out for,
  - selections - what ther are and how they work.`,

	// View

	"view":       "view line-number column # Sets view start at provided line number and column.",
	"view-down":  "view-down # Scrolls view down.",
	"view-left":  "view-down # Scrolls view left.",
	"view-right": "view-down # Scrolls view right.",
	"view-up":    "view-up # Scrolls view down.",

	/*
	   "match-brace":   struct{}{},
	   "match-bracket": struct{}{},
	   "match-paren":   struct{}{},
	*/
}
