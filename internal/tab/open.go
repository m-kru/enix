package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"

	"github.com/gdamore/tcell/v2"
)

func Empty(
	colors *cfg.Colorscheme,
	screen tcell.Screen,
	startX int,
	endX int,
	startY int,
	endY int,
) *Tab {
	t := &Tab{
		Colors:          colors,
		Screen:          screen,
		StartX:          startX,
		EndX:            endX,
		StartY:          startY,
		EndY:            endY,
		Name:            "No Name",
		Newline:         "\n",
		FileType:        "",
		HasChanges:      false,
		Lines:           line.Empty(colors),
		FirstVisLineIdx: 1,
	}

	c := &cursor.Cursor{
		Colors: colors,
		Screen: screen,
		Line:   t.Lines,
	}
	t.Cursor = c

	return t
}

// Open opens a new tab.
// If path is "", then new empty tab is opened.
func Open(
	colors *cfg.Colorscheme,
	screen tcell.Screen,
	startX int,
	endX int,
	startY int,
	endY int,
	path string,
	firstLine int, // First visible line
) *Tab {
	fileType := ""

	if path != "" {

	} else {
		path = "No Name"
	}

	return &Tab{
		Colors:          colors,
		Screen:          screen,
		StartX:          startX,
		EndX:            endX,
		StartY:          startY,
		EndY:            endY,
		Name:            path,
		Newline:         "\n",
		FileType:        fileType,
		HasChanges:      false,
		FirstVisLineIdx: firstLine,
	}
}
