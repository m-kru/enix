package search

import (
	"github.com/m-kru/enix/internal/find"
	"regexp"
)

type Context struct {
	PrevRegexp      *regexp.Regexp
	Regexp          *regexp.Regexp
	FirstVisLineNum int  // Line number of the first visible line
	Modified        bool // Flag indicating whether the search content was modified since last search
	Finds           []find.Find
	FirstVisFindIdx int // Index of potentially first visible find
}

func InitialContext() Context {
	return Context{
		PrevRegexp:      nil,
		Regexp:          nil,
		FirstVisLineNum: 0,
		Modified:        true, // true required for the first search to work correctly
		Finds:           nil,
		FirstVisFindIdx: 0,
	}
}

func (ctx Context) FindsFromVisible() []find.Find {
	if ctx.Regexp == nil || ctx.FirstVisFindIdx < 0 {
		return nil
	}

	return ctx.Finds[ctx.FirstVisFindIdx:]
}
