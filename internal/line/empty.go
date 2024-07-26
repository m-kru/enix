package line

import (
	"github.com/gdamore/tcell/v2"
	"github.com/m-kru/enix/internal/cfg"
)

func Empty(screen tcell.Screen, colors *cfg.Colorscheme) *Line {
	return &Line{
		Screen: screen,
		Colors: colors,
		Buf:    "",
		Prev:   nil,
		Next:   nil,
	}
}
