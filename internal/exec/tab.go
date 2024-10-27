package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

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
