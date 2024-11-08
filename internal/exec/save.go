package exec

import (
	"errors"
	"fmt"
	"os"

	"github.com/m-kru/enix/internal/tab"
)

func Save(args []string, tab *tab.Tab, trim bool) (string, error) {
	if len(args) > 1 {
		return "", fmt.Errorf("save: expected at most 1 arg, provided %d", len(args))
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

func save(tab *tab.Tab, path string) (string, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = tab.Save(file)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = file.Close()
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	return fmt.Sprintf("tab saved to file %s", path), nil
}

func safeSave(tab *tab.Tab, path string) (string, error) {
	// If file doesn't exist use regular save.
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return save(tab, path)
	}

	// First write content to the backup file
	backupPath := path + ".enix-bak"
	backupFile, err := os.OpenFile(backupPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = tab.Save(backupFile)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = backupFile.Close()
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	// Write content to the root file
	_, err = save(tab, path)
	if err != nil {
		return "", err
	}

	// Remove backup file
	err = os.Remove(backupPath)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("tab saved to file %s", path), nil
}
