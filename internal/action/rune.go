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
