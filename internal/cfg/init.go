package cfg

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"

	"github.com/m-kru/enix/internal/arg"

	homedir "github.com/mitchellh/go-homedir"
)

func init() {
	var err error

	ConfigDir = os.Getenv("ENIX_CONFIG_HOME")
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
func Init() (Config, Colorscheme, Keybindings, Keybindings, Keybindings, error) {
	config := ConfigDefault()
	colorscheme := ColorschemeDefault()
	keys := KeybindingsDefault()
	promptKeys := PromptKeybindingsDefault()
	insertKeys := InsertKeybindingsDefault()

	var err error

	if arg.Config != "" {
		config, err = configFromFile(arg.Config)
		if err != nil {
			goto exit
		}
	}

	if config.Colorscheme != "default" {
		colorscheme, err = colorschemeFromJSON(config.Colorscheme)
		if err != nil {
			goto exit
		}
	}

exit:
	return config, colorscheme, keys, promptKeys, insertKeys, err
}

func configFromFile(path string) (Config, error) {
	config := ConfigDefault()

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

	return config, nil
}
