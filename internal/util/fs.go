package util

import (
	"os"
	"time"
)

// FileModTime returns file modification time.
// In case of any errors it returns the default time value.
func FileModTime(path string) time.Time {
	fi, err := os.Stat(path)
	if err != nil {
		return time.Time{}
	}

	return fi.ModTime()
}
