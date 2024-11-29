package lang

import (
	"regexp"
)

type Match struct {
	start int
	end   int
}

type Regex struct {
	Regex              *regexp.Regexp
	NegativeLookbehind *regexp.Regexp
	PositiveLookahead  *regexp.Regexp
}

func (r Regex) FindAll(buf []byte) []Match {
	matches := make([]Match, 0, 8)
	var negLookBeh [][]int
	var posLookAhead [][]int

	finds := r.Regex.FindAllIndex(buf, -1)

	if len(finds) > 0 {
		if r.NegativeLookbehind != nil {
			negLookBeh = r.NegativeLookbehind.FindAllIndex(buf, -1)
		}
		if r.PositiveLookahead != nil {
			posLookAhead = r.PositiveLookahead.FindAllIndex(buf, -1)
		}
	}

	for _, f := range finds {
		ok := true
		for _, nlb := range negLookBeh {
			if nlb[1] == f[0] {
				ok = false
				break
			}
		}
		if !ok {
			continue
		}

		for i, pla := range posLookAhead {
			if f[1] == pla[0] {
				break
			}
			if i == len(posLookAhead)-1 {
				ok = false
			}
		}
		if !ok {
			continue
		}

		matches = append(matches, Match{start: f[0], end: f[1]})
	}

	return matches
}
