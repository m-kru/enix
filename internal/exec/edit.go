package exec

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/m-kru/enix/internal/tab"
	"github.com/m-kru/enix/internal/util"
)

func Edit(args []string, t *tab.Tab) (*tab.Tab, error) {
	if len(args) == 0 {
		return nil, fmt.Errorf("open: expected at least 1 arg, provided 0")
	}

	var newCurrentTab *tab.Tab

	paths := args

	// Handle opening at specific line
	openLine, openCol, err := util.ParseLineAndColumnString(paths[len(paths)-1])
	if err == nil {
		paths = paths[0 : len(paths)-1]
	}

	errMsg := ""
	for i, path := range paths {
		// Check if tab with given path already exists
		abspath := path
		if !filepath.IsAbs(abspath) {
			wd, err := os.Getwd()
			if err != nil {
				return nil, err
			}
			abspath = filepath.Join(wd, path)
		}
		t2 := t.First()
		for t2 != nil {
			abspath2 := t2.Path
			if !filepath.IsAbs(abspath2) {
				wd, err := os.Getwd()
				if err != nil {
					return nil, err
				}
				abspath2 = filepath.Join(wd, abspath2)
			}

			if abspath == abspath2 {
				return t2, nil
			}

			t2 = t2.Next
		}

		// Open new tab
		newT, err := tab.Open(t.Frame, path)
		newT.Go(openLine, openCol)
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
		path := "error.enix"
		idx := 2
		for t.Exists(path) {
			path = fmt.Sprintf("error-%d.enix", idx)
			idx++
		}

		errTab := tab.FromString(t.Frame, errMsg, path)
		t.Append(errTab)
		newCurrentTab = errTab
	}

	return newCurrentTab, nil
}
