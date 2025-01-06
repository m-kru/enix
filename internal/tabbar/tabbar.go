package tabbar

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/frame"
	ln "github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/tab"
	vw "github.com/m-kru/enix/internal/view"
)

var line *ln.Line

var lFrame frame.Frame // Left arrow frame
var iFrame frame.Frame // Items frame
var rFrame frame.Frame // Right arrow frame

var view vw.View

func SetFrame(f frame.Frame) {
	lFrame = f.ColumnSubframe(f.X, 2)
	iFrame = f.ColumnSubframe(f.X+2, f.Width-4)
	rFrame = f.ColumnSubframe(f.LastX()-1, 2)

	// Init view
	if view.Line == 0 {
		view = vw.View{
			Line:   1,
			Column: 1,
			Height: 1,
			Width:  iFrame.Width,
		}
	}
}

func Update(tabs *tab.Tab, currentTab *tab.Tab) {
	items = createItems(tabs)

	// Create line
	line, _ = ln.FromString("")

	rIdx := 0

	for x := range items {
		t := items[x].Tab

		items[x].StartIdx = rIdx

		line.Append([]byte(" "))
		rIdx++

		if t.HasChanges() {
			line.Append([]byte("*"))
			rIdx++
		}

		name := items[x].Name
		line.Append([]byte(name))
		rIdx += utf8.RuneCountInString(name)

		line.Append([]byte(" "))
		rIdx++

		items[x].EndIdx = rIdx
	}
}

func viewLeft() {
	for range 2 {
		view = view.Left()
	}
}

func viewRight() {
	lineCols := line.Columns()
	for range 2 {
		if view.LastColumn() >= lineCols {
			return
		}
		view = view.Right()
	}
}
