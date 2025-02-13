package cfg

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/m-kru/enix/internal/arg"

	"github.com/mattn/go-runewidth"
)

// Function Init initializes and returns various configurations at the program start.
func Init() error {
	Keys = DefaultKeybindings()
	KeysInsert = DefaultInsertKeybindings()
	KeysPrompt = DefaultPromptKeybindings()

	var err error

	// Reading config from command line arguemnt takes precedence
	// over default config initialization.
	if arg.Config != "" {
		err = initCfgFromFile(arg.Config)
	} else {
		err = initCfg()
	}
	if err != nil {
		return err
	}

	err = initColors()
	if err != nil {
		return err
	}

	err = initStyle()
	if err != nil {
		return err
	}

	return nil
}

func initCfg() error {
	Cfg = DefaultConfig()

	bytes, path, err := ReadConfigFile("config.json")
	if err != nil {
		return fmt.Errorf("reading config.json: %v", err)
	}

	err = json.Unmarshal(bytes, &Cfg)
	if err != nil {
		return fmt.Errorf("reading config from %s: %v", path, err)
	}

	err = configSanityChecks()
	if err != nil {
		return fmt.Errorf("reading config from %s: %v", path, err)
	}

	return nil
}

func initCfgFromFile(path string) error {
	Cfg = DefaultConfig()

	file, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("reading config from %s: %v", path, err)
	}

	err = json.Unmarshal(file, &Cfg)
	if err != nil {
		return fmt.Errorf("reading config from %s: %v", path, err)
	}

	err = configSanityChecks()
	if err != nil {
		return fmt.Errorf("reading config from %s: %v", path, err)
	}

	return nil
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

func initColors() error {
	bytes, path, err := ReadConfigFile(filepath.Join("colors", Cfg.Colors+".json"))
	if err != nil {
		return err
	}
	if path == "" {
		Colors = DefaultColors()
		return nil
	}

	err = json.Unmarshal(bytes, &ColorsJSON)
	if err != nil {
		return fmt.Errorf("reading colors from %s: %v", path, err)
	}

	Colors, err = ColorsJSON.ToColors()
	if err != nil {
		return fmt.Errorf("reading colors from file %s: %v", path, err)
	}

	return nil
}

func initStyle() error {
	bytes, path, err := ReadConfigFile(filepath.Join("style", Cfg.Style+".json"))
	if err != nil {
		return err
	}
	if path == "" {
		Style = DefaultStyle()
		return nil
	}

	err = json.Unmarshal(bytes, &StyleJSON)
	if err != nil {
		return fmt.Errorf("reading style from %s: %v", path, err)
	}

	Style, err = StyleJSON.ToStyle()
	if err != nil {
		return fmt.Errorf("reading style from file %s: %v", path, err)
	}

	return nil
}
