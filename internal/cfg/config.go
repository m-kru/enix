package cfg

var Cfg Config

type Config struct {
	Colors string
	Style  string

	AutoSave int
	// Trim trailing whitespaces on save.
	// It affects only saves explicitly called by the user.
	// Neither automatic nor backup saves depend on this value.
	TrimOnSave          bool
	SafeFileSave        bool
	HighlightCursorWord bool

	// Whiespace
	LineEndRune rune
	TabRune     rune
	TabPadRune  rune

	MouseScrollMultiplier int

	UndoSize int // Undo and Redo stack size

	Extensions     map[string]string
	FiletypeIndent map[string]string
}

func DefaultConfig() Config {
	return Config{
		Colors:                "",
		Style:                 "",
		AutoSave:              0,
		TrimOnSave:            true,
		SafeFileSave:          true,
		HighlightCursorWord:   true,
		LineEndRune:           '¬',
		TabRune:               '▸',
		TabPadRune:            '·',
		MouseScrollMultiplier: 5,
		UndoSize:              1024,
		Extensions:            map[string]string{},
		FiletypeIndent: map[string]string{
			"fbdl":   "  ",
			"python": "    ",
			"rust":   "    ",
			"tcl":    "    ",
			"vhdl":   "  ",
		},
	}
}

func (cfg Config) GetIndent(filetype string) string {
	if indent, ok := cfg.FiletypeIndent[filetype]; ok {
		return indent
	}
	return "\t"
}

func (cfg Config) GetFileType(fileExt string) string {
	if filetype, ok := cfg.Extensions[fileExt]; ok {
		return filetype
	}
	return ""
}
