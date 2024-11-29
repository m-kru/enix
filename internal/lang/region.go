package lang

import (
	"regexp"

	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

type RegionStartRegex struct {
	Regex              *regexp.Regexp
	NegativeLookbehind *regexp.Regexp
}

type RegionEndRegex struct {
	Regex              *regexp.Regexp
	NegativeLookbehind *regexp.Regexp
}

type Region struct {
	Name  string
	Style string // Region default style

	StartRegex RegionStartRegex
	EndRegex   RegionEndRegex

	CursorWord *regexp.Regexp

	Attribute       *regexp.Regexp
	Builtin         *regexp.Regexp
	Bold            *regexp.Regexp
	Code            *regexp.Regexp
	Comment         *regexp.Regexp
	Documentation   *regexp.Regexp
	EscapeSequence  *regexp.Regexp
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
}

type Matches struct {
	CursorWords [][2]int

	Attributes       [][2]int
	Builtins         [][2]int
	Bolds            [][2]int
	Codes            [][2]int
	Comments         [][2]int
	Documentations   [][2]int
	EscapeSequences  [][2]int
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

	if reg.Attribute != nil {
		attrs := reg.Attribute.FindAllStringIndex(str, -1)
		if len(attrs) > 0 {
			matches.Attributes = make([][2]int, 0, len(attrs))
			for _, a := range attrs {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, a[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, a[1]) + startIdx
				matches.Attributes = append(matches.Attributes, m)
			}
		}
	}

	if reg.Builtin != nil {
		builtins := reg.Builtin.FindAllStringIndex(str, -1)
		if len(builtins) > 0 {
			matches.Builtins = make([][2]int, 0, len(builtins))
			for _, b := range builtins {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, b[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, b[1]) + startIdx
				matches.Builtins = append(matches.Builtins, m)
			}
		}
	}

	if reg.EscapeSequence != nil {
		seqs := reg.EscapeSequence.FindAllStringIndex(str, -1)
		if len(seqs) > 0 {
			matches.EscapeSequences = make([][2]int, 0, len(seqs))
			for _, s := range seqs {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, s[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, s[1]) + startIdx
				matches.EscapeSequences = append(matches.EscapeSequences, m)
			}
		}
	}

	if reg.FormatSpecifier != nil {
		fmts := reg.FormatSpecifier.FindAllStringIndex(str, -1)
		if len(fmts) > 0 {
			matches.FormatSpecifiers = make([][2]int, 0, len(fmts))
			for _, f := range fmts {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, f[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, f[1]) + startIdx
				matches.FormatSpecifiers = append(matches.FormatSpecifiers, m)
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

	if reg.String != nil {
		strs := reg.String.FindAllStringIndex(str, -1)
		if len(strs) > 0 {
			matches.Strings = make([][2]int, 0, len(strs))
			for _, s := range strs {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, s[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, s[1]) + startIdx
				matches.Strings = append(matches.Strings, m)
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

	if reg.Variable != nil {
		vars := reg.Variable.FindAllStringIndex(str, -1)
		if len(vars) > 0 {
			matches.Variables = make([][2]int, 0, len(vars))
			for _, v := range vars {
				var m [2]int
				m[0] = util.ByteIdxToRuneIdx(str, v[0]) + startIdx
				m[1] = util.ByteIdxToRuneIdx(str, v[1]) + startIdx
				matches.Variables = append(matches.Variables, m)
			}
		}
	}

	return matches
}
