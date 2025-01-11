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

	Start struct {
		Regex              string
		NegativeLookBehind string
		PositiveLookBehind string
		NegativeLookAhead  string
		PositiveLookAhead  string
	}
	End struct {
		Regex              string
		NegativeLookBehind string
		PositiveLookBehind string
		NegativeLookAhead  string
		PositiveLookAhead  string
	}

	Attribute       string
	Builtin         string
	Bold            string
	Code            string
	Comment         string
	Documentation   string
	EscapeSequence  string
	FormatSpecifier string
	Function        string
	Heading         string
	Italic          string
	Keyword         string
	Link            string
	Meta            string
	Mono            string
	Number          string
	Operator        string
	String          string
	ToDo            string
	Title           string
	Type            string
	Value           string
	Variable        string
}

func (rj RegionJSON) ToRegion() (*Region, error) {
	var err error

	var sre *regexp.Regexp
	if rj.Start.Regex != "" {
		sre, err = regexp.Compile(rj.Start.Regex)
		if err != nil {
			return nil, fmt.Errorf("can't compile start regex: %v", err)
		}
	}

	var snlb *regexp.Regexp
	if rj.Start.NegativeLookBehind != "" {
		snlb, err = regexp.Compile(rj.Start.NegativeLookBehind)
		if err != nil {
			return nil, fmt.Errorf("can't compile start regex negative lookbehind: %v", err)
		}
	}

	var splb *regexp.Regexp
	if rj.Start.PositiveLookBehind != "" {
		splb, err = regexp.Compile(rj.Start.PositiveLookBehind)
		if err != nil {
			return nil, fmt.Errorf("can't compile start regex positive lookbehind: %v", err)
		}
	}

	var snla *regexp.Regexp
	if rj.Start.NegativeLookAhead != "" {
		snla, err = regexp.Compile(rj.Start.NegativeLookAhead)
		if err != nil {
			return nil, fmt.Errorf("can't compile start regex negative lookahead: %v", err)
		}
	}

	var spla *regexp.Regexp
	if rj.Start.PositiveLookAhead != "" {
		spla, err = regexp.Compile(rj.Start.PositiveLookAhead)
		if err != nil {
			return nil, fmt.Errorf("can't compile start regex positive lookahead: %v", err)
		}
	}

	var ere *regexp.Regexp
	if rj.End.Regex != "" {
		ere, err = regexp.Compile(rj.End.Regex)
		if err != nil {
			return nil, fmt.Errorf("can't compile end regex: %v", err)
		}
	}

	var enlb *regexp.Regexp
	if rj.End.NegativeLookBehind != "" {
		enlb, err = regexp.Compile(rj.End.NegativeLookBehind)
		if err != nil {
			return nil, fmt.Errorf("can't compile edn regex negative lookbehind: %v", err)
		}
	}

	var eplb *regexp.Regexp
	if rj.End.PositiveLookBehind != "" {
		eplb, err = regexp.Compile(rj.End.PositiveLookBehind)
		if err != nil {
			return nil, fmt.Errorf("can't compile edn regex positive lookbehind: %v", err)
		}
	}

	var enla *regexp.Regexp
	if rj.End.NegativeLookAhead != "" {
		enla, err = regexp.Compile(rj.End.NegativeLookAhead)
		if err != nil {
			return nil, fmt.Errorf("can't compile edn regex negative lookahead: %v", err)
		}
	}

	var epla *regexp.Regexp
	if rj.End.PositiveLookAhead != "" {
		epla, err = regexp.Compile(rj.End.PositiveLookAhead)
		if err != nil {
			return nil, fmt.Errorf("can't compile edn regex positive lookahead: %v", err)
		}
	}

	var attr *regexp.Regexp
	if rj.Attribute != "" {
		attr, err = regexp.Compile(rj.Attribute)
		if err != nil {
			return nil, fmt.Errorf("can't compile attribute: %v", err)
		}
	}

	var builtin *regexp.Regexp
	if rj.Builtin != "" {
		builtin, err = regexp.Compile(rj.Builtin)
		if err != nil {
			return nil, fmt.Errorf("can't compile builtin: %v", err)
		}
	}

	var bold *regexp.Regexp
	if rj.Bold != "" {
		bold, err = regexp.Compile(rj.Bold)
		if err != nil {
			return nil, fmt.Errorf("can't compile bold: %v", err)
		}
	}

	var code *regexp.Regexp
	if rj.Code != "" {
		code, err = regexp.Compile(rj.Code)
		if err != nil {
			return nil, fmt.Errorf("can't compile code: %v", err)
		}
	}

	var comment *regexp.Regexp
	if rj.Comment != "" {
		comment, err = regexp.Compile(rj.Comment)
		if err != nil {
			return nil, fmt.Errorf("can't compile comment: %v", err)
		}
	}

	var doc *regexp.Regexp
	if rj.Documentation != "" {
		doc, err = regexp.Compile(rj.Documentation)
		if err != nil {
			return nil, fmt.Errorf("can't compile documentation: %v", err)
		}
	}

	var escSeq *regexp.Regexp
	if rj.EscapeSequence != "" {
		escSeq, err = regexp.Compile(rj.EscapeSequence)
		if err != nil {
			return nil, fmt.Errorf("can't compile escape sequence: %v", err)
		}
	}

	var fmtSpec *regexp.Regexp
	if rj.FormatSpecifier != "" {
		fmtSpec, err = regexp.Compile(rj.FormatSpecifier)
		if err != nil {
			return nil, fmt.Errorf("can't compile format specifier: %v", err)
		}
	}

	var fun *regexp.Regexp
	if rj.Function != "" {
		fun, err = regexp.Compile(rj.Function)
		if err != nil {
			return nil, fmt.Errorf("can't compile format function: %v", err)
		}
	}

	var heading *regexp.Regexp
	if rj.Heading != "" {
		heading, err = regexp.Compile(rj.Heading)
		if err != nil {
			return nil, fmt.Errorf("can't compile heading: %v", err)
		}
	}

	var italic *regexp.Regexp
	if rj.Italic != "" {
		italic, err = regexp.Compile(rj.Italic)
		if err != nil {
			return nil, fmt.Errorf("can't compile italic: %v", err)
		}
	}

	var keyword *regexp.Regexp
	if rj.Keyword != "" {
		keyword, err = regexp.Compile(rj.Keyword)
		if err != nil {
			return nil, fmt.Errorf("can't compile keyword: %v", err)
		}
	}

	var link *regexp.Regexp
	if rj.Link != "" {
		link, err = regexp.Compile(rj.Link)
		if err != nil {
			return nil, fmt.Errorf("can't compile link: %v", err)
		}
	}

	var meta *regexp.Regexp
	if rj.Meta != "" {
		meta, err = regexp.Compile(rj.Meta)
		if err != nil {
			return nil, fmt.Errorf("can't compile meta: %v", err)
		}
	}

	var mono *regexp.Regexp
	if rj.Mono != "" {
		mono, err = regexp.Compile(rj.Mono)
		if err != nil {
			return nil, fmt.Errorf("can't compile mono: %v", err)
		}
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

	var title *regexp.Regexp
	if rj.Title != "" {
		title, err = regexp.Compile(rj.Title)
		if err != nil {
			return nil, fmt.Errorf("can't compile title: %v", err)
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
		Name:  rj.Name,
		Style: rj.Style,
		Start: regex.Regex{
			Regex:              sre,
			NegativeLookBehind: snlb,
			PositiveLookBehind: splb,
			NegativeLookAhead:  snla,
			PositiveLookAhead:  spla,
		},
		End: regex.Regex{
			Regex:              ere,
			NegativeLookBehind: enlb,
			PositiveLookBehind: eplb,
			NegativeLookAhead:  enla,
			PositiveLookAhead:  epla,
		},
		CursorWord:      nil,
		Attribute:       attr,
		Builtin:         builtin,
		Bold:            bold,
		Code:            code,
		Comment:         comment,
		Documentation:   doc,
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
		Title:           title,
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
