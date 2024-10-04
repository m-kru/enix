package arg

var Config string // Path to the config file.
var DumpConfig bool
var DumpKeys bool
var DumpPromptKeys bool

var Script string // Path to the script to be run.
var Line int = 1
var Column int = 1
var Files []string // Paths to text files.

func isValidFlag(f string) bool {
	flags := map[string]bool{
		"-dump-config":      true,
		"-dump-keys":        true,
		"-dump-prompt-keys": true,
		"-help":             true,
		"-version":          true,
	}
	if _, ok := flags[f]; ok {
		return true
	}
	return false
}

func isValidParam(p string) bool {
	params := map[string]bool{
		"-config": true, "-script": true,
	}
	if _, ok := params[p]; ok {
		return true
	}
	return false
}
