package tab

import (
	"github.com/m-kru/enix/internal/cursor"
	"github.com/m-kru/enix/internal/util"
)

// PrimaryClick handles mouse primary button click.
// x and y are tab frame coordinates.
func (t *Tab) PrimaryClick(x, y int) {
	lineNumWidth := util.IntWidth(t.LineCount())
	x -= lineNumWidth + 1
	if x < 0 {
		x = 0
	}

	line := t.Lines.Get(t.View.Line + y)
	if line == nil {
		line = t.Lines.Last()
	}

	cfg := t.Cursors.Config
	idx, _, ok := line.RuneIdx(t.View.Column+x, cfg.TabWidth)
	if !ok {
		idx = line.Len()
	}

	c := cursor.Cursor{
		Config: cfg,
		Line:   line,
		Idx:    idx,
		BufIdx: idx,
	}

	t.Cursors = &c
}
