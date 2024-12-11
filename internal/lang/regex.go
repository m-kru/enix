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
	NegativeLookahead  *regexp.Regexp
	PositiveLookahead  *regexp.Regexp
}

func (r Regex) FindAll(buf []byte) []Match {
	matches := make([]Match, 0, 8)
	var negLookBeh [][]int
	var negLookAhead [][]int
	var posLookAhead [][]int

	finds := r.Regex.FindAllIndex(buf, -1)

	if len(finds) > 0 {
		if r.NegativeLookbehind != nil {
			negLookBeh = r.NegativeLookbehind.FindAllIndex(buf, -1)
		}
		if r.NegativeLookahead != nil {
			negLookAhead = r.NegativeLookahead.FindAllIndex(buf, -1)
		}
		if r.PositiveLookahead != nil {
			posLookAhead = r.PositiveLookahead.FindAllIndex(buf, -1)
		}
	}

	// Note: Below code can be optimized.
	// If i'th find consumed j'th lookaround,
	// the for (i+1)'th find we can start from (j+1)'th lookaround.
	for _, f := range finds {
		nlbFound := false
		for _, nlb := range negLookBeh {
			if nlb[1] == f[0] {
				nlbFound = true
				break
			}
		}

		nlaFound := false
		for _, nla := range negLookAhead {
			if f[1] == nla[0] {
				nlaFound = true
				break
			}
		}

		plaFound := false
		for _, pla := range posLookAhead {
			if f[1] == pla[0] {
				plaFound = true
				break
			}
		}

		if r.NegativeLookbehind != nil && r.NegativeLookahead == nil && r.PositiveLookahead == nil {
			if nlbFound {
				continue
			}
		}
		if r.NegativeLookbehind != nil && r.NegativeLookahead != nil && r.PositiveLookahead == nil {
			if nlbFound && nlaFound {
				continue
			}
		}
		if r.NegativeLookbehind != nil && r.NegativeLookahead == nil && r.PositiveLookahead != nil {
			if nlbFound && !plaFound {
				continue
			}
		}
		if r.NegativeLookbehind != nil && r.NegativeLookahead != nil && r.PositiveLookahead != nil {
			if nlbFound && nlaFound && !plaFound {
				continue
			}
		}
		if r.NegativeLookbehind == nil && r.NegativeLookahead != nil && r.PositiveLookahead == nil {
			if nlaFound {
				continue
			}
		}
		if r.NegativeLookbehind == nil && r.NegativeLookahead != nil && r.PositiveLookahead != nil {
			if nlaFound && !plaFound {
				continue
			}
		}
		if r.NegativeLookbehind == nil && r.NegativeLookahead == nil && r.PositiveLookahead != nil {
			if !plaFound {
				continue
			}
		}

		matches = append(matches, Match{start: f[0], end: f[1]})
	}

	return matches
}
