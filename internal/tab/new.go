package tab

import (
	"errors"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"
	"os"
)

func Empty(config *cfg.Config, colors *cfg.Colorscheme, keys *cfg.Keybindings) *Tab {
	t := &Tab{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		Path:       "No Name",
		Newline:    "\n",
		FileType:   "",
		HasFocus:   true,
		HasChanges: false,
		Lines:      line.Empty(),
		View:       view.View{Line: 1, Column: 1},
	}

	c := &cursor.Cursor{Config: config, Line: t.Lines}
	t.Cursors = c

	return t
}

// Open opens a new tab.
// It panics if "", then new empty tab is opened.
func Open(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	keys *cfg.Keybindings,
	path string,
) *Tab {
	if path == "" {
		return Empty(config, colors, keys)
	}

	t := &Tab{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		Path:       path,
		Newline:    "\n",
		FileType:   util.FileNameToType(path),
		HasFocus:   true,
		HasChanges: false,
		View:       view.View{Line: 1, Column: 1},
	}

	// Lines initialization
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		t.Lines = line.FromString("")
	} else if err != nil {
		panic("unimplemented")
	} else {
		bytes, err := os.ReadFile(path)
		if err != nil {
			panic("unimplemented")
		}
		t.Lines = line.FromString(string(bytes))
	}

	// Cursor initialization
	c := &cursor.Cursor{Config: config, Line: t.Lines}
	t.Cursors = c

	return t
}

func FromString(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	keys *cfg.Keybindings,
	str string,
	path string,
) *Tab {
	t := &Tab{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		Path:       path,
		Newline:    "\n",
		FileType:   "None",
		HasFocus:   true,
		HasChanges: false,
		Lines:      line.FromString(str),
		View:       view.View{Line: 1, Column: 1},
	}

	c := &cursor.Cursor{Config: config, Line: t.Lines}
	t.Cursors = c

	return t
}
