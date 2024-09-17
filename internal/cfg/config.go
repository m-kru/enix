package cfg

type Config struct {
	// Trim trailing whitespaces on save.
	// It affects only saves explicitly called by the user.
	// Neither automatic nor backup saves depend on this value.
	TrimOnSave   bool
	SafeFileSave bool

	// Whiespace
	NewlineRune rune
	TabWidth    int
	TabRune     rune
	TabPadRune  rune
}

func ConfigDefault() Config {
	return Config{
		TrimOnSave:   false,
		SafeFileSave: true,
		NewlineRune:  '¬',
		TabWidth:     4,
		TabRune:      '▸',
		TabPadRune:   '·',
	}
}
