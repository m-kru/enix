package lang

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path"
	"path/filepath"
	"regexp"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/regex"
)

type RegionJSON struct {
	Name  string
	Style string

	Start regex.RegexJSON
	End   regex.RegexJSON

	Attribute       regex.RegexJSON
	Builtin         regex.RegexJSON
	Bold            regex.RegexJSON
	Comment         regex.RegexJSON
	EscapeSequence  regex.RegexJSON
	FormatSpecifier regex.RegexJSON
	Function        regex.RegexJSON
	Heading         regex.RegexJSON
	Italic          regex.RegexJSON
	Keyword         regex.RegexJSON
	Link            regex.RegexJSON
	Meta            regex.RegexJSON
	Mono            regex.RegexJSON
	Number          string
	Operator        string
	String          string
	ToDo            string
	Type            string
	Value           string
	Variable        string
}

func (rj RegionJSON) ToRegion() (*Region, error) {
	var err error

	start, err := rj.Start.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Start: %v", err)
	}

	end, err := rj.End.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile End: %v", err)
	}

	attr, err := rj.Attribute.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Attribute: %v", err)
	}

	builtin, err := rj.Builtin.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Builtin: %v", err)
	}

	bold, err := rj.Bold.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Bold: %v", err)
	}

	comment, err := rj.Comment.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Comment: %v", err)
	}

	escSeq, err := rj.EscapeSequence.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile EscapeSequence: %v", err)
	}

	fmtSpec, err := rj.FormatSpecifier.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile FormatSpecifier: %v", err)
	}

	fun, err := rj.Function.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Function: %v", err)
	}

	heading, err := rj.Heading.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Heading: %v", err)
	}

	italic, err := rj.Italic.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Italic: %v", err)
	}

	keyword, err := rj.Keyword.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Keyword: %v", err)
	}

	link, err := rj.Link.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Link: %v", err)
	}

	meta, err := rj.Meta.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Meta: %v", err)
	}

	mono, err := rj.Mono.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Mono: %v", err)
	}

	var number *regexp.Regexp
	if rj.Number != "" {
		number, err = regexp.Compile(rj.Number)
		if err != nil {
			return nil, fmt.Errorf("can't compile number: %v", err)
		}
	}

	var operator *regexp.Regexp
	if rj.Operator != "" {
		operator, err = regexp.Compile(rj.Operator)
		if err != nil {
			return nil, fmt.Errorf("can't compile operator: %v", err)
		}
	}

	var str *regexp.Regexp
	if rj.String != "" {
		str, err = regexp.Compile(rj.String)
		if err != nil {
			return nil, fmt.Errorf("can't compile string: %v", err)
		}
	}

	var todo *regexp.Regexp
	if rj.ToDo != "" {
		todo, err = regexp.Compile(rj.ToDo)
		if err != nil {
			return nil, fmt.Errorf("can't compile todo: %v", err)
		}
	}

	var typ *regexp.Regexp
	if rj.Type != "" {
		typ, err = regexp.Compile(rj.Type)
		if err != nil {
			return nil, fmt.Errorf("can't compile type: %v", err)
		}
	}

	var val *regexp.Regexp
	if rj.Value != "" {
		val, err = regexp.Compile(rj.Value)
		if err != nil {
			return nil, fmt.Errorf("can't compile value: %v", err)
		}
	}

	var variable *regexp.Regexp
	if rj.Variable != "" {
		variable, err = regexp.Compile(rj.Variable)
		if err != nil {
			return nil, fmt.Errorf("can't compile variable: %v", err)
		}
	}

	return &Region{
		Name:            rj.Name,
		Style:           rj.Style,
		Start:           start,
		End:             end,
		CursorWord:      nil,
		Attribute:       attr,
		Builtin:         builtin,
		Bold:            bold,
		Comment:         comment,
		EscapeSequence:  escSeq,
		FormatSpecifier: fmtSpec,
		Function:        fun,
		Heading:         heading,
		Italic:          italic,
		Keyword:         keyword,
		Link:            link,
		Meta:            meta,
		Mono:            mono,
		Number:          number,
		Operator:        operator,
		String:          str,
		ToDo:            todo,
		Type:            typ,
		Value:           val,
		Variable:        variable,
	}, nil
}

func readLangDefFromJSON(lang string) ([]RegionJSON, error) {
	langsDir := ""
	if arg.LangsDir != "" {
		langsDir = arg.LangsDir
	}
	if langsDir == "" {
		langsDir = path.Join(os.Getenv("ENIX_RC_DIR"), "langs")
	}
	if langsDir == "" {
		langsDir = filepath.Join(cfg.ConfigDir, "langs")
	}

	path := filepath.Join(langsDir, lang+".json")

	file, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("opening language file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return nil, fmt.Errorf("reading language file: %v", err)
	}

	var langDef []RegionJSON
	err = json.Unmarshal(data, &langDef)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json language file: %v", err)
	}

	return langDef, nil
}
