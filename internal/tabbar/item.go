package tabbar

import (
	"github.com/m-kru/enix/internal/tab"
)

func (tb *TabBar) getCurrentItem(currentTab *tab.Tab) *tabBarItem {
	for x := range tb.items {
		if tb.items[x].Tab == currentTab {
			return tb.items[x]
		}
	}

	// If the code gets here, then there is a bug.
	// However, to avoid panics return first item.
	return tb.items[0]
}

func (tb *TabBar) createItems(tabs *tab.Tab) {
	tb.items = make([]*tabBarItem, 0, tabs.Count())

	t := tabs
	for {
		if t == nil {
			break
		}

		i := tabBarItem{t, "", 0, 0}
		tb.items = append(tb.items, &i)

		t = t.Next
	}

	tb.assignItemNames()
}

func (tb *TabBar) assignItemNames() {
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

func (tb *TabBar) itemsNameConflicts() map[string][]int {
	names := make(map[string][]int)

	for x, i := range tb.items {
		names[i.Name] = append(names[i.Name], x)
	}

	for key, val := range names {
		if len(val) == 1 {
			delete(names, key)
		}
	}

	return names
}

func (tb *TabBar) extendItemsNames(confs map[string][]int, lvl int) {
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
