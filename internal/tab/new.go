package tab

import (
	"errors"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"
	"io/ioutil"
	"os"
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
// It panics if "", then new empty tab is opened.
func Open(
	colors *cfg.Colorscheme,
	path string,
) *Tab {
	if path == "" {
		return Empty(colors)
	}

	t := &Tab{
		Colors:     colors,
		Name:       "",
		Path:       path,
		Newline:    "\n",
		FileType:   util.FileNameToType(path),
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
		bytes, err := ioutil.ReadFile(path)
		if err != nil {
			panic("unimplemented")
		}
		t.Lines = line.FromString(string(bytes))
	}

	// Cursor initialization
	c := &cursor.Cursor{Line: t.Lines}
	t.Cursors = c

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
