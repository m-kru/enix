package cfg

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/mitchellh/go-homedir"
)

// ReadConfigFile tries to read configuration file from configuration directories.
// It returns:
//  1. config file content (if found),
//  2. absolute path of the read file,
//  3. error in case of any errors.
//
// If file doesn't exist, then it is not treated as an error.
// In such a case, the returned absolute path string is "".
func ReadConfigFile(file string) ([]byte, string, error) {
	var bytes []byte
	var err error

	path := os.Getenv("ENIX_CONFIG_DIR")
	if path == "" {
		goto xdg_dir_check
	}

	if strings.HasPrefix(path, "~") {
		usr, err := user.Current()
		if err != nil {
			return bytes, path, err
		}
		path = filepath.Join(usr.HomeDir, path[1:])
	}

	path = filepath.Join(path, file)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		goto xdg_dir_check
	}
	bytes, err = os.ReadFile(path)
	return bytes, path, err

xdg_dir_check:
	path = os.Getenv("XDG_CONFIG_HOME")
	hd, err := homedir.Dir()
	if err != nil {
		return bytes, "", err
	}
	if path == "" {
		path = filepath.Join(hd, ".config", "enix")
	}
	path = filepath.Join(path, file)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		goto install_dir_check
	}
	bytes, err = os.ReadFile(path)
	return bytes, path, err

install_dir_check:
	if runtime.GOOS == "windows" {
		panic("set correct default installation path here")
	} else {
		path = "/usr/local/share/enix"
	}
	path = filepath.Join(path, file)
	if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
		return bytes, "", nil
	}
	bytes, err = os.ReadFile(path)
	return bytes, path, err
}
