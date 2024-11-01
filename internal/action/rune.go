package action

import "github.com/m-kru/enix/internal/line"

type (
	NewlineDelete struct {
		Line     *line.Line
		PrevLine *line.Line
	}

	NewlineInsert struct {
		Line    *line.Line
		Idx     int
		NewLine *line.Line
	}

	RuneDelete struct {
		Line *line.Line
		Idx  int
	}

	RuneInsert struct {
		Line *line.Line
		Idx  int
	}
)

func (nd *NewlineDelete) isAction() {}
func (ni *NewlineInsert) isAction() {}
func (rd *RuneDelete) isAction()    {}
func (ri *RuneInsert) isAction()    {}
