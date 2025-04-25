package tab

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/m-kru/enix/internal/cfg"
	"github.com/m-kru/enix/internal/util"
)

func (tab *Tab) WriteTo(strWr io.StringWriter) error {
	l := tab.Lines
	i := 1
	for l != nil {
		nl := tab.Newline
		if l.Next == nil {
			nl = ""
		}
		_, err := strWr.WriteString(fmt.Sprintf("%s%s", l.String(), nl))
		if err != nil {
			return fmt.Errorf("%s:%d: %v", tab.Path, i, err)
		}

		l = l.Next
		i++
	}

	tab.UndoCount = 0
	tab.RedoCount = 0

	return nil
}

func (tab *Tab) Save(path string, trim bool) (string, error) {
	if trim {
		tab.Trim()
	}

	tab.ModMutex.Lock()
	defer tab.ModMutex.Unlock()

	var info string
	var err error
	if cfg.Cfg.SafeFileSave {
		info, err = tab.safeSave(path)
	} else {
		info, err = tab.save(path)
	}

	tab.ModTime = util.FileModTime(tab.Path)

	return info, err
}

func (tab *Tab) save(path string) (string, error) {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = tab.WriteTo(file)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = file.Close()
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	return fmt.Sprintf("tab saved to file %s", path), nil
}

func (tab *Tab) safeSave(path string) (string, error) {
	// If file doesn't exist use regular save.
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return tab.save(path)
	}

	// First write content to the backup file
	backupPath := path + ".enix-bak"
	backupFile, err := os.OpenFile(backupPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0664)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = tab.WriteTo(backupFile)
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	err = backupFile.Close()
	if err != nil {
		return "", fmt.Errorf("save: %v", err)
	}

	// Write content to the root file
	_, err = tab.save(path)
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

func (tab *Tab) AutoSave() {
	// Don't save if there are no changes.
	if !tab.HasChanges() {
		return
	}

	// Don't save if file doesn't yet exist.
	if _, err := os.Stat(tab.Path); errors.Is(err, os.ErrNotExist) {
		return
	}

	if tab.State == "insert" && len(tab.InsertActions) > 0 {
		tab.undoPushInInsert()
	}

	// Don't trim whitespaces in autosave.
	// Ignore infos and erros returned by Save.
	_, _ = tab.Save(tab.Path, false)
}
