package tab

import (
	"errors"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/lang"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mark"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"
	"os"
)

func Empty(config *cfg.Config, colors *cfg.Colorscheme, keys *cfg.Keybindings) *Tab {
	tab := &Tab{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		Path:       "No Name",
		Newline:    "\n",
		FileType:   "",
		HasFocus:   true,
		HasChanges: false,
		Lines:      line.Empty(),
		Marks:      make(map[string]mark.Mark),
		View:       view.View{Line: 1, Column: 1},
	}

	c := &cursor.Cursor{Config: config, Line: tab.Lines}
	tab.Cursors = c

	hl := lang.DefaultHighlighter()
	tab.Highlighter = &hl

	return tab
}

// Open opens a new tab.
// It panics if "", then new empty tab is opened.
//
// TODO: Allow opening without highlighter, useful for script mode.
func Open(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	keys *cfg.Keybindings,
	path string,
) (*Tab, error) {
	if path == "" {
		return Empty(config, colors, keys), nil
	}

	tab := &Tab{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		Path:       path,
		Newline:    "\n",
		FileType:   util.FileNameToType(path),
		HasFocus:   true,
		HasChanges: false,
		Marks:      make(map[string]mark.Mark),
		View:       view.View{Line: 1, Column: 1},
	}

	// Lines initialization
	_, err := os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		tab.Lines = line.FromString("")
	} else if err != nil {
		return nil, err
	} else {
		bytes, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		tab.Lines = line.FromString(string(bytes))
	}

	// Cursor initialization
	c := &cursor.Cursor{Config: config, Line: tab.Lines}
	tab.Cursors = c

	// Highlighter initialization
	hl, err := lang.NewHighlighter(tab.FileType)
	tab.Highlighter = &hl

	return tab, err
}

func FromString(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	keys *cfg.Keybindings,
	str string,
	path string,
) *Tab {
	tab := &Tab{
		Config:     config,
		Colors:     colors,
		Keys:       keys,
		Path:       path,
		Newline:    "\n",
		FileType:   "None",
		HasFocus:   true,
		HasChanges: false,
		Lines:      line.FromString(str),
		Marks:      make(map[string]mark.Mark),
		View:       view.View{Line: 1, Column: 1},
	}

	c := &cursor.Cursor{Config: config, Line: tab.Lines}
	tab.Cursors = c

	hl := lang.DefaultHighlighter()
	tab.Highlighter = &hl

	return tab
}
