package lang

type match struct {
	start int
	end   int
}

type matches struct {
	CursorWords []match

	Attributes []match
	Bolds      []match
	Comments   []match
	Headings   []match
	Italics    []match
	Keywords   []match
	Links      []match
	Metas      []match
	Monos      []match
	Numbers    []match
	Operators  []match
	Strings    []match
	Types      []match
	Values     []match
	Variables  []match
}

func defaultMatches() matches {
	return matches{
		CursorWords: nil,
		Attributes:  nil,
		Bolds:       nil,
		Comments:    nil,
		Headings:    nil,
		Italics:     nil,
		Keywords:    nil,
		Links:       nil,
		Metas:       nil,
		Monos:       nil,
		Numbers:     nil,
		Operators:   nil,
		Strings:     nil,
		Types:       nil,
		Values:      nil,
		Variables:   nil,
	}
}
