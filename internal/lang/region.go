package lang

import (
	"regexp"

	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/util"
)

type Region struct {
	Name  string
	Style string // Region default style

	Start Regex
	End   Regex

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

func DefaultRegion() *Region {
	return &Region{
		Name:  "Default",
		Style: "",
		Start: Regex{
			Regex:              nil,
			NegativeLookbehind: nil,
			PositiveLookahead:  nil,
		},
		End: Regex{
			Regex:              nil,
			NegativeLookbehind: nil,
			PositiveLookahead:  nil,
		},
		CursorWord:      nil,
		Attribute:       nil,
		Builtin:         nil,
		Bold:            nil,
		Code:            nil,
		Comment:         nil,
		Documentation:   nil,
		EscapeSequence:  nil,
		FormatSpecifier: nil,
		Function:        nil,
		Heading:         nil,
		Italic:          nil,
		Keyword:         nil,
		Link:            nil,
		Meta:            nil,
		Mono:            nil,
		Number:          nil,
		Operator:        nil,
		String:          nil,
		ToDo:            nil,
		Title:           nil,
		Type:            nil,
		Value:           nil,
		Variable:        nil,
	}
}

func (reg Region) Match(line *line.Line, startIdx int, endIdx int) matches {
	matches := defaultMatches()

	buf := line.Buf[startIdx:endIdx]

	if reg.CursorWord != nil {
		words := reg.CursorWord.FindAllIndex(buf, -1)
		if len(words) > 0 {
			matches.CursorWords = make([]match, 0, len(words))
			for _, w := range words {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, w[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, w[1]) + startIdx
				matches.CursorWords = append(matches.CursorWords, m)
			}
		}
	}

	if reg.Attribute != nil {
		attrs := reg.Attribute.FindAllIndex(buf, -1)
		if len(attrs) > 0 {
			matches.Attributes = make([]match, 0, len(attrs))
			for _, a := range attrs {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, a[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, a[1]) + startIdx
				matches.Attributes = append(matches.Attributes, m)
			}
		}
	}

	if reg.Builtin != nil {
		builtins := reg.Builtin.FindAllIndex(buf, -1)
		if len(builtins) > 0 {
			matches.Builtins = make([]match, 0, len(builtins))
			for _, b := range builtins {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, b[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, b[1]) + startIdx
				matches.Builtins = append(matches.Builtins, m)
			}
		}
	}

	if reg.EscapeSequence != nil {
		seqs := reg.EscapeSequence.FindAllIndex(buf, -1)
		if len(seqs) > 0 {
			matches.EscapeSequences = make([]match, 0, len(seqs))
			for _, s := range seqs {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, s[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, s[1]) + startIdx
				matches.EscapeSequences = append(matches.EscapeSequences, m)
			}
		}
	}

	if reg.FormatSpecifier != nil {
		fmts := reg.FormatSpecifier.FindAllIndex(buf, -1)
		if len(fmts) > 0 {
			matches.FormatSpecifiers = make([]match, 0, len(fmts))
			for _, f := range fmts {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, f[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, f[1]) + startIdx
				matches.FormatSpecifiers = append(matches.FormatSpecifiers, m)
			}
		}
	}

	if reg.Function != nil {
		funcs := reg.Function.FindAllIndex(buf, -1)
		if len(funcs) > 0 {
			matches.Functions = make([]match, 0, len(funcs))
			for _, f := range funcs {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, f[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, f[1]) + startIdx
				matches.Functions = append(matches.Functions, m)
			}
		}
	}

	if reg.Keyword != nil {
		keywords := reg.Keyword.FindAllIndex(buf, -1)
		if len(keywords) > 0 {
			matches.Keywords = make([]match, 0, len(keywords))
			for _, k := range keywords {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, k[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, k[1]) + startIdx
				matches.Keywords = append(matches.Keywords, m)
			}
		}
	}

	if reg.Meta != nil {
		metas := reg.Meta.FindAllIndex(buf, -1)
		if len(metas) > 0 {
			matches.Metas = make([]match, 0, len(metas))
			for _, x := range metas {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, x[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, x[1]) + startIdx
				matches.Metas = append(matches.Metas, m)
			}
		}
	}

	if reg.Number != nil {
		nums := reg.Number.FindAllIndex(buf, -1)
		if len(nums) > 0 {
			matches.Numbers = make([]match, 0, len(nums))
			for _, n := range nums {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, n[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, n[1]) + startIdx
				matches.Numbers = append(matches.Numbers, m)
			}
		}
	}

	if reg.Operator != nil {
		ops := reg.Operator.FindAllIndex(buf, -1)
		if len(ops) > 0 {
			matches.Operators = make([]match, 0, len(ops))
			for _, o := range ops {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, o[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, o[1]) + startIdx
				matches.Operators = append(matches.Operators, m)
			}
		}
	}

	if reg.String != nil {
		strs := reg.String.FindAllIndex(buf, -1)
		if len(strs) > 0 {
			matches.Strings = make([]match, 0, len(strs))
			for _, s := range strs {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, s[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, s[1]) + startIdx
				matches.Strings = append(matches.Strings, m)
			}
		}
	}

	if reg.ToDo != nil {
		todos := reg.ToDo.FindAllIndex(buf, -1)
		if len(todos) > 0 {
			matches.ToDos = make([]match, 0, len(todos))
			for _, t := range todos {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, t[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, t[1]) + startIdx
				matches.ToDos = append(matches.ToDos, m)
			}
		}
	}

	if reg.Type != nil {
		types := reg.Type.FindAllIndex(buf, -1)
		if len(types) > 0 {
			matches.Types = make([]match, 0, len(types))
			for _, t := range types {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, t[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, t[1]) + startIdx
				matches.Types = append(matches.Types, m)
			}
		}
	}

	if reg.Value != nil {
		values := reg.Value.FindAllIndex(buf, -1)
		if len(values) > 0 {
			matches.Values = make([]match, 0, len(values))
			for _, v := range values {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, v[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, v[1]) + startIdx
				matches.Values = append(matches.Values, m)
			}
		}
	}

	if reg.Variable != nil {
		vars := reg.Variable.FindAllIndex(buf, -1)
		if len(vars) > 0 {
			matches.Variables = make([]match, 0, len(vars))
			for _, v := range vars {
				var m match
				m.start = util.ByteIdxToRuneIdx(buf, v[0]) + startIdx
				m.end = util.ByteIdxToRuneIdx(buf, v[1]) + startIdx
				matches.Variables = append(matches.Variables, m)
			}
		}
	}

	return matches
}
