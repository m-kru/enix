package lang

import (
	"fmt"
)

type Highlighter struct {
	Regions []*Region
}

// DefaultHighlighter returns a highlighter highlighting only a cursor word.
func DefaultHighlighter() *Highlighter {
	return &Highlighter{Regions: []*Region{DefaultRegion()}}
}

func NewHighlighter(lang string) (*Highlighter, error) {
	if lang == "" {
		return DefaultHighlighter(), nil
	}

	langDef, err := readFiletypeDefFromJSON(lang)
	if err != nil {
		return DefaultHighlighter(),
			fmt.Errorf("creating highlighter for %s filetype: %v", lang, err)
	}
	if langDef == nil {
		return DefaultHighlighter(), nil
	}

	hl, err := langDefIntoHighlighter(langDef)
	if err != nil {
		return DefaultHighlighter(),
			fmt.Errorf("%s highlighter: %v", lang, err)
	}

	return &hl, nil
}

func langDefIntoHighlighter(regionsJSON []RegionJSON) (Highlighter, error) {
	hl := Highlighter{Regions: []*Region{}}

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
