package action

import "github.com/m-kru/enix/internal/line"

type (
	NewlineDelete struct {
		Line *line.Line // Pointer to the deleted line
	}

	NewlineInsert struct {
		Line    *line.Line // Newly inserted line
		RuneIdx int
	}

	RuneDelete struct {
		Line    *line.Line
		RuneIdx int
	}

	RuneInsert struct {
		Line    *line.Line
		RuneIdx int
	}
)

func (nd *NewlineDelete) isAction() {}
func (ni *NewlineInsert) isAction() {}
func (rd *RuneDelete) isAction()    {}
func (ri *RuneInsert) isAction()    {}
