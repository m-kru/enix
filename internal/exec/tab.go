package exec

import (
	"fmt"
	"regexp"

	"github.com/m-kru/enix/internal/tab"
)

func Tab(args []string, firstTab *tab.Tab) (*tab.Tab, error) {
	if len(args) != 1 {
		return firstTab, fmt.Errorf("tab: expected 1 arg, provided %d", len(args))
	}

	regex := args[0]
	re, err := regexp.Compile(regex)
	if err != nil {
		return firstTab, fmt.Errorf("can't compile regex: %v", err)
	}

	t := firstTab
	newT := t
	foundCnt := 0
	for t != nil {
		if re.MatchString(t.Path) {
			newT = t
			foundCnt++
		}
		t = t.Next
	}

	if foundCnt != 1 {
		return firstTab, fmt.Errorf(
			"found %d tabs matching regex '%s'", foundCnt, regex,
		)
	}

	return newT, nil
}

func TabNext(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) != 0 {
		return t, fmt.Errorf("tab-next: expected 0 args, provided %d", len(args))
	}

	if t.Next != nil {
		t = t.Next
	} else {
		t = t.First()
	}

	return t, nil
}

func TabPrev(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) != 0 {
		return t, fmt.Errorf("tab-prev: expected 0 args, provided %d", len(args))
	}

	if t.Prev != nil {
		t = t.Prev
	} else {
		t = t.Last()
	}

	return t, nil
}
