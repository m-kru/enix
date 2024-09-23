package cmd

import (
	"fmt"
	"os"

	"github.com/m-kru/enix/internal/tab"
)

func Save(args []string, tab *tab.Tab, trim bool) error {
	if len(args) > 0 {
		return fmt.Errorf("save: expected 0 args, provided %d", len(args))
	}

	if trim {
		tab.Trim()
	}

	if tab.Config.SafeFileSave {
		return safeSave(tab)
	} else {
		return save(tab)
	}
}

func save(tab *tab.Tab) error {
	file, err := os.OpenFile(tab.Path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
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

func safeSave(tab *tab.Tab) error {
	return nil
}
