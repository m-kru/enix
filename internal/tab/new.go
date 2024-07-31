package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/view"
)

func Empty(colors *cfg.Colorscheme) *Tab {
	t := &Tab{
		Colors:     colors,
		Name:       "No Name",
		Path:       "",
		Newline:    "\n",
		FileType:   "",
		HasChanges: false,
		Lines:      line.Empty(),
		View:       view.View{Line: 1, Column: 1},
	}

	c := &cursor.Cursor{Line: t.Lines}
	t.Cursors = c

	return t
}

// Open opens a new tab.
// If path is "", then new empty tab is opened.
func Open(
	colors *cfg.Colorscheme,
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
		Name:       "",
		Path:       path,
		Newline:    "\n",
		FileType:   fileType,
		HasChanges: false,
	}

	return t
}

func FromString(
	colors *cfg.Colorscheme,
	str string,
	name string,
) *Tab {
	t := &Tab{
		Colors:     colors,
		Name:       name,
		Path:       "",
		Newline:    "\n",
		FileType:   "None",
		HasChanges: false,
		Lines:      line.FromString(str),
		View:       view.View{Line: 1, Column: 1},
	}

	c := &cursor.Cursor{Line: t.Lines}
	t.Cursors = c

	return t
}
