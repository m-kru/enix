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

func (f Frame) GetContent(x int, y int) rune {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= f.Width (%d)", x, f.Width))
	}
	if y >= f.Height {
		panic(fmt.Sprintf("y (%d) >= f.Height (%d)", y, f.Height))
	}

	r, _, _, _ := f.Screen.GetContent(x+f.X, y+f.Y)

	return r
}

func (f Frame) SetContent(x int, y int, r rune, style tcell.Style) {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= f.Width (%d)", x, f.Width))
	}
	if y >= f.Height {
		panic(fmt.Sprintf("y (%d) >= f.Height (%d)", y, f.Height))
	}

	f.Screen.SetContent(x+f.X, y+f.Y, r, nil, style)
}

// Line returns frame f subframe for line rendering.
func (f Frame) Line(x int, y int) Frame {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= f.Width (%d)", x, f.Width))
	}
	if y >= f.Height {
		panic(fmt.Sprintf("y (%d) >= f.Height (%d)", y, f.Height))
	}

	return Frame{
		Screen: f.Screen,
		X:      f.X + x,
		Y:      f.Y + y,
		Width:  f.Width - x,
		Height: 1,
	}
}

// LastLine returns subframe containing only the last line of frame f.
func (f Frame) LastLine() Frame {
	if f.Height > 1 {
		f.Y = f.Y + f.Height - 1
		f.Height = 1
	}
	return f
}

func (f Frame) Column(x int, width int) Frame {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= frame.Width (%d)", x, f.Width))
	}
	if x+width > f.Width {
		panic(fmt.Sprintf("x (%d) + width (%d) > f.Width %d", x, width, f.Width))
	}

	return Frame{
		Screen: f.Screen,
		X:      x,
		Y:      f.Y,
		Width:  width,
		Height: f.Height,
	}
}
