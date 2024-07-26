package tab

import (
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/line"

	"github.com/gdamore/tcell/v2"
)

func Empty(
	screen tcell.Screen,
	startX int,
	endX int,
	startY int,
	endY int,
	colors *cfg.Colorscheme,
) *Tab {
	t := &Tab{
		Screen:          screen,
		StartX:          startX,
		EndX:            endX,
		StartY:          startY,
		EndY:            endY,
		Colors:          colors,
		Name:            "No Name",
		Newline:         "\n",
		FileType:        "",
		HasChanges:      false,
		Lines:           line.Empty(screen, colors),
		FirstVisLineIdx: 1,
	}

	c := &cursor.Cursor{
		Screen: screen,
		Colors: colors,
		Line:   t.Lines,
	}
	t.Cursor = c

	return t
}

// Open opens a new tab.
// If path is "", then new empty tab is opened.
func Open(
	screen tcell.Screen,
	startX int,
	endX int,
	startY int,
	endY int,
	colors *cfg.Colorscheme,
	path string,
	firstLine int, // First visible line
) *Tab {
	fileType := ""

	if path != "" {

	} else {
		path = "No Name"
	}

	return &Tab{
		Screen:          screen,
		StartX:          startX,
		EndX:            endX,
		StartY:          startY,
		EndY:            endY,
		Colors:          colors,
		Name:            path,
		Newline:         "\n",
		FileType:        fileType,
		HasChanges:      false,
		FirstVisLineIdx: firstLine,
	}
}
