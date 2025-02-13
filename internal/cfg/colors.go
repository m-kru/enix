package cfg

import (
	"strconv"

	"github.com/gdamore/tcell/v2"
)

type Colors struct {
	Background    tcell.Color
	Foreground    tcell.Color
	Black         tcell.Color
	Red           tcell.Color
	Green         tcell.Color
	Yellow        tcell.Color
	Blue          tcell.Color
	Magenta       tcell.Color
	Cyan          tcell.Color
	White         tcell.Color
	BrightBlack   tcell.Color
	BrightRed     tcell.Color
	BrightGreen   tcell.Color
	BrightYellow  tcell.Color
	BrightBlue    tcell.Color
	BrightMagenta tcell.Color
	BrightCyan    tcell.Color
	BrightWhite   tcell.Color
}

type ColorsJSON struct {
	Background    string
	Foreground    string
	Black         string
	Red           string
	Green         string
	Yellow        string
	Blue          string
	Magenta       string
	Cyan          string
	White         string
	BrightBlack   string
	BrightRed     string
	BrightGreen   string
	BrightYellow  string
	BrightBlue    string
	BrightMagenta string
	BrightCyan    string
	BrightWhite   string
}

func stringToColor(str string) (tcell.Color, error) {
	var col tcell.Color

	u64, err := strconv.ParseUint(str, 16, 64)
	if err != nil {
		return col, err
	}

	col = tcell.Color(u64) | tcell.ColorIsRGB

	return col, err
}

func (cj ColorsJSON) ToColors() (Colors, error) {
	var cols Colors
	var err error
	var c tcell.Color

	if c, err = stringToColor(cj.Background); err != nil {
		return cols, err
	}
	cols.Background = c

	if c, err = stringToColor(cj.Foreground); err != nil {
		return cols, err
	}
	cols.Foreground = c

	if c, err = stringToColor(cj.Black); err != nil {
		return cols, err
	}
	cols.Black = c

	if c, err = stringToColor(cj.Red); err != nil {
		return cols, err
	}
	cols.Red = c

	if c, err = stringToColor(cj.Green); err != nil {
		return cols, err
	}
	cols.Green = c

	if c, err = stringToColor(cj.Yellow); err != nil {
		return cols, err
	}
	cols.Yellow = c

	if c, err = stringToColor(cj.Blue); err != nil {
		return cols, err
	}
	cols.Blue = c

	if c, err = stringToColor(cj.Magenta); err != nil {
		return cols, err
	}
	cols.Magenta = c

	if c, err = stringToColor(cj.Cyan); err != nil {
		return cols, err
	}
	cols.Cyan = c

	if c, err = stringToColor(cj.White); err != nil {
		return cols, err
	}
	cols.White = c

	if c, err = stringToColor(cj.BrightBlack); err != nil {
		return cols, err
	}
	cols.BrightBlack = c

	if c, err = stringToColor(cj.BrightRed); err != nil {
		return cols, err
	}
	cols.BrightRed = c

	if c, err = stringToColor(cj.BrightGreen); err != nil {
		return cols, err
	}
	cols.BrightGreen = c

	if c, err = stringToColor(cj.BrightYellow); err != nil {
		return cols, err
	}
	cols.BrightYellow = c

	if c, err = stringToColor(cj.BrightBlue); err != nil {
		return cols, err
	}
	cols.BrightBlue = c

	if c, err = stringToColor(cj.BrightMagenta); err != nil {
		return cols, err
	}
	cols.BrightMagenta = c

	if c, err = stringToColor(cj.BrightCyan); err != nil {
		return cols, err
	}
	cols.BrightCyan = c

	if c, err = stringToColor(cj.BrightWhite); err != nil {
		return cols, err
	}
	cols.BrightWhite = c

	return cols, nil
}
