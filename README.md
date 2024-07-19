**enix** is a terminal-based text editor trying to follow the \*nix philosophy.

# Features

- Multi-cursor support.
- Tabs.

# No-features

- No integrated terminal.
  Simply `Ctrl-z` to execute shell commands.
- No splits.
  Just open multiple terminals and use your window manager capabilities or use a tmux like program.
- No embedded plugin system.
  Plugins can be written as regular programs in any programming language.
  Plugins can be run using the `sh` command.
  Each selection is a standard input for an externally called program and is replaced by the program standard output.

# Commands

enix is build on the concept of built-in commands.
Commands are executed by stroking keybindings or by using the command prompt.
Commands are *line-oriented* or *selection-oriented*.

Line-oriented commands operate on whole lines.
Line-oriented commands be default operate on the line where the cursor is located.
However, if a selection is present, then line oriented commands operate on all lines spanned by the selection.
Even if the selection starts or ends in the middle of a line.
An example of line-oriented command is `duplicate-line`.

Selection-oriented commands operate on selections.
An example of selection-oriented command is `sh`, which forwards selection as a standard input to an external program.

# Default keybindings

## Cursor Movement

- `Down` - cursor-down
- `Left` - cursor-left
- `Right` - cursor-right
- `Up` - cursor-up
- `Ctrl-Left` -
- `Ctrl-Right` -

## Tab

- `Ctrl-t` - open-tab
- `Ctrl-w` - close-tab
- `Ctrl-Tab` - next-tab

## Find

- `Ctrl-f` - find
- ` ` - find-next
- ` ` - find-prev

## Selection

- `Ctrl-a` - select-all
- `Ctrl-l` - select-line
- `Shift-Right` - select-next-char
- `Shift-Right` - select-prev-char

## Miscellaneous

- `Ctrl-c` - copy
- `Ctrl-k` - cut-line
- `Ctrl-o` - open-file
- `Ctrl-p` - paste
- `Ctrl-q` - quit
- `Ctrl-s` - save

# Help

To get help message open the command prompt (` ` - by default) and type the `help` command.
To get help message for a particular command type `help <command-name>`.
