package enix

import (
	"unicode/utf8"

	"github.com/gdamore/tcell/v2"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/view"
)

type menuItem struct {
	name     string
	startIdx int
	endIdx   int
}

type menu struct {
	frame *frame.Frame

	items       []menuItem
	currItemIdx int // Current item index

	line *line.Line
	view view.View

	style         tcell.Style
	currItemStyle tcell.Style
}

func newMenu(
	frame *frame.Frame,
	itemNames []string,
	currItemIdx int,
	style tcell.Style,
	currItemStyle tcell.Style,
) *menu {
	// Create line and item list
	items := make([]menuItem, len(itemNames))
	line, _ := line.FromString("")
	rIdx := 0
	for x, name := range itemNames {
		it := menuItem{
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

	width, _ := Screen.Size()
	view := view.View{
		Line:   1,
		Column: 1,
		Height: 1,
		Width:  width - 4,
	}

	return &menu{
		frame:         frame,
		items:         items,
		currItemIdx:   currItemIdx,
		line:          line,
		view:          view,
		style:         style,
		currItemStyle: currItemStyle,
	}
}

func (menu *menu) CurrentItemIdx() int {
	return menu.currItemIdx
}

func (menu *menu) CurrentItemName() string {
	return menu.items[menu.currItemIdx].name
}

func (menu *menu) updateView() {
	// Screen might be resized, update view width
	// Assume a menu always occupies full width.
	width, _ := Screen.Size()
	menu.view.Width = width - 4

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
func (menu *menu) Next() (int, string) {
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
func (menu *menu) Prev() (int, string) {
	menu.currItemIdx--
	if menu.currItemIdx < 0 {
		menu.currItemIdx = len(menu.items) - 1
	}

	menu.updateView()

	idx := menu.currItemIdx
	return idx, menu.items[idx].name
}

// RxMouseEvent handles mouse event.
//
// Return values are:
//   - -1, "" in case of events resulting with menu scroll
//     or clicks outside menu item frames,
//   - index and name of the item on which the event occurred in all other cases.
//
// Menu handles only mouse scroll and primary click events internally.
// All other events, for exmpale scroll click, must be handled by higher
// abstraction layers. In such a case, call this function to get index and
// name of the item on which the event occurred.
func (menu *menu) RxMouseEvent(ev mouse.Event) (int, string) {
	frame := menu.frame
	lFrame := frame.ColumnSubframe(frame.X, 2)
	iFrame := frame.ColumnSubframe(frame.X+2, frame.Width-4)
	rFrame := frame.ColumnSubframe(frame.LastX()-1, 2)

	idx := 0

	switch ev.(type) {
	case mouse.PrimaryClick, mouse.DoublePrimaryClick, mouse.TriplePrimaryClick:
		if lFrame.Within(ev.X(), ev.Y()) {
			menu.viewLeft()
			return -1, ""
		} else if rFrame.Within(ev.X(), ev.Y()) {
			menu.viewRight()
			return -1, ""
		} else {
			idx = menu.clickItemsFrame(ev.X() - iFrame.X)
			if idx >= 0 {
				menu.currItemIdx = idx
			}
		}
	case mouse.MiddleClick:
		idx = menu.clickItemsFrame(ev.X() - iFrame.X)
	case mouse.WheelDown:
		menu.viewRight()
		return -1, ""
	case mouse.WheelUp:
		menu.viewLeft()
		return -1, ""
	}

	name := ""
	if idx >= 0 {
		name = menu.items[idx].name
	}

	return idx, name
}

func (menu *menu) clickItemsFrame(x int) int {
	rIdx, _, ok := menu.line.RuneIdx(menu.view.Column + x)
	if !ok {
		return -1
	}

	for i, item := range menu.items {
		if item.startIdx <= rIdx && rIdx < item.endIdx {
			return i
		}
	}

	return -1
}

func (menu *menu) viewLeft() {
	for range 2 {
		menu.view = menu.view.Left()
	}
}

func (menu *menu) viewRight() {
	lineCols := menu.line.Columns()
	for range 2 {
		if menu.view.LastColumn() >= lineCols {
			return
		}
		menu.view = menu.view.Right()
	}
}

func (menu *menu) Render(frame frame.Frame) {
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

	iFrame := frame.ColumnSubframe(frame.X+2, frame.Width-4)
	line.Render(1, iFrame, menu.view, hls, nil)

	// Fill missing space
	for x := line.Columns(); x < iFrame.Width; x++ {
		iFrame.SetContent(x, 0, ' ', menu.style)
	}

	lFrame := frame.ColumnSubframe(frame.X, 2)
	menu.renderLeftArrow(lFrame)

	rFrame := frame.ColumnSubframe(frame.LastX()-1, 2)
	menu.renderRightArrow(rFrame)
}

func (menu *menu) renderLeftArrow(frame frame.Frame) {
	r := ' '
	if menu.view.Column > 1 {
		r = '<'
	}

	frame.SetContent(0, 0, r, menu.style)
	frame.SetContent(1, 0, ' ', menu.style)
}

func (menu *menu) renderRightArrow(frame frame.Frame) {
	r := ' '
	if menu.view.LastColumn() < menu.line.Columns() {
		r = '>'
	}

	frame.SetContent(0, 0, ' ', menu.style)
	frame.SetContent(1, 0, r, menu.style)
}
