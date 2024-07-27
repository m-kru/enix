package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"

	"github.com/gdamore/tcell/v2"
)

func Empty(colors *cfg.Colorscheme, screen tcell.Screen) *Tab {
	t := &Tab{
		Colors:     colors,
		Screen:     screen,
		Name:       "No Name",
		Newline:    "\n",
		FileType:   "",
		HasChanges: false,
		Lines:      line.Empty(),
	}

	t.FirstVisLine = t.Lines
	t.LastVisLine = t.Lines

	c := &cursor.Cursor{
		Colors: colors,
		Line:   t.Lines,
	}
	t.Cursors = c

	return t
}

// Open opens a new tab.
// If path is "", then new empty tab is opened.
func Open(
	colors *cfg.Colorscheme,
	screen tcell.Screen,
	path string,
	firstLine int, // First visible line
) *Tab {
	fileType := ""

	if path != "" {

	} else {
		path = "No Name"
	}

	t := &Tab{
		Colors:     colors,
		Screen:     screen,
		Name:       path,
		Newline:    "\n",
		FileType:   fileType,
		HasChanges: false,
	}

	t.FirstVisLine = t.Lines
	t.LastVisLine = t.Lines

	return t
}

func FromString(
	colors *cfg.Colorscheme,
	screen tcell.Screen,
	str string,
	name string,
) *Tab {
	t := &Tab{
		Colors:     colors,
		Screen:     screen,
		Name:       name,
		Newline:    "\n",
		FileType:   "",
		HasChanges: false,
		Lines:      line.FromString(str),
	}

	t.FirstVisLine = t.Lines
	t.LastVisLine = t.Lines

	c := &cursor.Cursor{
		Colors: colors,
		Line:   t.Lines,
	}
	t.Cursors = c

	return t
}
