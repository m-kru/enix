package cfg

type Config struct {
	SafeFileSave bool

	// Whiespace
	NewlineRune rune
	TabWidth    int
	TabRune     rune
	TabPadRune  rune
}

func ConfigDefault() Config {
	return Config{
		SafeFileSave: true,
		NewlineRune:  '¬',
		TabWidth:     4,
		TabRune:      '▸',
		TabPadRune:   '·',
	}
}
