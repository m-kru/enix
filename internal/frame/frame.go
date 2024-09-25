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

func (f Frame) HideCursor() { f.Screen.HideCursor() }

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

func (f Frame) ShowCursor(x, y int) {
	if x >= f.Width {
		panic(fmt.Sprintf("x (%d) >= f.Width (%d)", x, f.Width))
	}
	if y >= f.Height {
		panic(fmt.Sprintf("y (%d) >= f.Height (%d)", y, f.Height))
	}

	f.Screen.ShowCursor(x+f.X, y+f.Y)
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

// Within returns true if cell with x and y coordinates is located within the frame.
func (f Frame) Within(x, y int) bool {
	if f.X <= x && x < f.X+f.Width && f.Y <= y && y < f.Y+f.Height {
		return true
	}
	return false
}

// ToFramePosition transforms screen position with coordinates x, y to frame position.
// It is user's responsibility to make sure the initial point is within the frame.
func (f Frame) ToFramePosition(x, y int) (int, int) {
	return x - f.X, y - f.Y
}
