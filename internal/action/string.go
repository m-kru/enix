package action

import "github.com/m-kru/enix/internal/line"

type (
	StringDelete struct {
		Line         *line.Line
		Str          string
		StartRuneIdx int
		RuneCount    int
	}

	StringInsert struct {
		Line         *line.Line
		Str          string
		StartRuneIdx int
		RuneCount    int
	}
)

func (sd *StringDelete) isAction() {}
func (si *StringInsert) isAction() {}

func (sd *StringDelete) Reverse() Action {
	return &StringInsert{
		Line:         sd.Line,
		Str:          sd.Str,
		StartRuneIdx: sd.StartRuneIdx,
		RuneCount:    sd.RuneCount,
	}
}

func (si *StringInsert) Reverse() Action {
	return &StringDelete{
		Line:         si.Line,
		Str:          si.Str,
		StartRuneIdx: si.StartRuneIdx,
		RuneCount:    si.RuneCount,
	}
}
