package highlight

import "github.com/gdamore/tcell/v2"

type Highlight struct {
	LineNum      int
	StartRuneIdx int // Inclusive
	EndRuneIdx   int // Exclusive
	Style        tcell.Style
}

func (hl Highlight) CoversCell(lineNum int, idx int) bool {
	return lineNum == hl.LineNum && hl.StartRuneIdx <= idx && idx < hl.EndRuneIdx
}

func (hl Highlight) Contains(hl2 Highlight) bool {
	return hl.LineNum == hl2.LineNum &&
		hl.StartRuneIdx <= hl2.StartRuneIdx && hl2.StartRuneIdx < hl.EndRuneIdx &&
		hl.StartRuneIdx < hl2.EndRuneIdx && hl2.EndRuneIdx <= hl.EndRuneIdx
}

// Split divides hl into 1,2 or 3 new highlights.
// It is users responsibility to make sure hl contains hl2
// by calling the Contains function before splitting.
func (hl Highlight) Split(hl2 Highlight) []Highlight {
	hls := make([]Highlight, 0, 3)

	// Check if hl2 exactly covers hl.
	if hl.StartRuneIdx == hl2.StartRuneIdx && hl.EndRuneIdx == hl2.EndRuneIdx {
		hls = append(hls, hl2)
		return hls
	}

	if hl.StartRuneIdx < hl2.StartRuneIdx {
		hl1 := hl
		hl1.EndRuneIdx = hl2.StartRuneIdx
		hls = append(hls, hl1)
	}

	hls = append(hls, hl2)

	if hl.EndRuneIdx > hl2.EndRuneIdx {
		hl3 := hl
		hl3.StartRuneIdx = hl2.EndRuneIdx
		hls = append(hls, hl3)
	}

	return hls
}
