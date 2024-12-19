package cfg

// Path to the enix config home directory.
var ConfigDir string

var Cfg Config

type Config struct {
	Colorscheme string

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

	UndoSize int // Undo and Redo stack size

	Indent map[string]string
}

func DefaultConfig() Config {
	return Config{
		Colorscheme:         "default",
		TrimOnSave:          true,
		SafeFileSave:        true,
		HighlightCursorWord: true,
		LineEndRune:         '¬',
		TabRune:             '▸',
		TabPadRune:          '·',
		UndoSize:            1024,
		Indent: map[string]string{
			"python": "    ",
			"rust":   "    ",
			"vhdl":   "  ",
		},
	}
}

func (cfg Config) GetIndent(lang string) string {
	if indent, ok := cfg.Indent[lang]; ok {
		return indent
	}
	return "\t"
}
