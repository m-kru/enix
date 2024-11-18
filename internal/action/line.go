package action

import "github.com/m-kru/enix/internal/line"

type (
	LineDown struct {
		Line *line.Line // Pointer to the line moved down
	}

	LineUp struct {
		Line *line.Line // Pointer to the line moved up
	}

	NewlineDelete struct {
		Line1    *line.Line
		Line1Num int
		RuneIdx  int // Equals Line1.RuneCount() before delete
		Line2    *line.Line
		NewLine  *line.Line
	}

	NewlineInsert struct {
		Line     *line.Line
		LineNum  int
		RuneIdx  int
		NewLine1 *line.Line
		NewLine2 *line.Line
	}
)

func (ld *LineDown) isAction()      {}
func (lu *LineUp) isAction()        {}
func (nd *NewlineDelete) isAction() {}
func (ni *NewlineInsert) isAction() {}

func (ld *LineDown) Reverse() Action {
	return &LineUp{Line: ld.Line}
}

func (lu *LineUp) Reverse() Action {
	return &LineDown{Line: lu.Line}
}

func (nd *NewlineDelete) Reverse() Action {
	return &NewlineInsert{
		Line:     nd.NewLine,
		LineNum:  nd.Line1Num,
		RuneIdx:  nd.RuneIdx,
		NewLine1: nd.Line1,
		NewLine2: nd.Line2,
	}
}

func (ni *NewlineInsert) Reverse() Action {
	return &NewlineDelete{
		Line1:    ni.NewLine1,
		Line1Num: ni.LineNum,
		RuneIdx:  ni.RuneIdx,
		Line2:    ni.NewLine2,
		NewLine:  ni.Line,
	}
}
