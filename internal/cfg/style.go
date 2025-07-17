package cfg

import (
	"github.com/gdamore/tcell/v2"
)

var Style style
var StyleJSON styleJSON

type itemStyle struct {
	Fg     string // Foreground
	Bg     string // Background
	Bold   bool
	Italic bool
}

func (is itemStyle) ToTcellStyle(dflt tcell.Style) (tcell.Style, error) {
	s := dflt

	if is.Fg != "" {
		color, err := Colors.Get(is.Fg)
		if err != nil {
			return s, err
		}
		s = s.Foreground(color)
	}

	if is.Bg != "" {
		color, err := Colors.Get(is.Bg)
		if err != nil {
			return s, err
		}
		s = s.Background(color)
	}

	s = s.Bold(is.Bold)
	s = s.Italic(is.Italic)

	return s, nil
}

type styleJSON struct {
	Default         itemStyle
	Error           itemStyle
	Warning         itemStyle
	TabBar          itemStyle
	CurrentTab      itemStyle
	LineNum         itemStyle
	Whitespace      itemStyle
	Cursor          itemStyle
	CursorWord      itemStyle
	Find            itemStyle
	Selection       itemStyle
	MatchingBracket itemStyle
	StatusLine      itemStyle
	RepCount        itemStyle
	StateMark       itemStyle
	FindMark        itemStyle
	Prompt          itemStyle
	PromptShadow    itemStyle
	Menu            itemStyle
	MenuItem        itemStyle
	Attribute       itemStyle
	Bold            itemStyle
	Comment         itemStyle
	Heading         itemStyle
	Italic          itemStyle
	Keyword         itemStyle
	Meta            itemStyle
	Mono            itemStyle
	Number          itemStyle
	Operator        itemStyle
	String          itemStyle
	Type            itemStyle
	Value           itemStyle
	Variable        itemStyle
}

func (sj styleJSON) ToStyle() (style, error) {
	var s style
	var ts tcell.Style
	var err error

	if ts, err = sj.Default.ToTcellStyle(tcell.StyleDefault); err != nil {
		return s, err
	}
	s.Default = ts

	if ts, err = sj.Error.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Error = ts

	if ts, err = sj.Warning.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Warning = ts

	if ts, err = sj.TabBar.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.TabBar = ts

	if ts, err = sj.CurrentTab.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.CurrentTab = ts

	if ts, err = sj.LineNum.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.LineNum = ts

	if ts, err = sj.Whitespace.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Whitespace = ts

	if ts, err = sj.Cursor.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Cursor = ts

	if ts, err = sj.CursorWord.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.CursorWord = ts

	if ts, err = sj.Find.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Find = ts

	if ts, err = sj.Selection.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Selection = ts

	if ts, err = sj.MatchingBracket.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.MatchingBracket = ts

	if ts, err = sj.StatusLine.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.StatusLine = ts

	if ts, err = sj.RepCount.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.RepCount = ts

	if ts, err = sj.StateMark.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.StateMark = ts

	if ts, err = sj.FindMark.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.FindMark = ts

	if ts, err = sj.Prompt.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Prompt = ts

	if ts, err = sj.PromptShadow.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.PromptShadow = ts

	if ts, err = sj.Menu.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Menu = ts

	if ts, err = sj.MenuItem.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.MenuItem = ts

	// Syntax highlighting

	if ts, err = sj.Attribute.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Attribute = ts

	if ts, err = sj.Bold.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Bold = ts

	if ts, err = sj.Comment.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Comment = ts

	if ts, err = sj.Heading.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Heading = ts

	if ts, err = sj.Italic.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Italic = ts

	if ts, err = sj.Keyword.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Keyword = ts

	if ts, err = sj.Meta.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Meta = ts

	if ts, err = sj.Mono.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Mono = ts

	if ts, err = sj.Number.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Number = ts

	if ts, err = sj.Operator.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Operator = ts

	if ts, err = sj.String.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.String = ts

	if ts, err = sj.Type.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Type = ts

	if ts, err = sj.Value.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Value = ts

	if ts, err = sj.Variable.ToTcellStyle(s.Default); err != nil {
		return s, err
	}
	s.Variable = ts

	return s, nil
}

type style struct {
	Default tcell.Style

	// Error displaying
	Error   tcell.Style
	Warning tcell.Style

	TabBar     tcell.Style
	CurrentTab tcell.Style

	// Line Number
	LineNum tcell.Style

	Whitespace tcell.Style

	Cursor     tcell.Style
	CursorWord tcell.Style // Style of the word under cursor

	Find            tcell.Style
	Selection       tcell.Style
	MatchingBracket tcell.Style

	StatusLine tcell.Style
	RepCount   tcell.Style
	StateMark  tcell.Style
	FindMark   tcell.Style

	Prompt       tcell.Style
	PromptShadow tcell.Style

	Menu     tcell.Style
	MenuItem tcell.Style // Menu selected item

	// Syntax highlighting
	Attribute tcell.Style
	Bold      tcell.Style
	Comment   tcell.Style
	Heading   tcell.Style
	Italic    tcell.Style
	Keyword   tcell.Style
	Meta      tcell.Style
	Mono      tcell.Style
	Number    tcell.Style
	Operator  tcell.Style
	String    tcell.Style
	Type      tcell.Style
	Value     tcell.Style
	Variable  tcell.Style
}

func (s *style) Get(name string) tcell.Style {
	switch name {
	case "CursorWord":
		return s.CursorWord
	case "Attribute":
		return s.Attribute
	case "Bold":
		return s.Bold
	case "Comment":
		return s.Comment
	case "Error":
		return s.Error
	case "Heading":
		return s.Heading
	case "Italic":
		return s.Italic
	case "Keyword":
		return s.Keyword
	case "Meta":
		return s.Meta
	case "Mono":
		return s.Mono
	case "Number":
		return s.Number
	case "Operator":
		return s.Operator
	case "String":
		return s.String
	case "Type":
		return s.Type
	case "Value":
		return s.Value
	case "Variable":
		return s.Variable
	default:
		return s.Default
	}
}

// DefaultStyle return the default style.
//
// The default color scheme doesn't require any color scheme files to be installed
// as it is embedded into the program's binary.
//
// The default color scheme uses the same colors as the terminal.
func DefaultStyle() style {
	return style{
		Default: tcell.StyleDefault,

		Error:   tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		Warning: tcell.StyleDefault.Foreground(tcell.ColorOlive),

		TabBar:     tcell.StyleDefault.Background(tcell.ColorBlack),
		CurrentTab: tcell.StyleDefault.Background(tcell.ColorBlack).Bold(true),

		LineNum: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Whitespace: tcell.StyleDefault.Foreground(tcell.ColorBlack),

		Cursor:     tcell.StyleDefault.Reverse(true),
		CursorWord: tcell.StyleDefault.Foreground(tcell.ColorWhite),

		Find:            tcell.StyleDefault.Background(tcell.ColorOlive),
		Selection:       tcell.StyleDefault.Background(tcell.ColorMaroon).Foreground(tcell.ColorWhite),
		MatchingBracket: tcell.StyleDefault.Background(tcell.ColorGray).Bold(true),

		StatusLine: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray),
		RepCount:   tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorTeal),
		StateMark:  tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen),
		FindMark:   tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOlive),

		Prompt:       tcell.StyleDefault.Foreground(tcell.ColorWhite),
		PromptShadow: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Menu:     tcell.StyleDefault.Background(tcell.ColorBlack),
		MenuItem: tcell.StyleDefault.Background(tcell.ColorGray).Bold(true),

		Attribute: tcell.StyleDefault.Foreground(tcell.ColorTeal),
		Bold:      tcell.StyleDefault.Bold(true),
		Comment:   tcell.StyleDefault.Foreground(tcell.ColorGray),
		Heading:   tcell.StyleDefault.Foreground(tcell.ColorTeal),
		Italic:    tcell.StyleDefault.Italic(true),
		Keyword:   tcell.StyleDefault.Foreground(tcell.ColorNavy),
		Meta:      tcell.StyleDefault.Foreground(tcell.ColorPurple),
		Mono:      tcell.StyleDefault.Foreground(tcell.ColorPurple),
		Number:    tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		Operator:  tcell.StyleDefault.Foreground(tcell.ColorOlive),
		String:    tcell.StyleDefault.Foreground(tcell.ColorGreen),
		Type:      tcell.StyleDefault.Foreground(tcell.ColorOlive),
		Value:     tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		Variable:  tcell.StyleDefault.Foreground(tcell.ColorMaroon),
	}
}
