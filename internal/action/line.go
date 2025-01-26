package action

import "github.com/m-kru/enix/internal/line"

type (
	LineDelete struct {
		Line    *line.Line
		LineNum int        // Currently unused
		NewLine *line.Line // A line in the place of Line, not actually a new line.
	}

	LineInsert struct {
		Line    *line.Line
		LineNum int
	}

	LineDown struct {
		Line *line.Line // Pointer to the line moved down
	}

	LineUp struct {
		Line *line.Line // Pointer to the line moved up
	}

	NewlineDelete struct {
		Line         *line.Line
		LineNum      int
		RuneIdx      int // Equals Line.RuneCount() before delete
		TrimmedCount int // Number of runes trimmed from the Line2
		NewLine      *line.Line
	}

	NewlineInsert struct {
		Line     *line.Line
		LineNum  int
		RuneIdx  int
		NewLine1 *line.Line
		NewLine2 *line.Line
	}
)

func (ld *LineDelete) isAction()    {}
func (li *LineInsert) isAction()    {}
func (ld *LineDown) isAction()      {}
func (lu *LineUp) isAction()        {}
func (nd *NewlineDelete) isAction() {}
func (ni *NewlineInsert) isAction() {}

func (ld *LineDelete) Reverse() Action {
	return &LineInsert{Line: ld.Line, LineNum: ld.LineNum}
}

func (li *LineInsert) Reverse() Action {
	return &LineDelete{Line: li.Line, LineNum: li.LineNum, NewLine: nil}
}

func (ld *LineDown) Reverse() Action {
	return &LineUp{Line: ld.Line}
}

func (lu *LineUp) Reverse() Action {
	return &LineDown{Line: lu.Line}
}

func (nd *NewlineDelete) Reverse() Action {
	return &NewlineInsert{
		Line:     nd.NewLine,
		LineNum:  nd.LineNum,
		RuneIdx:  nd.RuneIdx,
		NewLine1: nd.Line,
		NewLine2: nd.Line.Next,
	}
}

func (ni *NewlineInsert) Reverse() Action {
	return &NewlineDelete{
		Line:         ni.NewLine1,
		LineNum:      ni.LineNum,
		RuneIdx:      ni.RuneIdx,
		TrimmedCount: 0,
		NewLine:      ni.Line,
	}
}
