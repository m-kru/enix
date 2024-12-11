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
	NegativeLookBehind *regexp.Regexp
	NegativeLookAhead  *regexp.Regexp
	PositiveLookAhead  *regexp.Regexp
}

func (r Regex) FindAll(buf []byte) []Match {
	matches := make([]Match, 0, 8)
	var negLookBeh [][]int
	var negLookAhead [][]int
	var posLookAhead [][]int

	finds := r.Regex.FindAllIndex(buf, -1)

	if len(finds) > 0 {
		if r.NegativeLookBehind != nil {
			negLookBeh = r.NegativeLookBehind.FindAllIndex(buf, -1)
		}
		if r.NegativeLookAhead != nil {
			negLookAhead = r.NegativeLookAhead.FindAllIndex(buf, -1)
		}
		if r.PositiveLookAhead != nil {
			posLookAhead = r.PositiveLookAhead.FindAllIndex(buf, -1)
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

		if r.NegativeLookBehind != nil && r.NegativeLookAhead == nil && r.PositiveLookAhead == nil {
			if nlbFound {
				continue
			}
		}
		if r.NegativeLookBehind != nil && r.NegativeLookAhead != nil && r.PositiveLookAhead == nil {
			if nlbFound && nlaFound {
				continue
			}
		}
		if r.NegativeLookBehind != nil && r.NegativeLookAhead == nil && r.PositiveLookAhead != nil {
			if nlbFound && !plaFound {
				continue
			}
		}
		if r.NegativeLookBehind != nil && r.NegativeLookAhead != nil && r.PositiveLookAhead != nil {
			if nlbFound && nlaFound && !plaFound {
				continue
			}
		}
		if r.NegativeLookBehind == nil && r.NegativeLookAhead != nil && r.PositiveLookAhead == nil {
			if nlaFound {
				continue
			}
		}
		if r.NegativeLookBehind == nil && r.NegativeLookAhead != nil && r.PositiveLookAhead != nil {
			if nlaFound && !plaFound {
				continue
			}
		}
		if r.NegativeLookBehind == nil && r.NegativeLookAhead == nil && r.PositiveLookAhead != nil {
			if !plaFound {
				continue
			}
		}

		matches = append(matches, Match{start: f[0], end: f[1]})
	}

	return matches
}
