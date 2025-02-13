package cfg

import (
	"fmt"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

var ColorsJSON colorsJSON
var Colors colors

type colors struct {
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

func DefaultColors() colors {
	return colors{
		Background:    tcell.ColorDefault,
		Foreground:    tcell.ColorDefault,
		Black:         tcell.ColorBlack,
		Red:           tcell.ColorMaroon,
		Green:         tcell.ColorGreen,
		Yellow:        tcell.ColorOlive,
		Blue:          tcell.ColorNavy,
		Magenta:       tcell.ColorPurple,
		Cyan:          tcell.ColorTeal,
		White:         tcell.ColorSilver,
		BrightBlack:   tcell.ColorGray,
		BrightRed:     tcell.ColorRed,
		BrightGreen:   tcell.ColorLime,
		BrightYellow:  tcell.ColorYellow,
		BrightBlue:    tcell.ColorBlue,
		BrightMagenta: tcell.ColorFuchsia,
		BrightCyan:    tcell.ColorAqua,
		BrightWhite:   tcell.ColorWhite,
	}
}

type colorsJSON struct {
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

func (cj colorsJSON) ToColors() (colors, error) {
	var cols colors
	var err error
	var c tcell.Color

	if cj.Background == "" {
		return cols, fmt.Errorf("missing value for Background")
	}
	if c, err = stringToColor(cj.Background); err != nil {
		return cols, err
	}
	cols.Background = c

	if cj.Foreground == "" {
		return cols, fmt.Errorf("missing value for Foreground")
	}
	if c, err = stringToColor(cj.Foreground); err != nil {
		return cols, err
	}
	cols.Foreground = c

	if cj.Black == "" {
		return cols, fmt.Errorf("missing value for Black")
	}
	if c, err = stringToColor(cj.Black); err != nil {
		return cols, err
	}
	cols.Black = c

	if cj.Red == "" {
		return cols, fmt.Errorf("missing value for Red")
	}
	if c, err = stringToColor(cj.Red); err != nil {
		return cols, err
	}
	cols.Red = c

	if cj.Green == "" {
		return cols, fmt.Errorf("missing value for Green")
	}
	if c, err = stringToColor(cj.Green); err != nil {
		return cols, err
	}
	cols.Green = c

	if cj.Yellow == "" {
		return cols, fmt.Errorf("missing value for Yellow")
	}
	if c, err = stringToColor(cj.Yellow); err != nil {
		return cols, err
	}
	cols.Yellow = c

	if cj.Blue == "" {
		return cols, fmt.Errorf("missing value for Blue")
	}
	if c, err = stringToColor(cj.Blue); err != nil {
		return cols, err
	}
	cols.Blue = c

	if cj.Magenta == "" {
		return cols, fmt.Errorf("missing value for Magenta")
	}
	if c, err = stringToColor(cj.Magenta); err != nil {
		return cols, err
	}
	cols.Magenta = c

	if cj.Cyan == "" {
		return cols, fmt.Errorf("missing value for Cyan")
	}
	if c, err = stringToColor(cj.Cyan); err != nil {
		return cols, err
	}
	cols.Cyan = c

	if cj.White == "" {
		return cols, fmt.Errorf("missing value for White")
	}
	if c, err = stringToColor(cj.White); err != nil {
		return cols, err
	}
	cols.White = c

	if cj.BrightBlack == "" {
		return cols, fmt.Errorf("missing value for BrightBlack")
	}
	if c, err = stringToColor(cj.BrightBlack); err != nil {
		return cols, err
	}
	cols.BrightBlack = c

	if cj.BrightRed == "" {
		return cols, fmt.Errorf("missing value for BrightRed")
	}
	if c, err = stringToColor(cj.BrightRed); err != nil {
		return cols, err
	}
	cols.BrightRed = c

	if cj.BrightGreen == "" {
		return cols, fmt.Errorf("missing value for BrightGreen")
	}
	if c, err = stringToColor(cj.BrightGreen); err != nil {
		return cols, err
	}
	cols.BrightGreen = c

	if cj.BrightYellow == "" {
		return cols, fmt.Errorf("missing value for BrightYellow")
	}
	if c, err = stringToColor(cj.BrightYellow); err != nil {
		return cols, err
	}
	cols.BrightYellow = c

	if cj.BrightBlue == "" {
		return cols, fmt.Errorf("missing value for BrightBlue")
	}
	if c, err = stringToColor(cj.BrightBlue); err != nil {
		return cols, err
	}
	cols.BrightBlue = c

	if cj.BrightMagenta == "" {
		return cols, fmt.Errorf("missing value for BrightMagenta")
	}
	if c, err = stringToColor(cj.BrightMagenta); err != nil {
		return cols, err
	}
	cols.BrightMagenta = c

	if cj.BrightCyan == "" {
		return cols, fmt.Errorf("missing value for BrightCyan")
	}
	if c, err = stringToColor(cj.BrightCyan); err != nil {
		return cols, err
	}
	cols.BrightCyan = c

	if cj.BrightWhite == "" {
		return cols, fmt.Errorf("missing value for BrightWhite")
	}
	if c, err = stringToColor(cj.BrightWhite); err != nil {
		return cols, err
	}
	cols.BrightWhite = c

	return cols, nil
}
