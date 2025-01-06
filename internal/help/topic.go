package help

var Topics = map[string]string{
	"enix": `Enix is a terminal-based modal text editor trying to follow the *nix philosophy.

Enix is created with simplicity and orthogonality in mind.
For example, there is only one way to enter the insert state.
Namely, the 'insert' commamnd (and its derivatives), there is no 'append'.

There is no custom plugin system.
However, it is possible to pass selections as standard input to external programs.
The selections are then replaced with the program standard output.
This is achieved using the 'sh' command.
Enter 'help sh' command for more details.
`,

	"cursors": `A tab at any time has at least one cursor or selection.
The number of cursors is unlimited.
Each cursor points to a line and rune index or line end.
The line end rune is displayed in the tab.
However, it is not part of the tab content.
It solely symbolizes the line end, not the newline.

At no point more than one cursor can point to the same position.
If, due to command execution, multiple cursors point to the same position, some of them are desroyed so that only one cursor pointing to a particular position exists.

The status bar displays position of the cursor on the right side in the following format: [line number]:[column].
In the case of multiple cursors, position of the last cursor is displayed.
The last cursor is the cursor which was created as the last one.
It is not necessarily cursor with the greatest line number or column position.

In the case of escaping from multiple cursors by executing the 'esc' command, it is only the last cursor that survives.`,

	"keybindings": `There are three independent keybinding sets.
Each of them is used in a different context:
1. keys - keybindings for a tab in the normal mode,
2. insert keys - keybindings for a tab in the insert mode,
3. prompt keys - keybindings for the prompt.

Enix tries to offer sane default keybindings so the user doesn't have to configure much out of the box.
The user can start using the editor right after installing it on various machines.

To get currently set keybindings, call enix with one of the following flags:
'-dump-keys', '-dump-insert-keys', or '-dump-prompt-keys'.

The default keybindings try to follow the following rules (with some exceptions):
- The most common commands are under primary keys
  The actions associated with them can be executed with a single keystroke.
- Keybindings with the Ctrl modifier are used to control things.
  For example, spawn new cursors, scroll view, create or iterate tabs.
- Keybindings with the Alt modifier are used to alternate the text.
  For example, join or move lines up or down.
- Keybindings with the Shift modifier are used to extend or shrink selections.

Setting custom keybindings

The user can overwrite default keybindings or define new custom ones.
To do so, one has to provide the following files in the enix configuration directory:
1. keys.json - custom keybindings for the tab normal mode,
2. insert-keys.json - custom keybindings for the tab insert mode,
3. prompt-keys.json - custom keybindings for the prompt.

These files shall contain a single JSON object representing a keybinding map.
A single key is a command name, and a single value is a key combination name.
A good starting point for defining custom keybindings might be dumping keys first.

Getting key names

Getting some key names might not be obvious.
For such cases, enix provides the 'key-name' command.
To get more details, open enix and execute the 'help key-name' command.

Risky keybindings

Enix receives key events from a terminal emulator, not from the keyboard.
As terminal emulators are weird and non-uniform applications, they eat or replace some key events.
This is why some keybindings with the Ctrl modifier never reach enix.

The following Ctrl keybindings are considered risky and are advised to be avoided.
- Ctrl+H - Backspace,
- Ctrl+I - Tab,
- Ctrl+M - Enter,
- Ctrl+; - some weird character,
- Ctrl+/ - some weird action,
- Ctrl+` + "`" + ` - Ctrl+Space,
- Ctrl+Tab - Tab,
- Ctrl+Enter - Enter,
- Ctrl+[1|2|3|4|5|6|7|8|9|0] - runes and control sequences.

Please note that all window manager or custom terminal emulator keybindings also won't reach enix.
The window manager or terminal emulator consumes them.`,
}
