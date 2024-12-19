package tab

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/lang"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mark"
	"github.com/m-kru/enix/internal/search"
	"github.com/m-kru/enix/internal/undo"
	"github.com/m-kru/enix/internal/util"
	"github.com/m-kru/enix/internal/view"
)

func Empty(
	keys *cfg.Keybindings,
	frame *frame.Frame,
) *Tab {
	lines := line.Empty()

	c := cursor.New(lines, 1, 0)
	curs := make([]*cursor.Cursor, 1, 16)
	curs[0] = c

	return &Tab{
		Keys:                 keys,
		Path:                 "No Name",
		Newline:              "\n",
		FileType:             "None",
		Indent:               cfg.Cfg.GetIndent(""),
		HasFocus:             true,
		State:                "",
		RepCount:             0,
		Lines:                lines,
		LineCount:            1,
		Cursors:              curs,
		Selections:           nil,
		InsertActions:        nil,
		PrevInsertCursors:    nil,
		PrevInsertSelections: nil,
		SearchCtx:            search.InitialContext(),
		Marks:                make(map[string]mark.Mark),
		Frame:                frame,
		View:                 view.View{Line: 1, Column: 1, Height: 1, Width: 1},
		Highlighter:          lang.DefaultHighlighter(),
		UndoStack:            undo.NewStack(cfg.Cfg.UndoSize),
		RedoStack:            undo.NewStack(cfg.Cfg.UndoSize),
		UndoCount:            0,
		RedoCount:            0,
		Prev:                 nil,
		Next:                 nil,
	}
}

// Open opens a new tab.
// It path is "", then new empty tab is opened.
//
// TODO: Allow opening without highlighter, useful for script mode.
func Open(
	keys *cfg.Keybindings,
	frame *frame.Frame,
	path string,
) (*Tab, error) {
	if path == "" {
		return Empty(keys, frame), nil
	}

	// Check existence of backup file. If exists, return an error.
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
	c := cursor.New(lines, 1, 0)
	curs := make([]*cursor.Cursor, 1, 16)
	curs[0] = c

	fileType := util.FileNameToType(filepath.Base(path))

	// Highlighter initialization
	hl, err := lang.NewHighlighter(fileType)

	return &Tab{
		Keys:                 keys,
		Path:                 path,
		Newline:              "\n",
		FileType:             fileType,
		Indent:               cfg.Cfg.GetIndent(fileType),
		HasFocus:             true,
		State:                "",
		RepCount:             0,
		Lines:                lines,
		LineCount:            lineCount,
		Cursors:              curs,
		Selections:           nil,
		InsertActions:        nil,
		PrevInsertCursors:    nil,
		PrevInsertSelections: nil,
		SearchCtx:            search.InitialContext(),
		Marks:                make(map[string]mark.Mark),
		Frame:                frame,
		View:                 view.View{Line: 1, Column: 1, Height: 1, Width: 1},
		Highlighter:          hl,
		UndoStack:            undo.NewStack(cfg.Cfg.UndoSize),
		RedoStack:            undo.NewStack(cfg.Cfg.UndoSize),
		UndoCount:            0,
		RedoCount:            0,
		Prev:                 nil,
		Next:                 nil,
	}, err
}

func FromString(
	keys *cfg.Keybindings,
	frame *frame.Frame,
	str string,
	path string,
) *Tab {
	lines, lineCount := line.FromString(str)

	c := cursor.New(lines, 1, 0)
	curs := make([]*cursor.Cursor, 1, 16)
	curs[0] = c

	return &Tab{
		Keys:                 keys,
		Path:                 path,
		Newline:              "\n",
		FileType:             "None",
		Indent:               cfg.Cfg.GetIndent(""),
		HasFocus:             true,
		State:                "",
		RepCount:             0,
		Lines:                lines,
		LineCount:            lineCount,
		Cursors:              curs,
		Selections:           nil,
		InsertActions:        nil,
		PrevInsertCursors:    nil,
		PrevInsertSelections: nil,
		SearchCtx:            search.InitialContext(),
		Marks:                make(map[string]mark.Mark),
		Frame:                frame,
		View:                 view.View{Line: 1, Column: 1, Height: 1, Width: 1},
		Highlighter:          lang.DefaultHighlighter(),
		UndoStack:            undo.NewStack(cfg.Cfg.UndoSize),
		RedoStack:            undo.NewStack(cfg.Cfg.UndoSize),
		UndoCount:            0,
		RedoCount:            0,
		Prev:                 nil,
		Next:                 nil,
	}
}
