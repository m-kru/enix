package highlight

import (
	"regexp"
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

type RegionToken struct {
	Region *Region
	Start  bool // Start (true) or end (false) token
	// Token start index for start token or token end index
	// for end token.
	Idx int
}
