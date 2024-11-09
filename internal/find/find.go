package find

type Find struct {
	LineNum      int
	StartRuneIdx int
	EndRuneIdx   int
}

func (f Find) CoversCell(lineNum int, idx int) bool {
	return lineNum == f.LineNum && f.StartRuneIdx <= idx && idx < f.EndRuneIdx
}

func (f Find) IsLastCell(lineNum int, idx int) bool {
	return lineNum == f.LineNum && idx == f.EndRuneIdx-1
}
