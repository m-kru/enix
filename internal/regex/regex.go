package regex

import (
	"regexp"
)

type Regex struct {
	Regex              *regexp.Regexp
	NegativeLookBehind *regexp.Regexp
	NegativeLookAhead  *regexp.Regexp
	PositiveLookAhead  *regexp.Regexp
	PositiveLookBehind *regexp.Regexp
}

func (r Regex) FindAll(buf []byte) []Match {
	matches := make([]Match, 0, 8)
	var negLookBeh [][]int
	var negLookAhead [][]int
	var posLookAhead [][]int
	var posLookBeh [][]int

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
		if r.PositiveLookBehind != nil {
			posLookBeh = r.PositiveLookBehind.FindAllIndex(buf, -1)
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

		plbFound := false
		for _, plb := range posLookBeh {
			if plb[1] == f[0] {
				plbFound = true
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

		behindOk := true
		if (r.NegativeLookBehind != nil && nlbFound) || (r.PositiveLookBehind != nil && !plbFound) {
			behindOk = false
		}

		aheadOk := true
		if (r.NegativeLookAhead != nil && nlaFound) || (r.PositiveLookAhead != nil && !plaFound) {
			aheadOk = false
		}

		if behindOk && aheadOk {
			matches = append(matches, Match{Start: f[0], End: f[1]})
		}
	}

	return matches
}
