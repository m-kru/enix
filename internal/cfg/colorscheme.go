package cfg

import "github.com/gdamore/tcell/v2"

// Bg stands for background, Fg stands for foreground.
type Colorscheme struct {
	Default tcell.Style

	// Error displaying
	Error tcell.Style

	// Line Number
	LineNum tcell.Style

	Cursor     tcell.Style
	CursorWord tcell.Style // Color of the word under cursor

	Selection tcell.Style

	// Syntax highlighting
	Comment tcell.Style
	Keyword tcell.Style
	Type    tcell.Style
	Value   tcell.Style
}

// ColorschemeDefault return the default color scheme.
//
// The default color scheme doesn't require any color scheme files to be installed
// as it is embedded into the program's binary.
//
// The default color scheme uses the same colors as the terminal.
func ColorschemeDefault() Colorscheme {
	return Colorscheme{
		Default: tcell.StyleDefault,

		Error: tcell.StyleDefault.Foreground(tcell.ColorMaroon),

		LineNum: tcell.StyleDefault.Foreground(tcell.ColorSilver),

		Cursor:     tcell.StyleDefault.Reverse(true),
		CursorWord: tcell.StyleDefault.Foreground(tcell.ColorWhite),
	}
}
