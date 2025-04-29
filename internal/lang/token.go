package lang

import (
	"sort"
)

type RegionToken struct {
	region      *Region
	startBufIdx int
	endBufIdx   int
}

func (rt RegionToken) Overlaps(rt2 RegionToken) bool {
	return (rt.startBufIdx <= rt2.startBufIdx && rt2.startBufIdx < rt.endBufIdx) ||
		(rt.startBufIdx < rt2.endBufIdx && rt2.endBufIdx <= rt.endBufIdx)
}

func lineStartTokens(line []byte, byteOffset int, regions []*Region, toks *[]RegionToken) {
	var tok RegionToken

	// Skip the default region
	for _, r := range regions[1:] {
		tok.region = r

		finds := r.Start.FindAll(line[byteOffset:])
		for _, f := range finds {
			tok.startBufIdx = byteOffset + f.Start
			tok.endBufIdx = byteOffset + f.End
			*toks = append(*toks, tok)
		}
	}

	less := func(i, j int) bool {
		ti := (*toks)[i]
		tj := (*toks)[j]

		if ti.startBufIdx < tj.startBufIdx {
			return true
		} else if ti.startBufIdx == tj.startBufIdx {
			// Longer start tokens take precedence over shorter ones.
			return ti.endBufIdx >= tj.endBufIdx
		}

		return false
	}
	sort.Slice(*toks, less)
}

func lineEndTokens(line []byte, byteOffset int, region *Region) []RegionToken {
	toks := []RegionToken{}
	var tok RegionToken

	tok.region = region

	finds := region.End.FindAll(line[byteOffset:])
	for _, f := range finds {
		tok.startBufIdx = byteOffset + f.Start
		tok.endBufIdx = byteOffset + f.End
		toks = append(toks, tok)
	}

	return toks
}
