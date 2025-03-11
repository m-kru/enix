package enix

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/highlight"
	"github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"
	"github.com/m-kru/enix/internal/view"
)

type tabBarItem struct {
	menuItem
	Tab *tab.Tab
}

func (item *tabBarItem) assignName(lvl int) {
	fields := strings.Split(item.Tab.Path, fmt.Sprintf("%c", os.PathSeparator))
	name := ""

	if lvl > 0 {
		for x := range lvl {
			if x >= len(fields)-1 {
				break
			}

			name = fmt.Sprintf(
				"%s%c%s",
				fields[len(fields)-2-x], os.PathSeparator, name,
			)
		}
	}

	name += fields[(len(fields) - 1)]

	item.name = name
}

type tabBar struct {
	items []*tabBarItem

	line   *line.Line
	lFrame frame.Frame // Left arrow frame
	iFrame frame.Frame // Items frame
	rFrame frame.Frame // Right arrow frame
	view   view.View
}

func (tb *tabBar) SetFrame(f frame.Frame) {
	tb.lFrame = f.ColumnSubframe(f.X, 2)
	tb.iFrame = f.ColumnSubframe(f.X+2, f.Width-4)
	tb.rFrame = f.ColumnSubframe(f.LastX()-1, 2)

	// Init view
	if tb.view.Line == 0 {
		tb.view = view.View{
			Line:   1,
			Column: 1,
			Height: 1,
			Width:  tb.iFrame.Width,
		}
	}
}

func (tb *tabBar) Update(tabs *tab.Tab, currentTab *tab.Tab) {
	tb.createItems(tabs)

	// Create line
	line, _ := line.FromString("")

	rIdx := 0

	for x := range tb.items {
		t := tb.items[x].Tab

		tb.items[x].startIdx = rIdx

		line.Append([]byte(" "))
		rIdx++

		if t.HasChanges() {
			line.Append([]byte("*"))
			rIdx++
		}

		name := tb.items[x].name
		line.Append([]byte(name))
		rIdx += utf8.RuneCountInString(name)

		line.Append([]byte(" "))
		rIdx++

		tb.items[x].endIdx = rIdx
	}

	tb.line = line
}

func (tb *tabBar) viewLeft() {
	for range 2 {
		tb.view = tb.view.Left()
	}
}

func (tb *tabBar) viewRight() {
	lineCols := tb.line.Columns()
	for range 2 {
		if tb.view.LastColumn() >= lineCols {
			return
		}
		tb.view = tb.view.Right()
	}
}

func (tb *tabBar) getCurrentItem(currentTab *tab.Tab) *tabBarItem {
	for x := range tb.items {
		if tb.items[x].Tab == currentTab {
			return tb.items[x]
		}
	}

	// If the code gets here, then there is a bug.
	// However, to avoid panics return first item.
	return tb.items[0]
}

func (tb *tabBar) createItems(tabs *tab.Tab) {
	tb.items = make([]*tabBarItem, 0, tabs.Count())

	t := tabs
	for {
		if t == nil {
			break
		}

		i := tabBarItem{
			menuItem: menuItem{"", 0, 0},
			Tab:      t,
		}
		tb.items = append(tb.items, &i)

		t = t.Next
	}

	tb.assignItemNames()
}

func (tb *tabBar) assignItemNames() {
	for x := range tb.items {
		tb.items[x].assignName(0)
	}

	for lvl := 1; ; lvl++ {
		confs := tb.itemsNameConflicts()
		if len(confs) == 0 {
			break
		}
		tb.extendItemsNames(confs, lvl)
	}
}

func (tb *tabBar) itemsNameConflicts() map[string][]int {
	names := make(map[string][]int)

	for x, i := range tb.items {
		names[i.name] = append(names[i.name], x)
	}

	for key, val := range names {
		if len(val) == 1 {
			delete(names, key)
		}
	}

	return names
}

func (tb *tabBar) extendItemsNames(confs map[string][]int, lvl int) {
	for _, idxs := range confs {
		its := []*tabBarItem{}
		for _, x := range idxs {
			its = append(its, tb.items[x])
		}
		for _, it := range its {
			it.assignName(lvl)
		}
	}
}

func (tb *tabBar) RxMouseEvent(ev mouse.Event) *tab.Tab {
	switch ev.(type) {
	case mouse.PrimaryClick, mouse.DoublePrimaryClick, mouse.TriplePrimaryClick:
		if tb.lFrame.Within(ev.X(), ev.Y()) {
			tb.viewLeft()
		} else if tb.rFrame.Within(ev.X(), ev.Y()) {
			tb.viewRight()
		} else {
			return tb.clickItemsFrame(ev.X() - tb.iFrame.X)
		}
	case mouse.WheelDown:
		tb.viewRight()
	case mouse.WheelUp:
		tb.viewLeft()
	}

	return nil
}

func (tb *tabBar) clickItemsFrame(x int) *tab.Tab {
	rIdx, _, ok := tb.line.RuneIdx(tb.view.Column + x)
	if !ok {
		return nil
	}

	for _, item := range tb.items {
		if item.startIdx <= rIdx && rIdx < item.endIdx {
			return item.Tab
		}
	}

	return nil
}

func (tb *tabBar) Render(currentTab *tab.Tab) {
	currentItem := tb.getCurrentItem(currentTab)

	hls := []highlight.Highlight{
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: 0,
			EndRuneIdx:   currentItem.startIdx,
			Style:        cfg.Style.TabBar,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currentItem.startIdx,
			EndRuneIdx:   currentItem.endIdx,
			Style:        cfg.Style.CurrentTab,
		},
		highlight.Highlight{
			LineNum:      1,
			StartRuneIdx: currentItem.endIdx,
			EndRuneIdx:   tb.line.RuneCount(),
			Style:        cfg.Style.TabBar,
		},
	}

	tb.line.Render(1, tb.iFrame, tb.view, hls, nil)

	// Fill missing space
	for x := tb.line.Columns(); x < tb.iFrame.Width; x++ {
		tb.iFrame.SetContent(x, 0, ' ', cfg.Style.TabBar)
	}

	tb.renderLeftArrow()
	tb.renderRightArrow()
}

func (tb *tabBar) renderLeftArrow() {
	r := ' '
	if tb.view.Column > 1 {
		r = '<'
	}
	tb.lFrame.SetContent(0, 0, r, cfg.Style.TabBar)
	tb.lFrame.SetContent(1, 0, ' ', cfg.Style.TabBar)
}

func (tb *tabBar) renderRightArrow() {
	tb.rFrame.SetContent(0, 0, ' ', cfg.Style.TabBar)
	r := ' '
	if tb.view.LastColumn() < tb.line.Columns() {
		r = '>'
	}
	tb.rFrame.SetContent(1, 0, r, cfg.Style.TabBar)
}
