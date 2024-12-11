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
			NegativeLookBehind: nil,
			NegativeLookAhead:  nil,
			PositiveLookAhead:  nil,
		},
		End: Regex{
			Regex:              nil,
			NegativeLookBehind: nil,
			NegativeLookAhead:  nil,
			PositiveLookAhead:  nil,
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

func (reg Region) match(line *line.Line, startIdx int, endIdx int) matches {
	matches := defaultMatches()

	buf := line.Buf[startIdx:endIdx]

	if reg.CursorWord != nil {
		words := reg.CursorWord.FindAllIndex(buf, -1)
		if len(words) > 0 {
			matches.CursorWords = make([]match, len(words))
			for i, w := range words {
				matches.CursorWords[i].start = util.ByteIdxToRuneIdx(buf, w[0]) + startIdx
				matches.CursorWords[i].end = util.ByteIdxToRuneIdx(buf, w[1]) + startIdx
			}
		}
	}

	if reg.Attribute != nil {
		attrs := reg.Attribute.FindAllIndex(buf, -1)
		if len(attrs) > 0 {
			matches.Attributes = make([]match, len(attrs))
			for i, a := range attrs {
				matches.Attributes[i].start = util.ByteIdxToRuneIdx(buf, a[0]) + startIdx
				matches.Attributes[i].end = util.ByteIdxToRuneIdx(buf, a[1]) + startIdx
			}
		}
	}

	if reg.Builtin != nil {
		builtins := reg.Builtin.FindAllIndex(buf, -1)
		if len(builtins) > 0 {
			matches.Builtins = make([]match, len(builtins))
			for i, b := range builtins {
				matches.Builtins[i].start = util.ByteIdxToRuneIdx(buf, b[0]) + startIdx
				matches.Builtins[i].end = util.ByteIdxToRuneIdx(buf, b[1]) + startIdx
			}
		}
	}

	if reg.EscapeSequence != nil {
		seqs := reg.EscapeSequence.FindAllIndex(buf, -1)
		if len(seqs) > 0 {
			matches.EscapeSequences = make([]match, len(seqs))
			for i, s := range seqs {
				matches.EscapeSequences[i].start = util.ByteIdxToRuneIdx(buf, s[0]) + startIdx
				matches.EscapeSequences[i].end = util.ByteIdxToRuneIdx(buf, s[1]) + startIdx
			}
		}
	}

	if reg.FormatSpecifier != nil {
		fmts := reg.FormatSpecifier.FindAllIndex(buf, -1)
		if len(fmts) > 0 {
			matches.FormatSpecifiers = make([]match, len(fmts))
			for i, f := range fmts {
				matches.FormatSpecifiers[i].start = util.ByteIdxToRuneIdx(buf, f[0]) + startIdx
				matches.FormatSpecifiers[i].end = util.ByteIdxToRuneIdx(buf, f[1]) + startIdx
			}
		}
	}

	if reg.Function != nil {
		funcs := reg.Function.FindAllIndex(buf, -1)
		if len(funcs) > 0 {
			matches.Functions = make([]match, len(funcs))
			for i, f := range funcs {
				matches.Functions[i].start = util.ByteIdxToRuneIdx(buf, f[0]) + startIdx
				matches.Functions[i].end = util.ByteIdxToRuneIdx(buf, f[1]) + startIdx
			}
		}
	}

	if reg.Keyword != nil {
		keywords := reg.Keyword.FindAllIndex(buf, -1)
		if len(keywords) > 0 {
			matches.Keywords = make([]match, len(keywords))
			for i, k := range keywords {
				matches.Keywords[i].start = util.ByteIdxToRuneIdx(buf, k[0]) + startIdx
				matches.Keywords[i].end = util.ByteIdxToRuneIdx(buf, k[1]) + startIdx
			}
		}
	}

	if reg.Meta != nil {
		metas := reg.Meta.FindAllIndex(buf, -1)
		if len(metas) > 0 {
			matches.Metas = make([]match, len(metas))
			for i, x := range metas {
				matches.Metas[i].start = util.ByteIdxToRuneIdx(buf, x[0]) + startIdx
				matches.Metas[i].end = util.ByteIdxToRuneIdx(buf, x[1]) + startIdx
			}
		}
	}

	if reg.Number != nil {
		nums := reg.Number.FindAllIndex(buf, -1)
		if len(nums) > 0 {
			matches.Numbers = make([]match, len(nums))
			for i, n := range nums {
				matches.Numbers[i].start = util.ByteIdxToRuneIdx(buf, n[0]) + startIdx
				matches.Numbers[i].end = util.ByteIdxToRuneIdx(buf, n[1]) + startIdx
			}
		}
	}

	if reg.Operator != nil {
		ops := reg.Operator.FindAllIndex(buf, -1)
		if len(ops) > 0 {
			matches.Operators = make([]match, len(ops))
			for i, o := range ops {
				matches.Operators[i].start = util.ByteIdxToRuneIdx(buf, o[0]) + startIdx
				matches.Operators[i].end = util.ByteIdxToRuneIdx(buf, o[1]) + startIdx
			}
		}
	}

	if reg.String != nil {
		strs := reg.String.FindAllIndex(buf, -1)
		if len(strs) > 0 {
			matches.Strings = make([]match, len(strs))
			for i, s := range strs {
				matches.Strings[i].start = util.ByteIdxToRuneIdx(buf, s[0]) + startIdx
				matches.Strings[i].end = util.ByteIdxToRuneIdx(buf, s[1]) + startIdx
			}
		}
	}

	if reg.ToDo != nil {
		todos := reg.ToDo.FindAllIndex(buf, -1)
		if len(todos) > 0 {
			matches.ToDos = make([]match, len(todos))
			for i, t := range todos {
				matches.ToDos[i].start = util.ByteIdxToRuneIdx(buf, t[0]) + startIdx
				matches.ToDos[i].end = util.ByteIdxToRuneIdx(buf, t[1]) + startIdx
			}
		}
	}

	if reg.Type != nil {
		types := reg.Type.FindAllIndex(buf, -1)
		if len(types) > 0 {
			matches.Types = make([]match, len(types))
			for i, t := range types {
				matches.Types[i].start = util.ByteIdxToRuneIdx(buf, t[0]) + startIdx
				matches.Types[i].end = util.ByteIdxToRuneIdx(buf, t[1]) + startIdx
			}
		}
	}

	if reg.Value != nil {
		values := reg.Value.FindAllIndex(buf, -1)
		if len(values) > 0 {
			matches.Values = make([]match, len(values))
			for i, v := range values {
				matches.Values[i].start = util.ByteIdxToRuneIdx(buf, v[0]) + startIdx
				matches.Values[i].end = util.ByteIdxToRuneIdx(buf, v[1]) + startIdx
			}
		}
	}

	if reg.Variable != nil {
		vars := reg.Variable.FindAllIndex(buf, -1)
		if len(vars) > 0 {
			matches.Variables = make([]match, len(vars))
			for i, v := range vars {
				matches.Variables[i].start = util.ByteIdxToRuneIdx(buf, v[0]) + startIdx
				matches.Variables[i].end = util.ByteIdxToRuneIdx(buf, v[1]) + startIdx
			}
		}
	}

	return matches
}
