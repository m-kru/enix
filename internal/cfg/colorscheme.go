package cfg

import "github.com/gdamore/tcell/v2"

// Bg stands for background, Fg stands for foreground.
type Colorscheme struct {
	Default tcell.Style

	// Error displaying
	Error tcell.Style

	TabBar     tcell.Style
	CurrentTab tcell.Style

	// Line Number
	LineNum tcell.Style

	Whitespace tcell.Style

	Cursor     tcell.Style
	CursorWord tcell.Style // Color of the word under cursor

	Selection tcell.Style

	StatusLine tcell.Style

	Prompt       tcell.Style
	PromptShadow tcell.Style

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

		TabBar:     tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray),
		CurrentTab: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorWhite),

		LineNum: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Whitespace: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Cursor:     tcell.StyleDefault.Reverse(true),
		CursorWord: tcell.StyleDefault.Foreground(tcell.ColorWhite),

		StatusLine: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray),

		Prompt:       tcell.StyleDefault.Foreground(tcell.ColorWhite),
		PromptShadow: tcell.StyleDefault.Foreground(tcell.ColorGray),
	}
}
