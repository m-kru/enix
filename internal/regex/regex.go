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
	finds := r.Regex.FindAllIndex(buf, -1)

	var negLookBeh [][]int
	var negLookAhead [][]int
	var posLookAhead [][]int
	var posLookBeh [][]int
	hasLookarounds := false

	if len(finds) > 0 {
		if r.NegativeLookBehind != nil {
			negLookBeh = r.NegativeLookBehind.FindAllIndex(buf, -1)
			hasLookarounds = true
		}
		if r.NegativeLookAhead != nil {
			negLookAhead = r.NegativeLookAhead.FindAllIndex(buf, -1)
			hasLookarounds = true
		}
		if r.PositiveLookAhead != nil {
			posLookAhead = r.PositiveLookAhead.FindAllIndex(buf, -1)
			hasLookarounds = true
		}
		if r.PositiveLookBehind != nil {
			posLookBeh = r.PositiveLookBehind.FindAllIndex(buf, -1)
			hasLookarounds = true
		}
	}

	if !hasLookarounds {
		matches := make([]Match, len(finds), len(finds))
		for i, f := range finds {
			matches[i].Start = f[0]
			matches[i].End = f[1]
		}
		return matches
	}

	matches := make([]Match, 0, len(finds))

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
