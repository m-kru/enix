[![Tests](https://github.com/m-kru/enix/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/m-kru/enix/actions?query=master)

**enix** is a terminal-based text editor trying to follow the \*nix philosophy.

# Features

- Whole documentation embedded into the binary.
  No need to access the Internet to get to know how things work.
- No runtime dependencies except configuration and syntax highlighting files.
- Multi-cursor support.
- Tabs.
- Built-in help.
- Mouse support.
- Improved whitespace management.
  - Automatic removal of trailing whitespaces.
  - Highlighting invalid indent whitespace (tab or space).
- Configuration and syntax highlighting loaded dynamically.
  No need to rebuilt or wait for weeks/months to get syntax highlighting fix or improvement.
  Easy to use custom syntax highlighting.

# No-features

- No integrated terminal.
  Simply `Ctrl-z` to execute shell commands.
- No splits.
  Instead, simply do one of the following:
  -  open multiple terminals and use your window manager capabilities,
  -  use a [tmux](https://github.com/tmux/tmux) like program,
  -  use a [tilix](https://github.com/gnunn1/tilix) like terminal.
- No embedded plugin system.
  Plugins can be written as regular programs in any programming language.
  Plugins can be run using the `sh` command.
  Each selection is a standard input for an externally called program and is replaced by the program standard output.

# Glossary

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

## Cursor

- `Down` - cursor-down
- `Left` - cursor-left
- `Right` - cursor-right
- `Up` - cursor-up
- `Ctrl-Left` - cursor-
- `Ctrl-Right` - cursor-
- `Ctrl-Down` - cursor-down-spawn
- `Ctrl-Up` - cursor-up-spawn
- `Ctrl-{` - cursor-match-brace
- `Ctrl-[` - cursor-match-bracket
- `Ctrl-(` - cursor-match-paren

## Tab

- `Ctrl-t` - tab-open
- `Ctrl-w` - tab-close
- `Ctrl-Tab` - tab-next

## Find

- `Ctrl-f` - find
- ` ` - find-next
- ` ` - find-prev

## Selection

- `Ctrl-a` - select-all
- `Ctrl-l` - select-line
- `Shift-Right` - select-next-char
- `Shift-Right` - select-prev-char
- `Ctrl-}` - select-in-brace
- `Ctrl-]` - select-in-bracket
- `Ctrl-)` - select-in-paren

## File

- `Ctrl-o` - file-open
- `Ctrl-s` - file-save
-
## Miscellaneous

- `Ctrl-c` - copy
- `Ctrl-k` - cut-line
- `Ctrl-p` - paste
- `Ctrl-q` - quit

# Help

To get help message open the command prompt (` ` - by default) and type the `help` command.
To get help message for a particular command type `help <command-name>`.
