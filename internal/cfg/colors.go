package cfg

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
	c := Colors{}

	return c, nil
}
