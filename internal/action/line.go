package action

import "github.com/m-kru/enix/internal/line"

type (
	LineDown struct {
		Line *line.Line // Pointer to the line moved down
	}

	LineUp struct {
		Line *line.Line // Pointer to the line moved up
	}
)

func (ld *LineDown) isAction() {}
func (lu *LineUp) isAction()   {}
