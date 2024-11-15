package action

import "github.com/m-kru/enix/internal/line"

type (
	NewlineDelete struct {
		Line    *line.Line
		LineNum int
		RuneIdx int // Equals Line.RuneCount() before delete
	}

	NewlineInsert struct {
		Line    *line.Line
		LineNum int
		RuneIdx int
	}

	RuneDelete struct {
		Line    *line.Line
		Rune    rune
		RuneIdx int
	}

	RuneInsert struct {
		Line    *line.Line
		Rune    rune
		RuneIdx int
	}
)

func (nd *NewlineDelete) isAction() {}
func (ni *NewlineInsert) isAction() {}
func (rd *RuneDelete) isAction()    {}
func (ri *RuneInsert) isAction()    {}

func (nd *NewlineDelete) Reverse() Action {
	return &NewlineInsert{Line: nd.Line, LineNum: nd.LineNum, RuneIdx: nd.RuneIdx}
}

func (ni *NewlineInsert) Reverse() Action {
	return &NewlineDelete{Line: ni.Line, LineNum: ni.LineNum, RuneIdx: ni.RuneIdx}
}

func (rd *RuneDelete) Reverse() Action {
	return &RuneInsert{Line: rd.Line, Rune: rd.Rune, RuneIdx: rd.RuneIdx}
}

func (ri *RuneInsert) Reverse() Action {
	return &RuneDelete{Line: ri.Line, Rune: ri.Rune, RuneIdx: ri.RuneIdx}
}
