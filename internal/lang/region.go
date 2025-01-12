package lang

import (
	"regexp"

	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/regex"
	"github.com/m-kru/enix/internal/util"
)

type Region struct {
	Name  string
	Style string // Region default style

	Start *regex.Regex
	End   *regex.Regex

	CursorWord *regexp.Regexp

	Attribute       *regex.Regex
	Builtin         *regex.Regex
	Bold            *regex.Regex
	Comment         *regex.Regex
	EscapeSequence  *regex.Regex
	FormatSpecifier *regex.Regex
	Function        *regex.Regex
	Heading         *regex.Regex
	Italic          *regex.Regex
	Keyword         *regex.Regex
	Link            *regex.Regex
	Meta            *regex.Regex
	Mono            *regex.Regex
	Number          *regex.Regex
	Operator        *regex.Regex
	String          *regexp.Regexp
	ToDo            *regexp.Regexp
	Type            *regexp.Regexp
	Value           *regexp.Regexp
	Variable        *regexp.Regexp
}

func DefaultRegion() *Region {
	return &Region{
		Name:            "Default",
		Style:           "",
		Start:           nil,
		End:             nil,
		CursorWord:      nil,
		Attribute:       nil,
		Builtin:         nil,
		Bold:            nil,
		Comment:         nil,
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
		attrs := reg.Attribute.FindAll(buf)
		if len(attrs) > 0 {
			matches.Attributes = make([]match, len(attrs))
			for i, a := range attrs {
				matches.Attributes[i].start = util.ByteIdxToRuneIdx(buf, a.Start) + startIdx
				matches.Attributes[i].end = util.ByteIdxToRuneIdx(buf, a.End) + startIdx
			}
		}
	}

	if reg.Builtin != nil {
		builtins := reg.Builtin.FindAll(buf)
		if len(builtins) > 0 {
			matches.Builtins = make([]match, len(builtins))
			for i, b := range builtins {
				matches.Builtins[i].start = util.ByteIdxToRuneIdx(buf, b.Start) + startIdx
				matches.Builtins[i].end = util.ByteIdxToRuneIdx(buf, b.End) + startIdx
			}
		}
	}

	if reg.Bold != nil {
		bolds := reg.Bold.FindAll(buf)
		if len(bolds) > 0 {
			matches.Bolds = make([]match, len(bolds))
			for i, b := range bolds {
				matches.Bolds[i].start = util.ByteIdxToRuneIdx(buf, b.Start) + startIdx
				matches.Bolds[i].end = util.ByteIdxToRuneIdx(buf, b.End) + startIdx
			}
		}
	}

	if reg.Comment != nil {
		comments := reg.Comment.FindAll(buf)
		if len(comments) > 0 {
			matches.Comments = make([]match, len(comments))
			for i, c := range comments {
				matches.Comments[i].start = util.ByteIdxToRuneIdx(buf, c.Start) + startIdx
				matches.Comments[i].end = util.ByteIdxToRuneIdx(buf, c.End) + startIdx
			}
		}
	}

	if reg.EscapeSequence != nil {
		seqs := reg.EscapeSequence.FindAll(buf)
		if len(seqs) > 0 {
			matches.EscapeSequences = make([]match, len(seqs))
			for i, s := range seqs {
				matches.EscapeSequences[i].start = util.ByteIdxToRuneIdx(buf, s.Start) + startIdx
				matches.EscapeSequences[i].end = util.ByteIdxToRuneIdx(buf, s.End) + startIdx
			}
		}
	}

	if reg.FormatSpecifier != nil {
		fmts := reg.FormatSpecifier.FindAll(buf)
		if len(fmts) > 0 {
			matches.FormatSpecifiers = make([]match, len(fmts))
			for i, f := range fmts {
				matches.FormatSpecifiers[i].start = util.ByteIdxToRuneIdx(buf, f.Start) + startIdx
				matches.FormatSpecifiers[i].end = util.ByteIdxToRuneIdx(buf, f.End) + startIdx
			}
		}
	}

	if reg.Function != nil {
		funcs := reg.Function.FindAll(buf)
		if len(funcs) > 0 {
			matches.Functions = make([]match, len(funcs))
			for i, f := range funcs {
				matches.Functions[i].start = util.ByteIdxToRuneIdx(buf, f.Start) + startIdx
				matches.Functions[i].end = util.ByteIdxToRuneIdx(buf, f.End) + startIdx
			}
		}
	}

	if reg.Heading != nil {
		headings := reg.Heading.FindAll(buf)
		if len(headings) > 0 {
			matches.Headings = make([]match, len(headings))
			for i, h := range headings {
				matches.Headings[i].start = util.ByteIdxToRuneIdx(buf, h.Start) + startIdx
				matches.Headings[i].end = util.ByteIdxToRuneIdx(buf, h.End) + startIdx
			}
		}
	}

	if reg.Italic != nil {
		italics := reg.Italic.FindAll(buf)
		if len(italics) > 0 {
			matches.Italics = make([]match, len(italics))
			for i, it := range italics {
				matches.Italics[i].start = util.ByteIdxToRuneIdx(buf, it.Start) + startIdx
				matches.Italics[i].end = util.ByteIdxToRuneIdx(buf, it.End) + startIdx
			}
		}
	}

	if reg.Keyword != nil {
		keywords := reg.Keyword.FindAll(buf)
		if len(keywords) > 0 {
			matches.Keywords = make([]match, len(keywords))
			for i, k := range keywords {
				matches.Keywords[i].start = util.ByteIdxToRuneIdx(buf, k.Start) + startIdx
				matches.Keywords[i].end = util.ByteIdxToRuneIdx(buf, k.End) + startIdx
			}
		}
	}

	if reg.Link != nil {
		links := reg.Link.FindAll(buf)
		if len(links) > 0 {
			matches.Links = make([]match, len(links))
			for i, l := range links {
				matches.Links[i].start = util.ByteIdxToRuneIdx(buf, l.Start) + startIdx
				matches.Links[i].end = util.ByteIdxToRuneIdx(buf, l.End) + startIdx
			}
		}
	}

	if reg.Meta != nil {
		metas := reg.Meta.FindAll(buf)
		if len(metas) > 0 {
			matches.Metas = make([]match, len(metas))
			for i, m := range metas {
				matches.Metas[i].start = util.ByteIdxToRuneIdx(buf, m.Start) + startIdx
				matches.Metas[i].end = util.ByteIdxToRuneIdx(buf, m.End) + startIdx
			}
		}
	}

	if reg.Mono != nil {
		monos := reg.Mono.FindAll(buf)
		if len(monos) > 0 {
			matches.Monos = make([]match, len(monos))
			for i, m := range monos {
				matches.Metas[i].start = util.ByteIdxToRuneIdx(buf, m.Start) + startIdx
				matches.Metas[i].end = util.ByteIdxToRuneIdx(buf, m.End) + startIdx
			}
		}
	}

	if reg.Number != nil {
		nums := reg.Number.FindAll(buf)
		if len(nums) > 0 {
			matches.Numbers = make([]match, len(nums))
			for i, n := range nums {
				matches.Numbers[i].start = util.ByteIdxToRuneIdx(buf, n.Start) + startIdx
				matches.Numbers[i].end = util.ByteIdxToRuneIdx(buf, n.End) + startIdx
			}
		}
	}

	if reg.Operator != nil {
		ops := reg.Operator.FindAll(buf)
		if len(ops) > 0 {
			matches.Operators = make([]match, len(ops))
			for i, o := range ops {
				matches.Operators[i].start = util.ByteIdxToRuneIdx(buf, o.Start) + startIdx
				matches.Operators[i].end = util.ByteIdxToRuneIdx(buf, o.End) + startIdx
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
