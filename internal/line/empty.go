package line

import (
	"github.com/m-kru/enix/internal/cfg"
)

func Empty(colors *cfg.Colorscheme) *Line {
	return &Line{
		Colors: colors,
		Buf:    "",
		Prev:   nil,
		Next:   nil,
	}
}
