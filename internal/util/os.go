package util

import (
	"os"
	"strings"
)

// DirEntires returns names of entries within provided path.
// Additionally it filters entires by provided prefix.
// Directories names are suffixed with os path separator.
// Errors returned by os.ReadDir are ignored.
func DirEntries(path string, prefix string) []string {
	if path == "" {
		path = "."
	}

	entries, _ := os.ReadDir(path)

	names := make([]string, 0, len(entries))

	for _, e := range entries {
		name := e.Name()
		if e.IsDir() {
			name += string(os.PathSeparator)
		}
		if strings.HasPrefix(name, prefix) {
			names = append(names, name)
		}
	}

	return names
}
