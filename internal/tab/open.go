package tab

import (
	"github.com/m-kru/enix/internal/cfg"

	"github.com/gdamore/tcell/v2"
)

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
) *Tab {
	fileType := ""

	if path != "" {

	} else {
		path = "No Name"
	}

	return &Tab{
		Screen:     screen,
		StartX:     startX,
		EndX:       endX,
		StartY:     startY,
		EndY:       endY,
		Colors:     colors,
		Name:       path,
		Newline:    "\n",
		FileType:   fileType,
		HasChanges: false,
	}
}
