package lang

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"regexp"

	"github.com/m-kru/enix/internal/arg"
	"github.com/m-kru/enix/internal/cfg"
)

type Highlighter struct {
	Regions []*Region
}

// DefaultHighlighter returns a highlighter highlighting only a cursor word.
func DefaultHighlighter() Highlighter {
	return Highlighter{Regions: []*Region{&Region{Name: "Default"}}}
}

func NewHighlighter(lang string) (Highlighter, error) {
	hl := Highlighter{}

	if lang == "" {
		return DefaultHighlighter(), nil
	}

	langDef, err := readLangDefFromJSON(lang)
	if err != nil {
		return hl, fmt.Errorf("creating highlighter for %s language: %v", lang, err)
	}

	hl, err = langDefIntoHighlighter(langDef)
	if err != nil {
		return hl, fmt.Errorf("%s highlighter: %v", lang, err)
	}

	return hl, nil
}

func readLangDefFromJSON(lang string) ([]any, error) {
	langsDir := filepath.Join(cfg.ConfigDir, "langs")
	if arg.LangsDir != "" {
		langsDir = arg.LangsDir
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

	var langDef []any
	err = json.Unmarshal([]byte(data), &langDef)
	if err != nil {
		return nil, fmt.Errorf("unmarshalling json language file: %v", err)
	}

	return langDef, nil
}

func fieldToRegexp(name string, regDef map[string]any, reg *Region) error {
	fieldAny, ok := regDef[name]
	if !ok {
		return nil
	}

	fieldStr, ok := fieldAny.(string)
	if !ok {
		return fmt.Errorf("invalid type for %q, expected string", name)
	}

	re, err := regexp.Compile(fieldStr)
	if err != nil {
		return fmt.Errorf("field %q: %v", name, err)
	}

	switch name {
	case "Attribute":
		reg.Attribute = re
	case "Bold":
		reg.Bold = re
	case "Builtin":
		reg.Builtin = re
	case "Code":
		reg.Code = re
	case "Comment":
		reg.Comment = re
	case "Documentation":
		reg.Documentation = re
	case "FormatSpecifier":
		reg.FormatSpecifier = re
	case "Function":
		reg.Function = re
	case "Heading":
		reg.Heading = re
	case "Italic":
		reg.Italic = re
	case "Keyword":
		reg.Keyword = re
	case "Link":
		reg.Link = re
	case "Meta":
		reg.Meta = re
	case "Mono":
		reg.Mono = re
	case "Number":
		reg.Number = re
	case "Operator":
		reg.Operator = re
	case "String":
		reg.String = re
	case "ToDo":
		reg.ToDo = re
	case "Title":
		reg.Title = re
	case "Type":
		reg.Type = re
	case "Value":
		reg.Value = re
	case "Variable":
		reg.Variable = re
	case "Other":
		reg.Other = re
	}

	return nil
}

func langDefIntoHighlighter(langDef []any) (Highlighter, error) {
	hl := Highlighter{}
	if len(langDef) == 0 {
		return hl, fmt.Errorf("no regions defined")
	}

	for i, regDef := range langDef {
		reg := Region{}

		regDef, ok := regDef.(map[string]any)
		if !ok {
			return hl, fmt.Errorf("invalid type for %d region", i)
		}

		// Name
		nameAny, ok := regDef["Name"]
		if !ok {
			return hl, fmt.Errorf("missing region name for region %d", i)
		}
		name, ok := nameAny.(string)
		if !ok {
			return hl, fmt.Errorf("invalid type for %d region \"Name\", expected string", i)
		}
		if i == 0 && name != "Default" {
			return hl, fmt.Errorf("name of the first region must be \"Default\"")
		}
		reg.Name = name

		// Style
		style := "Default"
		styleAny, ok := regDef["Style"]
		if i == 0 && ok {
			return hl, fmt.Errorf("default region doesn't accept \"Style\"")
		} else if i != 0 {
			style, ok = styleAny.(string)
			if !ok {
				return hl, fmt.Errorf(
					"invalid type for %q region \"Style\", expected string", name,
				)
			}
		}
		reg.Style = style

		// StartRegex
		startRegexAny, ok := regDef["StartRegex"]
		if i == 0 && ok {
			return hl, fmt.Errorf("default region doesn't accept \"StartRegex\"")
		} else if i != 0 && !ok {
			return hl, fmt.Errorf("missing \"StartRegex\" for region %q", name)
		} else if i != 0 {
			startRegex, ok := startRegexAny.(string)
			if !ok {
				return hl, fmt.Errorf(
					"invalid type for %q region \"StartRegex\", expected string", name,
				)
			}
			re, err := regexp.Compile(startRegex)
			if err != nil {
				return hl, fmt.Errorf("region %q, \"StartRegex\": %v", name, err)
			}
			reg.StartRegexp = re
		}

		// EndRegex
		endRegexAny, ok := regDef["EndRegex"]
		if i == 0 && ok {
			return hl, fmt.Errorf("default region doesn't accept \"EndRegex\"")
		} else if i != 0 && !ok {
			return hl, fmt.Errorf("missing \"EndRegex\" for region %q", name)
		} else if i != 0 {
			endRegex, ok := endRegexAny.(string)
			if !ok {
				return hl, fmt.Errorf(
					"invalid type for %q region \"EndRegex\", expected string", name,
				)
			}
			re, err := regexp.Compile(endRegex)
			if err != nil {
				return hl, fmt.Errorf("region %q, \"EndRegex\": %v", name, err)
			}
			reg.EndRegexp = re
		}

		if err := fieldToRegexp("Attribute", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Bold", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Builtin", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Code", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Comment", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Documentation", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("FormatSpecifier", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Function", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Heading", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Italic", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Keyword", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Link", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Meta", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Mono", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Number", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Operator", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("String", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("ToDo", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Title", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Type", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Value", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Variable", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}
		if err := fieldToRegexp("Other", regDef, &reg); err != nil {
			return hl, fmt.Errorf("region %q: %v", name, err)
		}

		hl.Regions = append(hl.Regions, &reg)
	}

	return hl, nil
}
