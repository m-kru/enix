package tab

import (
	"github.com/m-kru/enix/internal/cursor"
)

// PrimaryClick handles mouse primary button click.
// x and y are tab frame coordinates.
func (tab *Tab) PrimaryClick(x, y int) {
	x -= tab.LineNumWidth() + 1
	if x < 0 {
		x = 0
	}

	line := tab.Lines.Get(tab.View.Line + y)
	if line == nil {
		line = tab.Lines.Last()
	}

	cfg := tab.Config
	idx, _, ok := line.RuneIdx(tab.View.Column+x, cfg.TabWidth)
	if !ok {
		idx = line.Len()
	}

	c := cursor.Cursor{
		Config: cfg,
		Line:   line,
		Idx:    idx,
		BufIdx: idx,
	}

	tab.Cursors = &c
}

// PrimaryClickCtrl handles mouse primary button click with Ctrl modifier.
// x and y are tab frame coordinates.
func (tab *Tab) PrimaryClickCtrl(x, y int) {
	x -= tab.LineNumWidth() + 1
	if x < 0 {
		x = 0
	}

	line := tab.Lines.Get(tab.View.Line + y)
	if line == nil {
		line = tab.Lines.Last()
	}

	idx, _, ok := line.RuneIdx(tab.View.Column+x, tab.Config.TabWidth)
	if !ok {
		idx = line.Len()
	}

	tab.AddCursor(
		line.LineNum(), line.ColumnIdx(idx, tab.Config.TabWidth),
	)
}
