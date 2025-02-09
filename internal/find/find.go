package find

type Find struct {
	LineNum      int
	StartRuneIdx int
	EndRuneIdx   int
}

func (f Find) CoversRune(lineNum int, rIdx int) bool {
	return lineNum == f.LineNum && f.StartRuneIdx <= rIdx && rIdx < f.EndRuneIdx
}

func (f Find) IsLastRune(rIdx int) bool {
	return rIdx == f.EndRuneIdx-1
}
