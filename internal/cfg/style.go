package cfg

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strconv"

	"github.com/gdamore/tcell/v2"
)

var Style style

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
	CursorWord tcell.Style // Color of the word under cursor

	Find      tcell.Style
	Selection tcell.Style

	StatusLine tcell.Style
	RepCount   tcell.Style
	StateMark  tcell.Style
	FindMark   tcell.Style

	Prompt       tcell.Style
	PromptShadow tcell.Style

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
		CurrentTab: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOlive),

		LineNum: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Whitespace: tcell.StyleDefault.Foreground(tcell.ColorBlack),

		Cursor:     tcell.StyleDefault.Reverse(true),
		CursorWord: tcell.StyleDefault.Foreground(tcell.ColorWhite),

		Find:      tcell.StyleDefault.Background(tcell.ColorOlive),
		Selection: tcell.StyleDefault.Background(tcell.ColorMaroon).Foreground(tcell.ColorWhite),

		StatusLine: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray),
		RepCount:   tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorTeal),
		StateMark:  tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGreen),
		FindMark:   tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOlive),

		Prompt:       tcell.StyleDefault.Foreground(tcell.ColorWhite),
		PromptShadow: tcell.StyleDefault.Foreground(tcell.ColorGray),

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

// styleFromJSON reads style from file named "style/<name>.json".
func styleFromJSON(name string) (style, error) {
	data, path, err := ReadConfigFile(filepath.Join("style", name+".json"))
	if err != nil {
		return DefaultStyle(), fmt.Errorf("reading style file: %v", err)
	}

	var styleMap map[string]any
	err = json.Unmarshal(data, &styleMap)
	if err != nil {
		return DefaultStyle(), fmt.Errorf("unmarshalling json style file: %v", err)
	}

	cs, err := styleFromMap(styleMap)
	if err != nil {
		return DefaultStyle(), fmt.Errorf("%s: %v", path, err)
	}

	return cs, nil
}

// styleFromMap creates Style from style map read from JSON file.
func styleFromMap(sm map[string]any) (style, error) {
	var err error
	s := DefaultStyle()

	if s.Default, err = readStyleFromMap("Default", sm, &tcell.StyleDefault); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.Error, err = readStyleFromMap("Error", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.LineNum, err = readStyleFromMap("LineNum", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.Whitespace, err = readStyleFromMap("Whitespace", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.Cursor, err = readStyleFromMap("Cursor", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}
	if s.CursorWord, err = readStyleFromMap("CursorWord", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.StatusLine, err = readStyleFromMap("StatusLine", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.StateMark, err = readStyleFromMap("StateMark", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.Prompt, err = readStyleFromMap("Prompt", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	if s.PromptShadow, err = readStyleFromMap("PromptShadow", sm, &s.Default); err != nil {
		return s, fmt.Errorf("%v", err)
	}

	return s, nil
}

func readStyleFromMap(name string, sm map[string]any, dfltStyle *tcell.Style) (tcell.Style, error) {
	style := tcell.StyleDefault
	if dfltStyle != nil {
		style = *dfltStyle
	}

	styleDefAny, ok := sm[name]
	if !ok {
		return style, nil
	}
	styleDef, ok := styleDefAny.(map[string]any)
	if !ok {
		return style, fmt.Errorf("invalid type for style \"%s\" in json file", name)
	}

	colorsAny, ok := sm["Colors"]
	if !ok {
		return style, fmt.Errorf("colorscheme file misses colors definitions")
	}
	colors, ok := colorsAny.(map[string]any)
	if !ok {
		return style, fmt.Errorf("invalid type for \"Colors\" definition")
	}

	if col, ok := styleDef["Fg"]; ok {
		colStr, ok := col.(string)
		if !ok {
			return style, fmt.Errorf("invalid type for \"Fg\" in json file, expected string")
		}
		val, err := getColor(colors, colStr)
		if err != nil {
			return style, fmt.Errorf("%v", err)
		}
		style = style.Foreground(tcell.NewHexColor(val))
	}

	if col, ok := styleDef["Bg"]; ok {
		colStr, ok := col.(string)
		if !ok {
			return style, fmt.Errorf("invalid type for \"Bg\" in json file, expected string")
		}
		val, err := getColor(colors, colStr)
		if err != nil {
			return style, fmt.Errorf("%v", err)
		}
		style = style.Background(tcell.NewHexColor(val))
	}

	return style, nil
}

func getColor(cm map[string]any, name string) (int32, error) {
	var val int64
	var err error

	if value, ok := cm[name]; ok {
		if hex, ok := value.(string); ok {
			val, err = strconv.ParseInt(hex, 16, 32)
			if err != nil {
				return 0, fmt.Errorf("can't convert value for %s color: %v", name, err)
			}
		} else {
			return 0, fmt.Errorf("invalid value type for %s color, expected string", name)
		}
	} else {
		return 0, fmt.Errorf("missing definition of %s color", name)
	}

	return int32(val), nil
}
