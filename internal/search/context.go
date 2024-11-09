package search

import (
	"github.com/m-kru/enix/internal/find"
	"regexp"
)

type Context struct {
	Regexp   *regexp.Regexp
	Finds    []find.Find
	StartIdx int // Index of potentially first visible find
}

func (ctx Context) FindsFromVisible() []find.Find {
	if ctx.StartIdx < 0 {
		return []find.Find{}
	}

	return ctx.Finds[ctx.StartIdx:]
}
