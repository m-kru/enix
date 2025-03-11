package tabbar

import (
	"fmt"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/m-kru/enix/internal/frame"
	ln "github.com/m-kru/enix/internal/line"
	"github.com/m-kru/enix/internal/tab"
	vw "github.com/m-kru/enix/internal/view"
)

type tabBarItem struct {
	Tab      *tab.Tab
	Name     string
	StartIdx int // Start rune idx
	EndIdx   int // End rune idx
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

	item.Name = name
}

type TabBar struct {
	items []*tabBarItem

	line   *ln.Line
	lFrame frame.Frame // Left arrow frame
	iFrame frame.Frame // Items frame
	rFrame frame.Frame // Right arrow frame
	view   vw.View
}

func (tb *TabBar) SetFrame(f frame.Frame) {
	tb.lFrame = f.ColumnSubframe(f.X, 2)
	tb.iFrame = f.ColumnSubframe(f.X+2, f.Width-4)
	tb.rFrame = f.ColumnSubframe(f.LastX()-1, 2)

	// Init view
	if tb.view.Line == 0 {
		tb.view = vw.View{
			Line:   1,
			Column: 1,
			Height: 1,
			Width:  tb.iFrame.Width,
		}
	}
}

func (tb *TabBar) Update(tabs *tab.Tab, currentTab *tab.Tab) {
	tb.createItems(tabs)

	// Create line
	line, _ := ln.FromString("")

	rIdx := 0

	for x := range tb.items {
		t := tb.items[x].Tab

		tb.items[x].StartIdx = rIdx

		line.Append([]byte(" "))
		rIdx++

		if t.HasChanges() {
			line.Append([]byte("*"))
			rIdx++
		}

		name := tb.items[x].Name
		line.Append([]byte(name))
		rIdx += utf8.RuneCountInString(name)

		line.Append([]byte(" "))
		rIdx++

		tb.items[x].EndIdx = rIdx
	}

	tb.line = line
}

func (tb *TabBar) viewLeft() {
	for range 2 {
		tb.view = tb.view.Left()
	}
}

func (tb *TabBar) viewRight() {
	lineCols := tb.line.Columns()
	for range 2 {
		if tb.view.LastColumn() >= lineCols {
			return
		}
		tb.view = tb.view.Right()
	}
}
