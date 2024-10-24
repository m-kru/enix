package highlight

import "github.com/gdamore/tcell/v2"

type Highlight struct {
	Line     int
	StartIdx int
	EndIdx   int
	Style    tcell.Style
}

func (hl Highlight) CoversCell(lineNum int, idx int) bool {
	return lineNum == hl.Line && hl.StartIdx <= idx && idx <= hl.EndIdx
}

func (hl Highlight) Contains(hl2 Highlight) bool {
	return hl.Line == hl2.Line &&
		hl.StartIdx <= hl2.StartIdx && hl2.StartIdx <= hl.EndIdx &&
		hl.StartIdx <= hl2.EndIdx && hl2.EndIdx <= hl.EndIdx
}

// Split divides hl into 1,2 or 3 new highlights.
// It is users responsibility to make sure hl contains hl2
// by calling the Contains function before splitting.
func (hl Highlight) Split(hl2 Highlight) []Highlight {
	hls := make([]Highlight, 0, 3)

	// Check if hl2 exactly covers hl.
	if hl.StartIdx == hl2.StartIdx && hl.EndIdx == hl2.EndIdx {
		hls = append(hls, hl2)
		return hls
	}

	if hl.StartIdx < hl2.StartIdx {
		hl1 := hl
		hl1.EndIdx = hl2.StartIdx - 1
		hls = append(hls, hl1)
	}

	hls = append(hls, hl2)

	if hl.EndIdx > hl2.EndIdx {
		hl3 := hl
		hl3.StartIdx = hl2.EndIdx + 1
		hls = append(hls, hl3)
	}

	return hls
}
