package cfg

// Function Init initializes and returns various configurations at the program start.
func Init() (Colorscheme, Keybindings, error) {
	return ColorschemeDefault(), KeybindingsDefault(), nil
}
