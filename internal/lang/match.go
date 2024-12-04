package lang

type match struct {
	start int
	end   int
}

type matches struct {
	CursorWords []match

	Attributes       []match
	Builtins         []match
	Bolds            []match
	Codes            []match
	Comments         []match
	Documentations   []match
	EscapeSequences  []match
	FormatSpecifiers []match
	Functions        []match
	Headings         []match
	Italics          []match
	Keywords         []match
	Links            []match
	Metas            []match
	Monos            []match
	Numbers          []match
	Operators        []match
	Strings          []match
	ToDos            []match
	Titles           []match
	Types            []match
	Values           []match
	Variables        []match
}

func defaultMatches() matches {
	return matches{
		CursorWords:      nil,
		Attributes:       nil,
		Builtins:         nil,
		Bolds:            nil,
		Codes:            nil,
		Comments:         nil,
		Documentations:   nil,
		EscapeSequences:  nil,
		FormatSpecifiers: nil,
		Functions:        nil,
		Headings:         nil,
		Italics:          nil,
		Keywords:         nil,
		Links:            nil,
		Metas:            nil,
		Monos:            nil,
		Numbers:          nil,
		Operators:        nil,
		Strings:          nil,
		ToDos:            nil,
		Titles:           nil,
		Types:            nil,
		Values:           nil,
		Variables:        nil,
	}
}
