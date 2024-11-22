package tab

import (
	"errors"
	"fmt"
	"os"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/lang"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mark"
	"github.com/m-kru/enix/internal/undo"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"
)

func Empty(config *cfg.Config, colors *cfg.Colorscheme, keys *cfg.Keybindings) *Tab {
	lines := line.Empty()

	c := &cursor.Cursor{Line: lines, LineNum: 1, ColIdx: 1}
	curs := make([]*cursor.Cursor, 1, 16)
	curs[0] = c

	return &Tab{
		Config:      config,
		Colors:      colors,
		Keys:        keys,
		Path:        "No Name",
		Newline:     "\n",
		HasFocus:    true,
		Lines:       lines,
		LineCount:   1,
		Cursors:     curs,
		Marks:       make(map[string]mark.Mark),
		View:        view.View{Line: 1, Column: 1},
		Highlighter: lang.DefaultHighlighter(),
		UndoStack:   undo.NewStack(config.UndoSize),
		RedoStack:   undo.NewStack(config.UndoSize),
	}
}

// Open opens a new tab.
// It path is "", then new empty tab is opened.
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

	// Check existance of backup file. If exists, return an error.
	backupPath := path + ".enix-bak"
	_, err := os.Stat(backupPath)
	if err == nil {
		return nil, fmt.Errorf(
			"detected backup file '%[2]s'\n"+
				"resolve the issue manually, either:\n"+
				"1. Remove '%[2]s'\n"+
				"2. Replace '%[1]s' with '%[2]s'",
			path, backupPath,
		)
	}

	// Lines initialization
	lines, lineCount := line.FromString("")
	_, err = os.Stat(path)
	if errors.Is(err, os.ErrNotExist) {
		// Do notjing
	} else if err != nil {
		return nil, err
	} else {
		bytes, err := os.ReadFile(path)
		if err != nil {
			return nil, err
		}
		lines, lineCount = line.FromString(string(bytes))
	}

	// Cursor initialization
	c := &cursor.Cursor{Line: lines, LineNum: 1, ColIdx: 1}
	curs := make([]*cursor.Cursor, 1, 16)
	curs[0] = c

	fileType := util.FileNameToType(path)

	// Highlighter initialization
	hl, err := lang.NewHighlighter(fileType)

	return &Tab{
		Config:      config,
		Colors:      colors,
		Keys:        keys,
		Path:        path,
		Newline:     "\n",
		FileType:    fileType,
		HasFocus:    true,
		Lines:       lines,
		LineCount:   lineCount,
		Cursors:     curs,
		Marks:       make(map[string]mark.Mark),
		View:        view.View{Line: 1, Column: 1},
		Highlighter: hl,
		UndoStack:   undo.NewStack(config.UndoSize),
		RedoStack:   undo.NewStack(config.UndoSize),
	}, err
}

func FromString(
	config *cfg.Config,
	colors *cfg.Colorscheme,
	keys *cfg.Keybindings,
	str string,
	path string,
) *Tab {
	lines, lineCount := line.FromString(str)

	c := &cursor.Cursor{Line: lines, LineNum: 1, ColIdx: 1}
	curs := make([]*cursor.Cursor, 1, 16)
	curs[0] = c

	return &Tab{
		Config:      config,
		Colors:      colors,
		Keys:        keys,
		Path:        path,
		Newline:     "\n",
		FileType:    "None",
		HasFocus:    true,
		Lines:       lines,
		LineCount:   lineCount,
		Cursors:     curs,
		Marks:       make(map[string]mark.Mark),
		View:        view.View{Line: 1, Column: 1},
		Highlighter: lang.DefaultHighlighter(),
		UndoStack:   undo.NewStack(config.UndoSize),
		RedoStack:   undo.NewStack(config.UndoSize),
	}
}
