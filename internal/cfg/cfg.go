package cfg

// Function Init initializes and returns various configurations at the program start.
func Init() (Colorscheme, error) {
	return ColorschemeDefault(), nil
}
