package menu

import (
	"unicode/utf8"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/view"

	"github.com/gdamore/tcell/v2"
)

type item struct {
	name     string
	startIdx int
	endIdx   int
}

type Menu struct {
	lFrame frame.Frame // Left arrow frame
	iFrame frame.Frame // Items frame
	rFrame frame.Frame // Right arrow frame

	items       []item
	currItemIdx int // Current item index

	line *line.Line
	view view.View

	style         tcell.Style
	currItemStyle tcell.Style // Current item style
}

func New(
	frame frame.Frame,
	itemNames []string,
	currItemIdx int,
	style tcell.Style,
	currItemStyle tcell.Style,
) *Menu {
	// Create line and item list
	items := make([]item, len(itemNames))
	line, _ := line.FromString("")
	rIdx := 0
	for x, name := range itemNames {
		it := item{
			name:     name,
			startIdx: rIdx,
			endIdx:   0,
		}

		line.Append([]byte(" "))
		rIdx++

		line.Append([]byte(name))
		rIdx += utf8.RuneCountInString(name)

		line.Append([]byte(" "))
		rIdx++

		it.endIdx = rIdx

		items[x] = it
	}

	lFrame := frame.ColumnSubframe(frame.X, 2)
	iFrame := frame.ColumnSubframe(frame.X+2, frame.Width-4)
	rFrame := frame.ColumnSubframe(frame.LastX()-1, 2)

	view := view.View{
		Line:   1,
		Column: 1,
		Height: 1,
		Width:  iFrame.Width,
	}

	return &Menu{
		lFrame:        lFrame,
		iFrame:        iFrame,
		rFrame:        rFrame,
		items:         items,
		currItemIdx:   currItemIdx,
		line:          line,
		view:          view,
		style:         style,
		currItemStyle: currItemStyle,
	}
}

func (menu *Menu) CurrentItemIdx() int {
	return menu.currItemIdx
}

func (menu *Menu) CurrentItemName() string {
	return menu.items[menu.currItemIdx].name
}

func (menu *Menu) updateView() {
	item := menu.items[menu.currItemIdx]
	sIdx := item.startIdx
	eIdx := item.endIdx
	sCol := menu.line.ColumnIdx(sIdx)
	eCol := menu.line.ColumnIdx(eIdx)

	if menu.view.IsColumnVisible(sCol) {
		if menu.view.IsColumnVisible(eCol) {
			return
		}
		for range eCol - menu.view.LastColumn() - 1 {
			menu.view = menu.view.Right()
		}
	} else {
		if sCol < menu.view.Column {
			for range menu.view.Column - sCol {
				menu.view = menu.view.Left()
			}
		} else {
			for range eCol - menu.view.LastColumn() - 1 {
				menu.view = menu.view.Right()
			}
		}
	}
}

// Next goes to the next item.
// If current item is the last one, then it wraps to the first item.
func (menu *Menu) Next() (int, string) {
	menu.currItemIdx++
	if menu.currItemIdx == len(menu.items) {
		menu.currItemIdx = 0
	}

	menu.updateView()

	idx := menu.currItemIdx
	return idx, menu.items[idx].name
}

// Prev goes to the previous item.
// If current item is the first one, then it wraps to the last item.
func (menu *Menu) Prev() (int, string) {
	menu.currItemIdx--
	if menu.currItemIdx < 0 {
		menu.currItemIdx = len(menu.items) - 1
	}

	menu.updateView()

	idx := menu.currItemIdx
	return idx, menu.items[idx].name
}

// RxMouseEvent handles mouse event.
// Returned values are the current item index and name.
func (menu *Menu) RxMouseEvent(ev mouse.Event) (int, string) {
	switch ev.(type) {
	case mouse.PrimaryClick, mouse.DoublePrimaryClick, mouse.TriplePrimaryClick:
		if menu.lFrame.Within(ev.X(), ev.Y()) {
			menu.viewLeft()
		} else if menu.rFrame.Within(ev.X(), ev.Y()) {
			menu.viewRight()
		} else {
			menu.clickItemsFrame(ev.X() - menu.iFrame.X)
		}
	case mouse.WheelDown:
		menu.viewRight()
	case mouse.WheelUp:
		menu.viewLeft()
	}

	idx := menu.currItemIdx
	return idx, menu.items[idx].name
}

func (menu *Menu) clickItemsFrame(x int) {
	rIdx, _, ok := menu.line.RuneIdx(menu.view.Column + x)
	if !ok {
		return
	}

	for i, item := range menu.items {
		if item.startIdx <= rIdx && rIdx < item.endIdx {
			menu.currItemIdx = i
			return
		}
	}
}

func (menu *Menu) viewLeft() {
	for range 2 {
		menu.view = menu.view.Left()
	}
}

func (menu *Menu) viewRight() {
	lineCols := menu.line.Columns()
	for range 2 {
		if menu.view.LastColumn() >= lineCols {
			return
		}
		menu.view = menu.view.Right()
	}
}

func (menu *Menu) Render() {
	currItem := menu.items[menu.currItemIdx]
	line := menu.line

	hls := []highlight.Highlight{
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: 0,
			EndRuneIdx:   currItem.startIdx,
			Style:        menu.style,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currItem.startIdx,
			EndRuneIdx:   currItem.endIdx,
			Style:        menu.currItemStyle,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currItem.endIdx,
			EndRuneIdx:   line.RuneCount(),
			Style:        menu.style,
		},
	}

	iFrame := menu.iFrame
	line.Render(1, iFrame, menu.view, hls, nil)

	// Fill missing space
	for x := line.Columns(); x < iFrame.Width; x++ {
		iFrame.SetContent(x, 0, ' ', menu.style)
	}

	menu.renderLeftArrow()
	menu.renderRightArrow()
}

func (menu *Menu) renderLeftArrow() {
	r := ' '
	if menu.view.Column > 1 {
		r = '<'
	}
	menu.lFrame.SetContent(0, 0, r, menu.style)
	menu.lFrame.SetContent(1, 0, ' ', menu.style)
}

func (menu *Menu) renderRightArrow() {
	menu.rFrame.SetContent(0, 0, ' ', menu.style)
	r := ' '
	if menu.view.LastColumn() < menu.line.Columns() {
		r = '>'
	}
	menu.rFrame.SetContent(1, 0, r, menu.style)
}
