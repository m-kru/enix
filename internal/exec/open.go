package exec

import (
	"fmt"
	"github.com/m-kru/enix/internal/tab"
)

func Open(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("open: expected at least 1 arg, provided 0")
	}

	var newCurrentTab *tab.Tab

	errMsg := ""
	for i, path := range args {
		newT, err := tab.Open(t.Keys, t.Frame, path)
		if newT != nil {
			t.Append(newT)
			if i == 0 {
				newCurrentTab = newT
			}
		}
		if err != nil {
			errMsg += err.Error() + "\n\n"
		}
	}

	if len(errMsg) > 0 {
		path := "enix-error"
		idx := 2
		for {
			if !t.Exists(path) {
				break
			}
			path = fmt.Sprintf("enix-error-%d", idx)
			idx++
		}

		errTab := tab.FromString(t.Keys, t.Frame, errMsg, path)
		t.Append(errTab)
		newCurrentTab = errTab
	}

	return newCurrentTab, nil
}
