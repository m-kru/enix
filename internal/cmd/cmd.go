package cmd

var cmdDescriptions = map[string]string{
	// Command

	"cmd": `cmd # Start command prompt.`,
	"cmd-error": `cmd-error # Displays error message in the command prompt.
This command is useful for enix debugging and development.`,
	"cmd-info": `cmd-info # Displays message in the command prompt.
This command is useful for enix debugging and development.`,
	"cmd-list": `cmd-list # Lists all available commands in alphabetical order with one sentence summary.`,
	"cmd-prev": `cmd-prev # Executes previous command.`,

	// Cursor

	"cursor-down": `cursor-down # Moves cursor down.
If cursor is in the last line, nothing happens.`,

	"cursor-left": `cursor-left # Moves cursors left.
If cursor is in the first column of a line, then it is moved into the last
column of the previous line. Unless this is the first line. In such a case,
nothing happens.`,

	"cursor-right": `cursor-right # Moves cursors right.
If cursor is in the last column of a line, then it is moved into the first
column of the next line. Unless this is the lastt line. In such a case,
nothing happens.`,

	"cursor-up": `cursor-up # Moves cursors up.
If cursor is in the fisrt line, nothing happens.`,

	// Deletion

	"del":      `del # Deletes text under cursors/selections.`,
	"del-word": `del-word # Deletes words under cursors/selections.`,
	"del-line": `del-line # Deletes lines with cursors/selections.`,

	// File

	"file-open": `file-open [path/to/file] # Opens file in a new tab.
If path to file is not provided, then it behaves the same as the
tab-open command.`,

	"file-save": `file-save [path/to/file] # Saves file using the provided path to file.
If path to file is not provided, then it uses the current file path.`,

	"file-type": `file-type type # Enforces file type.`,

	// Indent

	"indent-increase": `indent-increase # Increases indent of lines with cursor.
In the case of selections, it increases all selected lines. Even if the
selection starts/ends in the middle of a line.`,

	"indent-decrease": `indent-increase # Decreases indent of lines with cursor.
In the case of selections, it decreases all selected lines. Even if the
selection starts/ends in the middle of a line.`,

	// Tab

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

	"escape": `escape # Escapes the current context.
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
	/*
	   "cursor-down-spawn":    struct{}{},
	   "cursor-up-spawn":      struct{}{},
	   "cursor-match-brace":   struct{}{},
	   "cursor-match-bracket": struct{}{},
	   "cursor-match-paren":   struct{}{},
	*/
}

// IsValid returns true if given command is a valid command.
func IsValid(cmd string) bool {
	if _, ok := cmdDescriptions[cmd]; ok {
		return true
	}
	return false
}
