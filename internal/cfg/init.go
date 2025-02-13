package cfg

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/m-kru/enix/internal/arg"

	"github.com/mattn/go-runewidth"
)

// Function Init initializes and returns various configurations at the program start.
func Init() error {
	Cfg = DefaultConfig()
	Style = DefaultStyle()
	Keys = DefaultKeybindings()
	KeysInsert = DefaultInsertKeybindings()
	KeysPrompt = DefaultPromptKeybindings()

	var err error

	if arg.Config != "" {
		Cfg, err = configFromFile(arg.Config)
		if err != nil {
			goto exit
		}
	}

	if Cfg.Style != "default" {
		Style, err = styleFromJSON(Cfg.Style)
		if err != nil {
			goto exit
		}
	}

exit:
	return err
}

func configFromFile(path string) (Config, error) {
	config := DefaultConfig()

	file, err := os.ReadFile(path)
	if err != nil {
		return config, fmt.Errorf(
			"reading config from %s: %v", path, err,
		)
	}

	err = json.Unmarshal(file, &config)
	if err != nil {
		return config, fmt.Errorf(
			"reading config from %s: %v", path, err,
		)
	}

	err = configSanityChecks()
	if err != nil {
		return config, fmt.Errorf(
			"reading config from %s: %v", path, err,
		)
	}

	return config, nil
}

func configSanityChecks() error {
	if Cfg.AutoSave < 0 {
		return fmt.Errorf("AutoSave must be natural, current value %d", Cfg.AutoSave)
	}

	rw := runewidth.RuneWidth(Cfg.LineEndRune)
	if rw != 1 {
		return fmt.Errorf(
			"width of LineEndRune must equal 1, width of '%c' equals %d",
			Cfg.LineEndRune, rw,
		)
	}

	if Cfg.UndoSize < 0 {
		return fmt.Errorf("UndoSize must be natural, current value %d", Cfg.UndoSize)
	}

	return nil
}
