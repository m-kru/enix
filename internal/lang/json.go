package lang

import (
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/regex"
)

type RegionJSON struct {
	Name  string
	Style string

	Start regex.RegexJSON
	End   regex.RegexJSON

	Attribute regex.RegexJSON
	Bold      regex.RegexJSON
	Comment   regex.RegexJSON
	Heading   regex.RegexJSON
	Italic    regex.RegexJSON
	Keyword   regex.RegexJSON
	Meta      regex.RegexJSON
	Mono      regex.RegexJSON
	Number    regex.RegexJSON
	Operator  regex.RegexJSON
	String    regex.RegexJSON
	Type      regex.RegexJSON
	Value     regex.RegexJSON
	Variable  regex.RegexJSON
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

	bold, err := rj.Bold.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Bold: %v", err)
	}

	comment, err := rj.Comment.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Comment: %v", err)
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

	meta, err := rj.Meta.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Meta: %v", err)
	}

	mono, err := rj.Mono.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Mono: %v", err)
	}

	number, err := rj.Number.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Number: %v", err)
	}

	operator, err := rj.Operator.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Operator: %v", err)
	}

	str, err := rj.String.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile String: %v", err)
	}

	typ, err := rj.Type.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Type: %v", err)
	}

	val, err := rj.Value.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Value: %v", err)
	}

	variable, err := rj.Variable.ToRegex()
	if err != nil {
		return nil, fmt.Errorf("can't compile Variable: %v", err)
	}

	return &Region{
		Name:       rj.Name,
		Style:      rj.Style,
		Start:      start,
		End:        end,
		CursorWord: nil,
		Attribute:  attr,
		Bold:       bold,
		Comment:    comment,
		Heading:    heading,
		Italic:     italic,
		Keyword:    keyword,
		Meta:       meta,
		Mono:       mono,
		Number:     number,
		Operator:   operator,
		String:     str,
		Type:       typ,
		Value:      val,
		Variable:   variable,
	}, nil
}

func readFiletypeDefFromJSON(lang string) ([]RegionJSON, error) {
	data, path, err := cfg.ReadConfigFile(filepath.Join("filetype", lang+".json"))
	if err != nil {
		return nil, fmt.Errorf("reading filetype from %s: %v", path, err)
	}

	if path == "" {
		return nil, nil
	}

	var langDef []RegionJSON
	err = json.Unmarshal(data, &langDef)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling filetype json: %v", err)
	}

	return langDef, nil
}

func langDefIntoHighlighter(regionsJSON []RegionJSON) (*Highlighter, error) {
	hl := &Highlighter{
		Regions:          []*Region{},
		matchingBrackets: nil,
		lineNum:          0,
		firstVisLineNum:  0,
		lastVisLineNum:   0,
		region:           nil,
		startTokens:      []RegionToken{},
	}

	if len(regionsJSON) == 0 {
		return hl, fmt.Errorf("no regions defined")
	}

	for _, rj := range regionsJSON {
		r, err := rj.ToRegion()
		if err != nil {
			return hl, fmt.Errorf("region '%s' :%v", rj.Name, err)
		}

		hl.Regions = append(hl.Regions, r)
	}

	// Some sanity checks
	if hl.Regions[0].Name != "Default" {
		return hl, fmt.Errorf(
			"name of the first region must be 'Default', curerent name '%s'",
			hl.Regions[0].Name,
		)
	}

	for i := 1; i < len(hl.Regions); i++ {
		r := hl.Regions[i]
		if r.Start.Regex == nil {
			panic(fmt.Sprintf("missing start regex for region '%s'", r.Name))
			//return hl, fmt.Errorf("missing start regex for region '%s'", r.Name)
		}
		if r.End.Regex == nil {
			panic(fmt.Sprintf("missing end regex for region '%s'", r.Name))
			//return hl, fmt.Errorf("missing end regex for region '%s'", r.Name)
		}
	}

	return hl, nil
}
