package cfg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/m-kru/enix/internal/arg"

	homedir "github.com/mitchellh/go-homedir"

	"github.com/mattn/go-runewidth"
)

func init() {
	var err error

	ConfigDir = os.Getenv("ENIX_CONFIG_DIR")
	if ConfigDir == "" {
		ConfigDir = os.Getenv("XDG_CONFIG_HOME")

		if ConfigDir == "" {
			ConfigDir, err = homedir.Dir()
			if err != nil {
				log.Fatalf("can't determine home directory: %v", err)
			}
			ConfigDir = filepath.Join(ConfigDir, ".config")
		}

		ConfigDir = filepath.Join(ConfigDir, "enix")
	}
}

// Function Init initializes and returns various configurations at the program start.
func Init() (Config, Keybindings, Keybindings, Keybindings, error) {
	config := DefaultConfig()
	Colors = DefaultColorscheme()
	keys := DefaultKeybindings()
	promptKeys := DefaultPromptKeybindings()
	insertKeys := DefaultInsertKeybindings()

	var err error

	if arg.Config != "" {
		config, err = configFromFile(arg.Config)
		if err != nil {
			goto exit
		}
	}

	if config.Colorscheme != "default" {
		Colors, err = colorschemeFromJSON(config.Colorscheme)
		if err != nil {
			goto exit
		}
	}

exit:
	return config, keys, promptKeys, insertKeys, err
}

func configFromFile(path string) (Config, error) {
	config := DefaultConfig()

	file, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf(
			"reading config from file %s: %v", path, err,
		)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf(
			"reading config from file %s: %v", path, err,
		)
	}

	rw := runewidth.RuneWidth(config.LineEndRune)
	if rw != 1 {
		return config, fmt.Errorf(
			"reading config from file %s, width of line end rune must equal 1, width of '%c' equals %d",
			path, config.LineEndRune, rw,
		)
	}

	return config, nil
}
