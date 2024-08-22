package cfg

// Function Init initializes and returns various configurations at the program start.
func Init() (Config, Colorscheme, Keybindings, Keybindings, error) {
	return ConfigDefault(), ColorschemeDefault(), KeybindingsDefault(), PromptKeybindingsDefault(), nil
}
