package cmd

var cmdDescriptions = map[string]string{
	// Cursor
	"cursor-down": `Usage:
  cursor-down
Description:
  Moves cursors down. If cursor is in the last line, nothing happens.`,

	"cursor-left": `Usage:
  cursor-left
Description:
  Moves cursors left. If cursor is in the first column of a line, then it is
  moved into the last column of the previous line. Unless this is the first line.
  In such a case, nothing happens.`,
	"cursor-right": `Usage:

  cursor-right
Description:
  Moves cursors right. If cursor is in the last column of a line, then it is
  moved into the first column of the next line. Unless this is the lastt line.
  In such a case, nothing happens.`,

	"cursor-up": `Usage:
  cursor-up
Description:
  Moves cursors up. If cursor is in the fisrt line, nothing happens.`,

	// File

	"file-open": `Usage:
  file-open [path/to/file]
Description:
  Opens file in a new tab. If path to file is not provided, then it behaves
  the same as the tab-open command.`,

	"file-save": `Usage:
  file-save [path/to/file]
Description:
  Saves file using the provided path to file. If path to file is not provided,
  then it uses the current file path.`,

	// Tab

	"tab-open": `Usage:
  tab-open
Description:
  Opens a new empty tab.`,

	"tab-next": `Usage:
  tab-next
Description:
  Cycles to the next tab. If the current tab is the last tab, then it wraps
  to the first tab.`,

	"tab-prev": `Usage:
  tab-prev
Description:
  Cycles to the previous tab. If the current tab is the first tab, then it wraps
  to the last tab.`,

	// Miscellaneous

	"escape": `Usage:
  escape
Description:
  Escapes current context. The actual action depends on the context.
  For example, if the focues is on the command prompt, then escape command
  escapes the command prompt and allows user to continue file editing.
  If the focus in on a file editing and there is any selection, then escape
  command clears all selections.`,

	"help": `Usage:
  help [topic|command-name]
Description:
  Displays help message for a given topic or command in a newly open tab.
  If neither topic nor command name is provided displays help message
  for the help command. The same as 'help help'. Valid topics are:
    - XXX - lorem ipsum ...
    - YYY - lorem ipsum ...
    - ZZZ - lorem ipsum ...`,
	/*
	   "cursor-down-spawn":    struct{}{},
	   "cursor-up-spawn":      struct{}{},
	   "cursor-match-brace":   struct{}{},
	   "cursor-match-bracket": struct{}{},
	   "cursor-match-paren":   struct{}{},
	*/
}

// IsValid returns true if given command is a valid command.
func IsValid(cmd string) {
	if _, ok := cmdDescriptions[cmd]; ok {
		return true
	}
	return false
}
