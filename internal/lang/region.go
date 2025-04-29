package lang

import (
	"regexp"
	"sort"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/regex"
	"github.com/m-kru/enix/internal/util"

	"github.com/gdamore/tcell/v2"
)

type match struct {
	start int // Rune start index
	end   int // Rune end index
	style tcell.Style
}

type Region struct {
	Name  string
	Style string // Region default style

	Start *regex.Regex
	End   *regex.Regex

	CursorWord *regexp.Regexp

	Attribute *regex.Regex
	Bold      *regex.Regex
	Comment   *regex.Regex
	Heading   *regex.Regex
	Italic    *regex.Regex
	Keyword   *regex.Regex
	Meta      *regex.Regex
	Mono      *regex.Regex
	Number    *regex.Regex
	Operator  *regex.Regex
	String    *regex.Regex
	Type      *regex.Regex
	Value     *regex.Regex
	Variable  *regex.Regex
}

func DefaultRegion() *Region {
	return &Region{
		Name:       "Default",
		Style:      "",
		Start:      nil,
		End:        nil,
		CursorWord: nil,
		Attribute:  nil,
		Bold:       nil,
		Comment:    nil,
		Heading:    nil,
		Italic:     nil,
		Keyword:    nil,
		Meta:       nil,
		Mono:       nil,
		Number:     nil,
		Operator:   nil,
		String:     nil,
		Type:       nil,
		Value:      nil,
		Variable:   nil,
	}
}

func (reg Region) match(buf []byte) []match {
	matches := make([]match, 0, 32)

	if reg.CursorWord != nil {
		finds := reg.CursorWord.FindAllIndex(buf, -1)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f[0]),
				end:   util.ByteIdxToRuneIdx(buf, f[1]),
				style: cfg.Style.CursorWord,
			}
			matches = append(matches, m)
		}
	}

	if reg.Attribute != nil {
		finds := reg.Attribute.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Attribute,
			}
			matches = append(matches, m)
		}
	}

	if reg.Bold != nil {
		finds := reg.Bold.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Bold,
			}
			matches = append(matches, m)
		}
	}

	if reg.Comment != nil {
		finds := reg.Comment.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Comment,
			}
			matches = append(matches, m)
		}
	}

	if reg.Heading != nil {
		finds := reg.Heading.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Heading,
			}
			matches = append(matches, m)
		}
	}

	if reg.Italic != nil {
		finds := reg.Italic.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Italic,
			}
			matches = append(matches, m)
		}
	}

	if reg.Keyword != nil {
		finds := reg.Keyword.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Keyword,
			}
			matches = append(matches, m)
		}
	}

	if reg.Meta != nil {
		finds := reg.Meta.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Meta,
			}
			matches = append(matches, m)
		}
	}

	if reg.Mono != nil {
		finds := reg.Mono.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Mono,
			}
			matches = append(matches, m)
		}
	}

	if reg.Number != nil {
		finds := reg.Number.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Number,
			}
			matches = append(matches, m)
		}
	}

	if reg.Operator != nil {
		finds := reg.Operator.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Operator,
			}
			matches = append(matches, m)
		}
	}

	if reg.String != nil {
		finds := reg.String.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.String,
			}
			matches = append(matches, m)
		}
	}

	if reg.Type != nil {
		finds := reg.Type.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Type,
			}
			matches = append(matches, m)
		}
	}

	if reg.Value != nil {
		finds := reg.Value.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Value,
			}
			matches = append(matches, m)
		}
	}

	if reg.Variable != nil {
		finds := reg.Variable.FindAll(buf)
		for _, f := range finds {
			m := match{
				start: util.ByteIdxToRuneIdx(buf, f.Start),
				end:   util.ByteIdxToRuneIdx(buf, f.End),
				style: cfg.Style.Variable,
			}
			matches = append(matches, m)
		}
	}

	less := func(i, j int) bool {
		mi := matches[i]
		mj := matches[j]

		return mi.start < mj.start
	}
	sort.Slice(matches, less)

	return matches
}
