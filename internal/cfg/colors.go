package cfg

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/m-kru/enix/internal/arg"
)

type Colors struct {
	Background    uint64
	Foreground    uint64
	Black         uint64
	Red           uint64
	Green         uint64
	Yellow        uint64
	Blue          uint64
	Magenta       uint64
	Cyan          uint64
	White         uint64
	BrightBlack   uint64
	BrightRed     uint64
	BrightGreen   uint64
	BrightYellow  uint64
	BrightBlue    uint64
	BrightMagenta uint64
	BrightCyan    uint64
	BrightWhite   uint64
}

// Read reads colors from file named "name.json".
func Read(name string) (Colors, error) {

	colorsDir := filepath.Join(ConfigDir, "colors")
	if arg.ColorsDir != "" {
		colorsDir = arg.ColorsDir
	}

	path := filepath.Join(colorsDir, name+".json")

	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("opening colors file: %v", err)
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		log.Fatalf("reading colors file: %v", err)
	}

	var colors map[string]any
	err = json.Unmarshal([]byte(data), &colors)
	if err != nil {
		log.Fatalf("unmarshalling colors file: %v", err)
	}

	c := Colors{
		getColorFromJSON(colors, "Background"),
		getColorFromJSON(colors, "Foreground"),
		getColorFromJSON(colors, "Black"),
		getColorFromJSON(colors, "Red"),
		getColorFromJSON(colors, "Green"),
		getColorFromJSON(colors, "Yellow"),
		getColorFromJSON(colors, "Blue"),
		getColorFromJSON(colors, "Magenta"),
		getColorFromJSON(colors, "Cyan"),
		getColorFromJSON(colors, "White"),
		getColorFromJSON(colors, "BrightBlack"),
		getColorFromJSON(colors, "BrightRed"),
		getColorFromJSON(colors, "BrightGreen"),
		getColorFromJSON(colors, "BrightYellow"),
		getColorFromJSON(colors, "BrightBlue"),
		getColorFromJSON(colors, "BrightMagenta"),
		getColorFromJSON(colors, "BrightCyan"),
		getColorFromJSON(colors, "BrightWhite"),
	}

	return c, nil
}

func getColorFromJSON(colors map[string]any, name string) uint64 {
	var val uint64
	var err error

	if value, ok := colors[name]; ok {
		if hex, ok := value.(string); ok {
			val, err = strconv.ParseUint(hex, 16, 64)
			if err != nil {
				log.Fatalf("can't convert value for %s color: %v", name, err)
			}
		} else {
			log.Fatalf("invalid value type for %s color, expected string", name)
		}
	} else {
		log.Fatalf("missing definition of %s color", name)
	}

	return val
}
