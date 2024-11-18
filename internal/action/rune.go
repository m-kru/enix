package action

import "github.com/m-kru/enix/internal/line"

type (
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

func (rd *RuneDelete) isAction() {}
func (ri *RuneInsert) isAction() {}

func (rd *RuneDelete) Reverse() Action {
	return &RuneInsert{Line: rd.Line, Rune: rd.Rune, RuneIdx: rd.RuneIdx}
}

func (ri *RuneInsert) Reverse() Action {
	return &RuneDelete{Line: ri.Line, Rune: ri.Rune, RuneIdx: ri.RuneIdx}
}
