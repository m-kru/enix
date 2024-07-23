package cmd

var topicKeybindings = `Keybindings

enix tries to offer sane default keybindings so that the user doesn't have to
configure much out of the box and can use the editor straight after installation
on various machines.

The default keybindings follows the following rules (with some exceptions):
- Keybindings with the Ctrl modifier are used to control things. For example,
  control cursor position, spawn new cursors, open or save file, open or close
  tabs, text selections.
- Keybindings with the Alt modifier are used to alternate the text buffer.
  For example, move lines up or down, increase or decrease the indent.
- Keybindings with the Shift modifier are used to extend or shirnk selections.

However, some keybindings with the Ctrl modifier are so omnipresent and have
unified behavior that enix follows them and has a few exceptions from its
default keybinding rules.
- Ctrl+V - paste text into the text buffer,
- Ctrl+X - cut selections/line.

The following default Ctrl keybindings also follow common conventions.
However, they are not exceptions as they don't alternate the text buffer:
- Ctrl+A - select whole text buffer,
- Ctrl+C - copy text,
- Ctrl+F - find text,
- Ctrl+O - open file,
- Ctrl+Q - quit tab/enix,
- Ctrl+S - save file,
- Ctrl+T - open new tab.

Risky keybindings

enix receives key events sent from a terminal emulator, not from the keyboard.
As terminal emulators are weird and non-uniform applications they eat or
replace some key events. This is why some keybindings with the Ctrl modifier
never reach enix.

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

Please note, that all window manager or custom terminal emulator keybindings
also won't reach enix, as they are simply consumed by your window manager or
terminal emulator.
`
