package cfg

type Config struct {
	// Whiespace
	NewlineRune rune
	TabWidth    int
	TabRune     rune
	TabPadRune  rune
}

func ConfigDefault() Config {
	return Config{
		NewlineRune: '¬',
		TabWidth:    4,
		TabRune:     '▸',
		TabPadRune:  '·',
	}
}
