package lang

import (
	"regexp"

	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

type Region struct {
	Name  string
	Style string // Region default style

	StartRegexp *regexp.Regexp
	EndRegexp   *regexp.Regexp

	CursorWord *regexp.Regexp

	Attribute       *regexp.Regexp
	Builtin         *regexp.Regexp
	Bold            *regexp.Regexp
	Code            *regexp.Regexp
	Comment         *regexp.Regexp
	Documentation   *regexp.Regexp
	FormatSpecifier *regexp.Regexp
	Function        *regexp.Regexp
	Heading         *regexp.Regexp
	Italic          *regexp.Regexp
	Keyword         *regexp.Regexp
	Link            *regexp.Regexp
	Meta            *regexp.Regexp
	Mono            *regexp.Regexp
	Number          *regexp.Regexp
	Operator        *regexp.Regexp
	String          *regexp.Regexp
	ToDo            *regexp.Regexp
	Title           *regexp.Regexp
	Type            *regexp.Regexp
	Value           *regexp.Regexp
	Variable        *regexp.Regexp
	Other           *regexp.Regexp
}

type Matches struct {
	CursorWords [][2]int

	Attributes       [][2]int
	Builtins         [][2]int
	Bolds            [][2]int
	Codes            [][2]int
	Comments         [][2]int
	Documentations   [][2]int
	FormatSpecifiers [][2]int
	Functions        [][2]int
	Headings         [][2]int
	Italics          [][2]int
	Keywords         [][2]int
	Links            [][2]int
	Metas            [][2]int
	Monos            [][2]int
	Numbers          [][2]int
	Operators        [][2]int
	Strings          [][2]int
	ToDos            [][2]int
	Titles           [][2]int
	Types            [][2]int
	Values           [][2]int
	Variables        [][2]int
	Others           [][2]int
}

func (reg Region) Match(line *line.Line, startIdx int, endIdx int) Matches {
	matches := Matches{}

	str := string(line.Buf[startIdx:endIdx])

	if reg.CursorWord != nil {
		words := reg.CursorWord.FindAllStringIndex(str, -1)
		if len(words) > 0 {
			matches.CursorWords = make([][2]int, 0, len(words))
			for _, w := range words {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, w[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, w[1]) + startIdx
				matches.CursorWords = append(matches.CursorWords, m)
			}
		}
	}

	if reg.Keyword != nil {
		keywords := reg.Keyword.FindAllStringIndex(str, -1)
		if len(keywords) > 0 {
			matches.Keywords = make([][2]int, 0, len(keywords))
			for _, k := range keywords {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, k[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, k[1]) + startIdx
				matches.Keywords = append(matches.Keywords, m)
			}
		}
	}

	if reg.Meta != nil {
		metas := reg.Meta.FindAllStringIndex(str, -1)
		if len(metas) > 0 {
			matches.Metas = make([][2]int, 0, len(metas))
			for _, x := range metas {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, x[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, x[1]) + startIdx
				matches.Metas = append(matches.Metas, m)
			}
		}
	}

	if reg.Number != nil {
		nums := reg.Number.FindAllStringIndex(str, -1)
		if len(nums) > 0 {
			matches.Numbers = make([][2]int, 0, len(nums))
			for _, n := range nums {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, n[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, n[1]) + startIdx
				matches.Numbers = append(matches.Numbers, m)
			}
		}
	}

	if reg.Operator != nil {
		ops := reg.Operator.FindAllStringIndex(str, -1)
		if len(ops) > 0 {
			matches.Operators = make([][2]int, 0, len(ops))
			for _, o := range ops {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, o[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, o[1]) + startIdx
				matches.Operators = append(matches.Operators, m)
			}
		}
	}

	if reg.Type != nil {
		types := reg.Type.FindAllStringIndex(str, -1)
		if len(types) > 0 {
			matches.Types = make([][2]int, 0, len(types))
			for _, t := range types {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, t[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, t[1]) + startIdx
				matches.Types = append(matches.Types, m)
			}
		}
	}

	if reg.Value != nil {
		values := reg.Value.FindAllStringIndex(str, -1)
		if len(values) > 0 {
			matches.Values = make([][2]int, 0, len(values))
			for _, v := range values {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, v[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, v[1]) + startIdx
				matches.Values = append(matches.Values, m)
			}
		}
	}

	return matches
}
