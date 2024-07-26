package frame

import (
	"fmt"

	"github.com/gdamore/tcell/v2"
)

type Frame struct {
	Screen tcell.Screen
	X      int // Start X coordinate
	Y      int // Start Y coordinate
	Width  int
	Height int
}

func (f Frame) SetContent(x int, y int, r rune, style tcell.Style) {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= frame.Width (%d)", x, f.Width))
	}
	if y >= f.Height {
		panic(fmt.Sprintf("y (%d) >= frame.Height (%d)", y, f.Height))
	}

	f.Screen.SetContent(x+f.X, y+f.Y, r, nil, style)
}

// Line returns frame f subframe for line rendering.
func (f Frame) Line(x int, y int) Frame {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= frame.Width (%d)", x, f.Width))
	}
	if y >= f.Height {
		panic(fmt.Sprintf("y (%d) >= frame.Height (%d)", y, f.Height))
	}

	return Frame{
		Screen: f.Screen,
		X:      f.X + x,
		Y:      f.Y + y,
		Width:  f.Width - x,
		Height: 1,
	}
}
