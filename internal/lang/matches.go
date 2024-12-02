package lang

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

func DefaultMatches() Matches {
	return Matches{
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
