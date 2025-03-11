package enix

import (
	"fmt"
	"os"
	"strings"

	"github.com/m-kru/enix/internal/frame"
	"github.com/m-kru/enix/internal/mouse"
	"github.com/m-kru/enix/internal/tab"
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
	items      []*tabBarItem
	menu       *menu
	updateView bool
}

func (tb *tabBar) Init() {
	tb.Update()
	tb.menu = newMenu([]string{""}, 0)
}

func (tb *tabBar) getCurrentItem() (*tabBarItem, int) {
	for idx, item := range tb.items {
		if item.Tab == CurrentTab {
			return item, idx
		}
	}

	// If the code gets here, then there is a bug.
	// However, to avoid panics return first item.
	return tb.items[0], 0
}

// Update rebuilts information on tab bar items.
// Must be called when tabs are created or removed.
func (tb *tabBar) Update() {
	tb.items = make([]*tabBarItem, 0, Tabs.Count())

	t := Tabs
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

	tb.updateView = true
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
	idx, _ := tb.menu.RxMouseEvent(ev)
	return tb.items[idx].Tab
}

func (tb *tabBar) Render(frame frame.Frame) {
	_, currIdx := tb.getCurrentItem()

	names := make([]string, 0, len(tb.items))

	for _, it := range tb.items {
		name := it.name
		if it.Tab.HasChanges() {
			name = "*" + name
		}
		names = append(names, name)
	}

	view := tb.menu.view
	tb.menu = newMenu(names, currIdx)
	tb.menu.view = view
	if tb.updateView {
		tb.menu.updateView()
		tb.updateView = false
	}

	tb.menu.Render(frame)
}
