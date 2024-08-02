package cfg

type Config struct {
	TabWidth   int
	TabRune    rune
	TabPadRune rune
}

func ConfigDefault() Config {
	return Config{
		TabWidth: 4,
		TabRune: '▸',
		TabPadRune: '·',
	}
}
