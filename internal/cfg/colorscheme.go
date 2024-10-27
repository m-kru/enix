package cfg

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"

	"github.com/m-kru/enix/internal/arg"

	"github.com/gdamore/tcell/v2"
)

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
	RepCount   tcell.Style
	StateMark  tcell.Style

	Prompt       tcell.Style
	PromptShadow tcell.Style

	// Syntax highlighting
	Attribute       tcell.Style
	Bold            tcell.Style
	Builtin         tcell.Style
	Code            tcell.Style
	Comment         tcell.Style
	Documentation   tcell.Style
	FormatSpecifier tcell.Style
	Function        tcell.Style
	Heading         tcell.Style
	Italic          tcell.Style
	Keyword         tcell.Style
	Link            tcell.Style
	Meta            tcell.Style
	Mono            tcell.Style
	Number          tcell.Style
	Operator        tcell.Style
	String          tcell.Style
	ToDo            tcell.Style
	Title           tcell.Style
	Type            tcell.Style
	Value           tcell.Style
	Variable        tcell.Style
	Other           tcell.Style
}

func (cs *Colorscheme) Style(name string) tcell.Style {
	switch name {
	case "CursorWord":
		return cs.CursorWord
	case "Attribute":
		return cs.Attribute
	case "Bold":
		return cs.Bold
	case "Builtin":
		return cs.Builtin
	case "Code":
		return cs.Code
	case "Comment":
		return cs.Comment
	case "Documentation":
		return cs.Documentation
	case "FormatSpecifier":
		return cs.FormatSpecifier
	case "Function":
		return cs.Function
	case "Heading":
		return cs.Heading
	case "Italic":
		return cs.Italic
	case "Keyword":
		return cs.Keyword
	case "Link":
		return cs.Link
	case "Meta":
		return cs.Meta
	case "Mono":
		return cs.Mono
	case "Number":
		return cs.Number
	case "Operator":
		return cs.Operator
	case "String":
		return cs.String
	case "ToDo":
		return cs.ToDo
	case "Title":
		return cs.Title
	case "Type":
		return cs.Type
	case "Value":
		return cs.Value
	case "Variable":
		return cs.Variable
	case "Other":
		return cs.Other
	default:
		return cs.Default
	}
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

		TabBar:     tcell.StyleDefault.Background(tcell.ColorBlack),
		CurrentTab: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOlive),

		LineNum: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Whitespace: tcell.StyleDefault.Foreground(tcell.ColorBlack),

		Cursor:     tcell.StyleDefault.Reverse(true),
		CursorWord: tcell.StyleDefault.Foreground(tcell.ColorWhite),

		StatusLine: tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorGray),
		RepCount:   tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorTeal),
		StateMark:  tcell.StyleDefault.Background(tcell.ColorBlack).Foreground(tcell.ColorOlive),

		Prompt:       tcell.StyleDefault.Foreground(tcell.ColorWhite),
		PromptShadow: tcell.StyleDefault.Foreground(tcell.ColorGray),

		Attribute:       tcell.StyleDefault.Foreground(tcell.ColorTeal),
		Builtin:         tcell.StyleDefault.Bold(true),
		Comment:         tcell.StyleDefault.Foreground(tcell.ColorGray),
		FormatSpecifier: tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		Keyword:         tcell.StyleDefault.Foreground(tcell.ColorNavy),
		Meta:            tcell.StyleDefault.Foreground(tcell.ColorPurple),
		Number:          tcell.StyleDefault.Foreground(tcell.ColorMaroon),
		Operator:        tcell.StyleDefault.Foreground(tcell.ColorOlive),
		String:          tcell.StyleDefault.Foreground(tcell.ColorGreen),
		Type:            tcell.StyleDefault.Foreground(tcell.ColorOlive),
		Value:           tcell.StyleDefault.Foreground(tcell.ColorMaroon),
	}
}

// readFromJSON reads colorscheme from file named "name.json".
func colorschemeFromJSON(name string) (Colorscheme, error) {
	cs := Colorscheme{}

	colorsDir := filepath.Join(ConfigDir, "colors")
	if arg.ColorsDir != "" {
		colorsDir = arg.ColorsDir
	}

	path := filepath.Join(colorsDir, name+".json")

	file, err := os.Open(path)
	if err != nil {
		return cs, fmt.Errorf("opening colorscheme file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return cs, fmt.Errorf("reading colors file: %v", err)
	}

	var colorschemeMap map[string]any
	err = json.Unmarshal([]byte(data), &colorschemeMap)
	if err != nil {
		return cs, fmt.Errorf("unmarshalling json colorscheme file: %v", err)
	}

	cs, err = colorschemeFromMap(colorschemeMap)
	if err != nil {
		return cs, fmt.Errorf("%s: %v", path, err)
	}

	return cs, nil
}

// colorschemeFromMap creates Colorscheme from colorscheme map read from JSON file.
func colorschemeFromMap(csm map[string]any) (Colorscheme, error) {
	var err error
	cs := Colorscheme{}

	if cs.Default, err = readStyleFromMap("Default", csm, &tcell.StyleDefault); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.Error, err = readStyleFromMap("Error", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.LineNum, err = readStyleFromMap("LineNum", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.Whitespace, err = readStyleFromMap("Whitespace", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.Cursor, err = readStyleFromMap("Cursor", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}
	if cs.CursorWord, err = readStyleFromMap("CursorWord", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.StatusLine, err = readStyleFromMap("StatusLine", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.StateMark, err = readStyleFromMap("StateMark", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.Prompt, err = readStyleFromMap("Prompt", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	if cs.PromptShadow, err = readStyleFromMap("PromptShadow", csm, &cs.Default); err != nil {
		return cs, fmt.Errorf("%v", err)
	}

	return cs, nil
}

func readStyleFromMap(name string, csm map[string]any, dfltStyle *tcell.Style) (tcell.Style, error) {
	style := tcell.StyleDefault
	if dfltStyle != nil {
		style = *dfltStyle
	}

	styleDefAny, ok := csm[name]
	if !ok {
		return style, nil
	}
	styleDef, ok := styleDefAny.(map[string]any)
	if !ok {
		return style, fmt.Errorf("invalid type for style \"%s\" in json file", name)
	}

	colorsAny, ok := csm["Colors"]
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
