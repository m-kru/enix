package tabbar

import (
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/tab"
	vw "github.com/m-kru/enix/internal/view"
)

var lFrame frame.Frame // Left arrow frame
var iFrame frame.Frame // Items frame
var rFrame frame.Frame // Right arrow frame

var view vw.View

func SetFrame(f frame.Frame) {
	lFrame = f.ColumnSubframe(f.X, 2)
	iFrame = f.ColumnSubframe(f.X+2, f.Width-4)
	rFrame = f.ColumnSubframe(f.LastX()-1, 2)
}

func Update(tabs *tab.Tab, currentTab *tab.Tab) {
	items = createItems(tabs)
}
