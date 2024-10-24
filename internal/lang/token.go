package lang

type RegionToken struct {
	Region *Region
	Start  bool // Start (true) or end (false) token
	// Token start index for start token or token end index
	// for end token.
	StartIdx int
	EndIdx   int
}

func (rt RegionToken) Overlaps(rt2 RegionToken) bool {
	return (rt.StartIdx <= rt2.StartIdx && rt2.StartIdx < rt.EndIdx) ||
		(rt.StartIdx < rt2.EndIdx && rt2.EndIdx <= rt.EndIdx)
}
