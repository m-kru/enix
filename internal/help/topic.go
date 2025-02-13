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

	"config": `While starting enix reads the following configuration files:
  1. 'config.json' - general configuration file,
  2. 'keys.json' - keybindings in the normal mode,
  3. 'keys-insert.json' - keybindings in the insert mode,
  4. 'keys-prompt.json' - keybidnings for the prompt,
  5. 'colors/<colors>.json' - colors definition,
  6. 'style/<style>.json' - style definition.

The <colors> and <style> can be defined in the 'config.json' file using the "Colors" and "Style" keys.
If they are not defined, then the default colors (termianl colors) and the default style (style built into binary) are used.

How to set custom keybindings is described in the keybindings topic.
Please check 'help keybindings'.

During the runtime, enix also reads files containing definitions of the syntax highlighting for particular filetypes.
This happens every time the user opens a new tab, or implicitly changes filetype of a given tab using the 'filetype' command.
Filetype files are looked for under the following path, 'filetype/<filetype>.json'.
How to define highlighting for a new filetype is described in the highlighting topic.
Please check 'help highlighting'.

Enix looks for configuration files in the following directories:
  1. A path defined in the 'ENIX_CONFIG_DIR' environment variable.
  2. 'enix' directory within the path defined in the 'XDG_CONFIG_HOME' environment variable.
     Bear in mind, that if 'XDG_CONFIG_HOME' is not defined, then it defaults to the '$HOME/.config'.
  3. Enix default config installation path:
     - '/usr/local/share/enix' for *nix-like systems.
If the requested configuration file is not found in the searched directory, then enix proceeds to the next directory.
If the requested configuration file is not found in any configuration directory, then enix uses the default configuration built into the binary.

Colorscheme

In enix, the colorscheme is composed of two independent logical elements: colors definition and style definition.
Users can arbitrarily mix colors and styles.
However, a particular style is usually defined to go in hand with a particular colors definition.
Mixing arbitrarily colors and styles can lead to unreadable colorscheme.

General configuration

The general configuration can be defined in the 'config.json' file.
To get all configuration settings with their current value call enix with the '-dump config' flag.

This section tries to describe in detail all general configuration settings.
Settings are described in the following format:

name : type
  description

AutoSave : int
  Autosave value in seconds.
  The value must be natural.
  If value equals 0, then the autosave feature is disabled.
  The tab is autosaved every n seconds, not n seconds after the last rune insert.
  Setting low n value on a constrained system may lead to performance drop.
  The tab is not autosaved if the corresponding file in the file system doesn't exist.

Colors : string
  Name of the file with colors definition.
  Enix tries to read colors from the 'colors/<colors>.json' file.

Extensions : map[string]string
  Custom file extensions.
  This setting is useful for enabling highlighting in files with custom extensions.
  The key in the map is a custom file extension.
  The value in the map is a filetype recognized by enix.

HighlightCursorWord : bool
  Highlight word under cursor and all occurrence of the word in the view.

Indent : map[string]string
  Indent configuration for various filetypes.
  If filetype is not recognized, then the default indent is a signle tab.
  The key in the map is the filetype.
  The value in the map is a string representing indent for the provided filetype.

LineEndRune : int
  Line end rune.
  Unfortunately, the value type is int, so the rune is not visible in the config file.

TabPadRune : int
  Tab padding rune.
  Unfortunately, the value type is int, so the rune is not visible in the config file.

TabRune : int
  Tab rune.
  Unfortunately, the value type is int, so the rune is not visible in the config file.

TrimOnSave : bool
  Trim trailing whitespaces while saving the tab.
  TrimOnSave doesn't apply for autosaves.

SafeFileSave : bool
  Automatically create file backup before save.
  This helps to avoid file loss if there is a power failure during the save.
  After successful save, the backup file is removed.
  The backup file has an additional '.enix-bak' extension.

Style : string
  Name of the file with style definition.
  Enix tries to read style from the 'style/<style>.json' file.

UndoSize : int
  Size of the undo and redo stacks.
  The value must be natural.
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

	"highlighting": `

Use region instead of regex with lookarounds when at least one of the following conditions is met:
  1. Highlighted text might span more than one line.
  2. Highlighted text might have inner highlights with different style.

Regex with lookarounds is also better than the region in terms of performance.
This is because correctly splitting a tab into regions requires analyzing the whole tab content.
Whereas regex with lookarounds is analyzed only for currently visible lines.`,

	"keybindings": `There are three independent keybinding sets.
Each of them is used in a different context:
  1. keys - keybindings for a tab in the normal mode,
  2. insert keys - keybindings for a tab in the insert mode,
  3. prompt keys - keybindings for the prompt.

Enix tries to offer sane default keybindings so the user doesn't have to configure much out of the box.
The user can start using the editor right after installing it on various machines.

To get currently set keybindings, call enix with one of the following flags:
'-dump-keys', '-dump-keys-insert', or '-dump-keys-prompt'.

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
  1. 'keys.json' - custom keybindings for the tab normal mode,
  2. 'keys-insert.json' - custom keybindings for the tab insert mode,
  3. 'keys-prompt.json' - custom keybindings for the prompt.

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
