package cmd

import (
	"fmt"
	"os"

	"github.com/m-kru/enix/internal/tab"
)

func Save(args []string, tab *tab.Tab, trim bool) error {
	if len(args) > 1 {
		return fmt.Errorf("save: expected at most 1 arg, provided %d", len(args))
	}

	path := tab.Path
	if len(args) == 1 {
		path = args[0]
	}

	if trim {
		tab.Trim()
	}

	if tab.Config.SafeFileSave {
		return safeSave(tab, path)
	} else {
		return save(tab, path)
	}
}

func save(tab *tab.Tab, path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return fmt.Errorf("save: %v", err)
	}

	err = tab.Save(file)
	if err != nil {
		return fmt.Errorf("save: %v", err)
	}

	err = file.Close()
	if err != nil {
		return fmt.Errorf("save: %v", err)
	}

	return nil
}

func safeSave(tab *tab.Tab, path string) error {
	return nil
}
