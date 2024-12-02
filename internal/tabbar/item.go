package tabbar

import (
	"fmt"
	"os"
	"strings"

	"github.com/m-kru/enix/internal/tab"
)

type item struct {
	Tab      *tab.Tab
	Name     string
	StartIdx int
	EndIdx   int
}

func createItems(tabs *tab.Tab) []item {
	items := make([]item, 0, tabs.Count())

	t := tabs
	for {
		if t == nil {
			break
		}

		i := item{t, "", 0, 0}
		items = append(items, i)

		t = t.Next
	}

	items = assignItemNames(items)

	return items
}

func assignItemNames(items []item) []item {
	for x := range items {
		assignItemName(&items[x], 0)
	}

	for lvl := 1; ; lvl++ {
		confs := itemsNameConflicts(items)
		if len(confs) == 0 {
			break
		}
		items = extendItemsNames(items, confs, lvl)
	}

	return items
}

func assignItemName(item *item, lvl int) {
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

func itemsNameConflicts(items []item) map[string][]int {
	names := make(map[string][]int)

	for x, i := range items {
		names[i.Name] = append(names[i.Name], x)
	}

	for key, val := range names {
		if len(val) == 1 {
			delete(names, key)
		}
	}

	return names
}

func extendItemsNames(items []item, confs map[string][]int, lvl int) []item {
	for _, idxs := range confs {
		its := []*item{}
		for _, x := range idxs {
			its = append(its, &items[x])
		}
		for _, it := range its {
			assignItemName(it, lvl)
		}
	}

	return items
}
