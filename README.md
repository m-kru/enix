[![Tests](https://github.com/m-kru/enix/actions/workflows/tests.yml/badge.svg?branch=master)](https://github.com/m-kru/enix/actions?query=master)

**enix** is a terminal-based modal text editor trying to follow the \*nix philosophy.

# Features

- Whole documentation embedded into the binary.
  No need to access the Internet to get to know how things work.
- No runtime dependencies except optional configuration and syntax highlighting files.
- Multi-cursor support.
- Marks.
- Tabs with visible, clickable tab bar.
- Built-in help.
- Mouse support.
- Optional automatic removal of trailing whitespaces.
- Configurable character and style of tab and newline.
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

## Rune related

- rune length - number of bytes required to store a rune.
- rune width - number of screen cells required to display a rune.

## Index related

- buffer index - index of the byte in the line byte buffer.
- column index - index of the column in the tab.
- rune index - index of the rune in the line byte buffer.

## Execution related

- event - an external act caused by an implicit user activity, for example, mouse click or key press.
- command - a procedure executed by the editor in response to an event, for example, tab save.
- action - a change in the tab content made during command execution, for exmaple, rune insert or line deletion.

A usual chain of flow is as follows.
The user triggers an event, for example, by pressing a key on the keyboard.
The event is mapped to the specific command.
The command is executed, which leads to actions modifying the tab content.
However, if an event is not mapped to any command, then no command is executed.
Moreover, not all commands lead to actions.
For example, spawning a new cursor doesn't modify the tab content, so the command doesn't cause any actions.

# Commands

enix is build on the concept of built-in commands.
Commands are executed by stroking keybindings or by using the command prompt.
Commands are *line-oriented* or *selection-oriented*.

Line-oriented commands operate on whole lines.
Line-oriented commands by default operate on the line where the cursor is located.
However, if a selection is present, then line oriented commands operate on all lines spanned by the selection.
Even if the selection starts or ends in the middle of a line.

Selection-oriented commands operate on selections.
An example of selection-oriented command is `sh`, which forwards selection as a standard input to an external program.

# Default keybindings

To get the default keybindings simply run `enix -dump-keys` in the shell.

# Help

To get help message open the command prompt (`:` - by default) and type the `help` command.
To get help message for a particular command type `help <command-name>`.
